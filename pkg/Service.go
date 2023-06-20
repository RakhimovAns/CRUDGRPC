package pkg

import (
	"context"
	"fmt"
	pb "github.com/RakhimovAns/GRPC/proto"
	"github.com/RakhimovAns/GRPC/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var pool *pgxpool.Pool
var err error

func DataBaseConnection() {
	dsn := "postgresql://postgres:Ansar@localhost:5432/test"
	connectCtx, _ := context.WithTimeout(context.Background(), time.Second*5)
	pool, err = pgxpool.Connect(connectCtx, dsn)
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
		return
	}
	fmt.Println("Database connection successful...")

}

type Server struct {
	pb.UnimplementedMovieServiceServer
}

func (*Server) CreatedMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Println("Create Movie")
	movie := req.GetMovie()
	movie.Id = uuid.New().String()
	data := types.Movie{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}
	_, err := pool.Exec(ctx, `
insert into movies(id, title, genre) VALUES ($1,$2,$3)
`, data.ID, data.Title, data.Genre)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func (*Server) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read Movie", req.GetId())
	movie := &types.Movie{}
	err := pool.QueryRow(ctx, `
select id, title, genre, created, updated from movies where id=$1
`, req.GetId()).Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}
func (*Server) GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	fmt.Println("Read Movies")
	movies := []*pb.Movie{}
	rows, err := pool.Query(ctx, `
select id,title,genre from movies
`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		movie := &pb.Movie{}
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Genre)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (*Server) UpdateMovies(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Println("Update Movie")
	movie := &types.Movie{}
	reqMovie := req.GetMovie()
	_, err := pool.Exec(ctx, `
update movies set title=$1,genre=$2,updated=current_timestamp where id=$3
`, reqMovie.Title, reqMovie.Genre, reqMovie.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = pool.QueryRow(ctx, `
select id, title, genre, created, updated from movies where id=$1
`, reqMovie.Id).Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Println("Delete Movie")
	_, err := pool.Exec(ctx, `
delete from movies where id=$1
`, req.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
