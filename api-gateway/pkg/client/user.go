package client

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/config"
	pb "api_gateway/pkg/proto/user"
	"api_gateway/pkg/utils/models"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type userClient struct {
	Client pb.UserClient
}

func NewUserClient(cfg config.Config) interfaces.UserClient {

	grpcConnection, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewUserClient(grpcConnection)

	return &userClient{
		Client: grpcClient,
	}

}

func (a *userClient) UserLogin(adminDetails models.UserLogin) (string, error) {

	fmt.Println(a.Client)

	res, err := a.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})
	if err != nil {
		return "", err
	}

	if res.Error != "" {
		return "", errors.New(res.Error)
	}

	return res.Token, nil

}

func (a *userClient) UserSignUp(admin models.UserSignUp) (int, error) {

	res, err := a.Client.UserSignUp(context.Background(), &pb.UserSignUpRequest{
		Name:            admin.Name,
		Email:           admin.Email,
		Password:        admin.Password,
		Confirmpassword: admin.ConfirmPassword,
	})

	if err != nil {
		return 400, err
	}

	return int(res.Status), nil

}
