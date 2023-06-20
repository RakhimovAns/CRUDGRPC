package main

import (
	"context"
	"flag"
	GRPC "github.com/RakhimovAns/GRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	Text := flag.Arg(0)
	if Text == "" {
		log.Fatal("No text")
	}
	data := strings.Split(Text, ";")
	conn, err := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := GRPC.NewMovieServiceClient(conn)
	if len(data) == 2 {
		if data[0] == "GET" && data[1] == "movies" {
			res, err := c.GetMovies(context.Background(), &GRPC.ReadMoviesRequest{})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
	}
	if len(data) == 3 {
		if data[0] == "POST" && data[1] == "movies" {
			Movie := strings.Split(data[2], ",")
			movie := GRPC.Movie{
				Genre: Movie[1],
				Title: Movie[0],
			}
			res, err := c.CreatedMovie(context.Background(), &GRPC.CreateMovieRequest{Movie: &movie})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
		if data[0] == "GET" && data[1] == "movies" {
			res, err := c.GetMovie(context.Background(), &GRPC.ReadMovieRequest{Id: data[2]})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
		if data[0] == "PUT" && data[1] == "movies" {
			Movie := strings.Split(data[2], ",")
			movie := &GRPC.Movie{
				Id:    Movie[0],
				Genre: Movie[2],
				Title: Movie[1],
			}
			res, err := c.UpdateMovies(context.Background(), &GRPC.UpdateMovieRequest{Movie: movie})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
		if data[0] == "DELETE" && data[1] == "movies" {
			res, err := c.DeleteMovie(context.Background(), &GRPC.DeleteMovieRequest{Id: data[2]})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
	}
}
