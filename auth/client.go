package auth

import (
	"context"

	pb "bitbucket.org/muulin/interlib/auth/service"
	"bitbucket.org/muulin/interlib/core"
	"github.com/94peter/sterna/auth"
	"google.golang.org/grpc/metadata"
)

type AuthClient interface {
	core.MyGrpc
	ValidateToken(host, token string) (auth.ReqUser, error)
	IsAccountsExist(host string, accounts []string) (notExistAccounts []string, err error)
}

func NewGrpcClient(address string) (AuthClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
}

func (gclt *grpcClt) ValidateToken(host, token string) (auth.ReqUser, error) {
	clt := pb.NewAuthServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	resp, err := clt.ValidateToken(ctx, &pb.ValidateTokenRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}
	return auth.NewReqUser(host, resp.Id, resp.Account, resp.Name, resp.Perms), nil
}

func (gclt *grpcClt) IsAccountsExist(host string, accounts []string) (notExistAccounts []string, err error) {
	clt := pb.NewAuthServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	resp, err := clt.IsAccountsExist(ctx, &pb.IsAccountsExistRequest{
		Accounts: accounts,
	})
	if err != nil {
		return nil, err
	}
	return resp.NotExistAccounts, nil
}
