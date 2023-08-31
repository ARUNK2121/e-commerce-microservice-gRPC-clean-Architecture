package service

import (
	"context"
	pb "users/pkg/proto"
	interfaces "users/pkg/usecase/interface"
	"users/pkg/utils/models"
)

type UserServer struct {
	adminUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}

func NewUserServer(useCase interfaces.UserUseCase) pb.UserServer {

	return &UserServer{
		adminUseCase: useCase,
	}

}

func (a *UserServer) UserSignUp(ctx context.Context, userSignUpDetails *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {

	userCreateDetails := models.UserSignUp{
		Name:     userSignUpDetails.Name,
		Email:    userSignUpDetails.Email,
		Phone:    userSignUpDetails.Phone,
		Password: userSignUpDetails.Password,
	}

	status, err := a.adminUseCase.UserSignUp(userCreateDetails)
	if err != nil {
		return &pb.UserSignUpResponse{
			Status: int64(status),
			Error:  err.Error(),
		}, err
	}

	return &pb.UserSignUpResponse{
		Status: int64(status),
		Error:  "",
	}, nil

}

func (a *UserServer) UserLogin(ctx context.Context, UserLoginDetails *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	userLogin := models.UserLogin{
		Email:    UserLoginDetails.Email,
		Password: UserLoginDetails.Password,
	}

	token, err := a.adminUseCase.UserLogin(userLogin)
	if err != nil {
		return &pb.UserLoginResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.UserLoginResponse{
		Token: token,
	}, nil

}
