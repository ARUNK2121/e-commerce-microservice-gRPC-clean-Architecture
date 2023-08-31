package client

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/config"
	pb "api_gateway/pkg/proto/admin"
	"api_gateway/pkg/utils/models"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}

}

func (a *adminClient) AdminLogin(adminDetails models.AdminLogin) (string, error) {

	fmt.Println(a.Client)

	res, err := a.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
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

func (a *adminClient) AdminSignUp(admin models.AdminSignUp) (int, error) {

	res, err := a.Client.AdminSignUp(context.Background(), &pb.AdminSignUpRequest{
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
