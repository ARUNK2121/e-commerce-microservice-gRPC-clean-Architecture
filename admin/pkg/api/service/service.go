package service

import (
	pb "admin/pkg/proto"
	interfaces "admin/pkg/usecase/interface"
	"admin/pkg/utils/models"
	"context"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}

func (a *AdminServer) AdminSignUp(ctx context.Context, userSignUpDetails *pb.AdminSignUpRequest) (*pb.AdminSignUpResponse, error) {

	adminCreateDetails := models.AdminSignUp{
		Name:     userSignUpDetails.Name,
		Email:    userSignUpDetails.Email,
		Password: userSignUpDetails.Password,
	}

	status, err := a.adminUseCase.AdminSignUp(adminCreateDetails)
	if err != nil {
		return &pb.AdminSignUpResponse{
			Status: int64(status),
			Error:  err.Error(),
		}, err
	}

	return &pb.AdminSignUpResponse{
		Status: int64(status),
		Error:  "",
	}, nil

}

func (a *AdminServer) AdminLogin(ctx context.Context, adminLoginDetails *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {

	adminLogin := models.AdminLogin{
		Email:    adminLoginDetails.Email,
		Password: adminLoginDetails.Password,
	}

	token, err := a.adminUseCase.AdminLogin(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.AdminLoginResponse{
		Token: token,
	}, nil

}
