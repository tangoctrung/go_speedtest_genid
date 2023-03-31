package utils

import (
	"context"
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/tangoctrung/speed_test_id/pb"
)

type server struct{}

func (*server) GenIDWithUUID(ctx context.Context) (*pb.GenIDResponse, error) {
	log.Println("sum called...")
	res := &pb.GenIDResponse{
		Id: uuid.New().String(),
	}
	return res, nil
}

func (*server) GenIDWithSnowflake(ctx context.Context) (*pb.GenIDResponse, error) {
	log.Println("sum called...")
	n, _ := snowflake.NewNode(1)
	res := &pb.GenIDResponse{
		Id: n.Generate().String(),
	}
	return res, nil
}
