package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/koga456/rock-paper-scissors/pb"
	"github.com/koga456/rock-paper-scissors/pkg"
)

func PlayGame(handShapes int32) {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)
	playRequest := pb.PlayRequest{
		HandShapes: pkg.EncodeHandShapes(handShapes),
	}

	reply, err := client.PlayGame(ctx, &playRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	marchResult := reply.GetMatchResult()
	fmt.Println("***********************************")
	fmt.Printf("Your hand shapes: %s \n", marchResult.YourHandShapes.String())
	fmt.Printf("Opponent hand shapes: %s \n", marchResult.OpponentHandShapes.String())
	fmt.Printf("Result: %s \n", marchResult.Result.String())
	fmt.Println("***********************************")
	fmt.Println()
}

func ReportMatchResults() {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)
	reportRequest := pb.ReportRequest{}

	reply, err := client.ReportMatchResults(ctx, &reportRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	report := reply.GetReport()
	if len(report.MatchResults) == 0 {
		fmt.Println("***********************************")
		fmt.Println("There are no match results.")
		fmt.Println("***********************************")
		fmt.Println()
		return
	}

	fmt.Println("***********************************")
	for k, v := range report.MatchResults {
		fmt.Println(k + 1)
		fmt.Printf("Your hand shapes: %s \n", v.YourHandShapes.String())
		fmt.Printf("Opponent hand shapes: %s \n", v.OpponentHandShapes.String())
		fmt.Printf("Result: %s \n", v.Result.String())
		fmt.Printf("Datetime of match: %s \n", v.CreateTime.AsTime().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(time.ANSIC))
		fmt.Println()
	}

	fmt.Printf("Number of games: %d \n", reply.GetReport().NumberOfGames)
	fmt.Printf("Number of wins: %d \n", reply.GetReport().NumberOfWins)
	fmt.Println("***********************************")
	fmt.Println()
}
