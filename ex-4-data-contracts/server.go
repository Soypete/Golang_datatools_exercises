package server

import (
	"context"

	gengo "github.com/soypete/Golang_datatools_exercises/ex-4-data-contracts/gen/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	gengo.UnimplementedUserActionServer
}

func setupServer() *Server {
	return &Server{}
}

func (s *Server) Close(ctx context.Context) error {
	ctx.Done()
	return nil
}

func (s *Server) UpdateUser(ctx context.Context, in *gengo.User) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) GetUserInfo(ctx context.Context, in *gengo.UserIdentifier) (*gengo.User, error) {
	return nil, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *gengo.UserIdentifier) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) SendUserAction(ctx context.Context, in *gengo.Action) (*emptypb.Empty, error) {
	return nil, nil
}
