package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/tangoctrung/speed_test_id/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GenIDWithUUID(ctx context.Context, req *pb.GenIDRequest) (*pb.GenIDResponse, error) {
	log.Println("sum called...")
	res := &pb.GenIDResponse{
		Id: uuid.New().String(),
	}
	return res, nil
}

func (*server) GenIDWithSnowflake(ctx context.Context, req *pb.GenIDRequest) (*pb.GenIDResponse, error) {
	log.Println("sum called...")
	n, _ := snowflake.NewNode(1)
	res := &pb.GenIDResponse{
		Id: n.Generate().String(),
	}
	return res, nil
}

func main() {
	log.Println("Hello world !!!")

	//   -------------Goi ham gen ID-----------------
	// SequentialGenID()
	// utils.SyncGenUUID()
	// utils.SyncGenSnowflake()

	//   --------------HTTP gen ID-------------------
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Server running!!!",
	// 		"data":    nil,
	// 		"success": true,
	// 	})
	// })

	// api := r.Group("api")
	// {
	// 	api.GET("/uuid", controllers.GenUUID)
	// 	api.GET("/snowflake", controllers.GenSnowflakeID)
	// }

	// r.Run(":8000")

	// -----------------gRPC gen ID----------------------
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	grpc.WithInsecure()

	pb.RegisterGenIDServiceServer(s, &server{})

	fmt.Println("server is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}

}
