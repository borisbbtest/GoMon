package handlers

import (
	"context"

	"google.golang.org/grpc/status"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// GetSession - gRPC метод получения сессии
func (grpc *GRPC) GetSession(ctx context.Context, in *pb.GetSessionRequest) (*pb.GetSessionResponse, error) {
	var resp pb.GetSessionResponse
	session, err := grpc.App.GetSession(ctx, in.Login, in.Id)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Ss = session
	resp.Code = "OK"
	return &resp, nil
}

// CreationSession - gRPC метод создания сессии
func (grpc *GRPC) CreationSession(ctx context.Context, in *pb.CreationSessionRequest) (*pb.CreationSessionResponse, error) {
	var resp pb.CreationSessionResponse
	_, err := grpc.App.CreateSession(ctx, in.Person)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// DeletionSession - gRPC метод удаления сессии
func (grpc *GRPC) DeletionSession(ctx context.Context, in *pb.DeletionSessionRequest) (*pb.DeletionSessionResponse, error) {
	var resp pb.DeletionSessionResponse
	err := grpc.App.DeleteSession(ctx, in.Login, in.Id)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// GetAllSession - gRPC метод получения списка всех сессий
func (grpc *GRPC) GetAllSession(ctx context.Context, in *pb.GetAllSessionRequest) (*pb.GetAllSessionResponse, error) {
	var resp pb.GetAllSessionResponse
	result, err := grpc.App.GetAllSessions(ctx)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	resp.Ss = result
	return &resp, nil
}

// CreationUser - gRPC метод создания пользователя
func (grpc *GRPC) CreationUser(ctx context.Context, in *pb.CreationUserRequest) (*pb.CreationUserResponse, error) {
	var resp pb.CreationUserResponse
	err := grpc.App.CreateUser(ctx, in.Persone)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// DeletionUser - gRPC метод удаления пользователя
func (grpc *GRPC) DeletionUser(ctx context.Context, in *pb.DeletionUserRequest) (*pb.DeletionUserResponse, error) {
	var resp pb.DeletionUserResponse
	err := grpc.App.DeleteUser(ctx, in.Login)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	return &resp, nil
}

// GetUser - gRPC метод получения пользователя
func (grpc *GRPC) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var resp pb.GetUserResponse
	user, err := grpc.App.GetUser(ctx, in.Person)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	resp.Person = user
	return &resp, nil
}

// GetListUserAll - gRPC метод получения списка всех пользователей
func (grpc *GRPC) GetListUserAll(ctx context.Context, in *pb.GetListUserAllRequest) (*pb.GetListUserAllResponse, error) {
	var resp pb.GetListUserAllResponse
	users, err := grpc.App.GetAllUsers(ctx)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	resp.Persons = users
	return &resp, nil
}

// Authorization - gRPC метод авторизации пользователя
func (grpc *GRPC) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	var resp pb.AuthorizationResponse
	session, err := grpc.App.AuthorizeUser(ctx, in.Login, in.Password)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	resp.Ss = session
	return &resp, nil
}

// Registration - gRPC метод регистрации нового пользователя
func (grpc *GRPC) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	var resp pb.RegistrationResponse
	session, err := grpc.App.RegisterUser(ctx, in.Person)
	if err != nil {
		code, gcode := ErrCodesMapping(err)
		resp.Code = code
		return &resp, status.Error(gcode, err.Error())
	}
	resp.Code = "OK"
	resp.Ss = session
	return &resp, nil
}
