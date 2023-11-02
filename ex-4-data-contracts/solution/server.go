package server

import (
	"context"
	"errors"
	"strings"

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
	//validate required fields
	if in.Email == "" || in.Password == "" {
		return nil, errors.New("missing required fields")
	}
	//validate email
	if !strings.Contains(in.Email, "@") {
		return nil, errors.New("invalid email")
	}

	// NOTE: This is where we would call the database to update the user
	// Database index would start at 1 and is not required because we can use Sequential ID
	return &emptypb.Empty{}, nil
}

func (s *Server) GetUserInfo(ctx context.Context, in *gengo.UserIdentifier) (*gengo.User, error) {
	// validate required fields, ID default is 0. 0 is not a valid ID
	if in.UserId == 0 {
		return nil, errors.New("required ID fields")
	}

	// NOTE: This is where we would call the database to get the user
	// Database index would start at 1
	return &gengo.User{
		Id:    in.UserId,
		Name:  "SoyPete",
		Email: "captainnobody1@email.com",
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *gengo.UserIdentifier) (*emptypb.Empty, error) {
	// validate required fields, ID default is 0. 0 is not a valid ID
	if in.UserId == 0 {
		return nil, errors.New("required ID fields")
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) SendUserAction(ctx context.Context, in *gengo.Action) (*emptypb.Empty, error) {
	// validate required fields. 0 Is the Enum default value
	if in.Action == 0 {
		return nil, errors.New("required Action fields")
	}

	// ID default is 0. 0 is not a valid ID
	if in.UserId == 0 {
		return nil, errors.New("required UserId fields")
	}

	// NOTE: This is where we would call the database to get the user

	return &emptypb.Empty{}, nil
}
