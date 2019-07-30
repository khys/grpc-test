package service

import (
	"context"
	"errors"
	"afc-mano/proto"
)

type ResourceService struct {
}

func (s *ResourceService) GetResources(ctx context.Context, message *proto.GetResourcesMessage) (*proto.ResourcesResponse, error) {
	switch message.TargetResource {
	case "cpu":
		return &proto.ResourcesResponse{
			Id: "0",
			Value: "32",
		}, nil
	case "ram":
		return &proto.ResourcesResponse{
			Id: "1",
			Value: "128",
		}, nil
	}
	return nil, errors.New("Not Found the Resources")
}
