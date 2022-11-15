package handlers

import (
	"context"

	"google.golang.org/grpc/status"

	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// PushObject - gRPC метод создания КЕ
func (grpc *GRPC) PushObject(ctx context.Context, in *pb.PushObjectRequest) (*pb.PushObjectResponse, error) {
	var resp pb.PushObjectResponse
	err := grpc.App.CreateObject(ctx, in.Item)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// GetObject - gRPC метод получения КЕ
func (grpc *GRPC) GetObject(ctx context.Context, in *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	var resp pb.GetObjectResponse
	ci, err := grpc.App.GetObject(ctx, in.Name)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Item = ci
	resp.Code = "OK"
	return &resp, nil
}

// DeleteObject - gRPC метод удаления КЕ
func (grpc *GRPC) DeleteObject(ctx context.Context, in *pb.DeleteObjectRequest) (*pb.DeleteObjectResponse, error) {
	var resp pb.DeleteObjectResponse
	err := grpc.App.DeleteObject(ctx, in.Name)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// PushBatchObject - gRPC метод создания КЕ пакетно
func (grpc *GRPC) PushBatchObject(ctx context.Context, in *pb.PushBatchObjectsRequest) (*pb.PushBatchObjectsResponse, error) {
	var resp pb.PushBatchObjectsResponse
	err := grpc.App.CreateBatchObjects(ctx, in.Item)
	if err != nil {
		code, _ := ErrCodesMapping(err)
		resp.Code = code
		return &resp, nil
	}
	resp.Code = "OK"
	return &resp, nil
}

// GetBatchObject - gRPC метод получения КЕ пакетно
func (grpc *GRPC) GetBatchObject(ctx context.Context, in *pb.GetBatchObjectsRequest) (*pb.GetBatchObjectsResponse, error) {
	var resp pb.GetBatchObjectsResponse
	cis, err := grpc.App.GetBatchObjects(ctx, in.Item)
	if err != nil {
		code, _ := ErrCodesMapping(err)
		resp.Item = cis
		resp.Code = code
		return &resp, nil
	}
	resp.Item = cis
	resp.Code = "OK"
	return &resp, nil
}

// DeleteBatchObject - gRPC метод удаления КЕ пакетно
func (grpc *GRPC) DeleteBatchObject(ctx context.Context, in *pb.GetBatchObjectsRequest) (*pb.GetBatchObjectsResponse, error) {
	var resp pb.GetBatchObjectsResponse
	err := grpc.App.DeleteBatchObject(ctx, in.Item)
	if err != nil {
		code, _ := ErrCodesMapping(err)
		resp.Code = code
		return &resp, nil
	}
	resp.Code = "OK"
	return &resp, nil
}
