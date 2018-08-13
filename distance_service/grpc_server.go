package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/andrewdruzhinin/travel-cost-app/distance_service/trippb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct{}

func (s *gRPCServer) GetTripInfo(ctx context.Context, trip *pb.Trip) (*pb.TripInfoResponse, error) {
	distanceInfo, err := GetDistanceInfo(trip)
	if err != nil {
		return nil, err
	}
	distance := int64(distanceInfo.Distance)
	duration := int64(distanceInfo.Duration)
	return &pb.TripInfoResponse{Distance: distance, Duration: duration}, nil
}

// StartGRPCServer :
func StartGRPCServer() {
	gRPCPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTripInfoServer(s, &gRPCServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
