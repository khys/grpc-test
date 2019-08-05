package service

import (
	"context"
	"errors"
	"afc-mano/proto"
)

type ResourceService struct {
}

func (s *ResourceService) GetCpu(ctx context.Context, message *proto.GetCpuMessage) (*proto.CpuResponse, error) {
	switch message.Id {
	case "0":
		return &proto.CpuResponse{
			Max: "32",
			Used: "16",
			Price: "8",
		}, nil
	case "1":
		return &proto.CpuResponse{
			Max: "32",
			Used: "16",
			Price: "8",
		}, nil
	}
	return nil, errors.New("Not Found the Resources")
}

func (s *ResourceService) GetMemory(ctx context.Context, message *proto.GetMemoryMessage) (*proto.MemoryResponse, error) {
	switch message.Id {
	case "0":
		return &proto.MemoryResponse{
			Max: "128",
			Used: "120",
			Price: "2",
		}, nil
	case "1":
		return &proto.MemoryResponse{
			Max: "128",
			Used: "120",
			Price: "2",
		}, nil
	}
	return nil, errors.New("Not Found the Resources")
}