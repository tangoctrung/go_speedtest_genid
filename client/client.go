package main

import (
	"context"
	"log"
	"time"

	"github.com/tangoctrung/speed_test_id/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:8000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf(" err while dial %v", err)
	}
	defer cc.Close()

	client := pb.NewGenIDServiceClient(cc)

	callGenIDWithUUID(client, 1*time.Second) // bi timeout
}

func callGenIDWithUUID(c pb.GenIDServiceClient, timeout time.Duration) {
	log.Println("calling gen id with uuid api")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.GenIDWithUUID(ctx, &pb.GenIDRequest{
		// Message: "Hello server",
	})

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("calling gen id with uuid DeadlineExceeded")
			} else {
				log.Printf("calling gen id with uuid api err %v", err)
			}
		} else {
			log.Fatalf("calling gen id with uuid unknown err %v", err)
		}
		return
	}

	log.Printf("gen id with uuid response %v\n", resp.GetId())
}
