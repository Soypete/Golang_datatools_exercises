package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gen "github.com/soypete/Golang_datatools_exercises/grpc-gateway-demo/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	gen.UnimplementedUserActionServer
	GWServer *http.Server
}

func SetupGrpc() *Server {
	return &Server{}
}

// SetupGateway creates the Rest server via grpc connection/
func (s *Server) SetupGateway(ctx context.Context, port string, grpcPort string) error {
	conn, err := grpc.DialContext(
		ctx,
		"localhost"+grpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return fmt.Errorf("failed to setup grpc client connection: %w", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	if err = gen.RegisterUserActionHandler(ctx, gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	s.GWServer = &http.Server{
		Addr:    "localhost" + port,
		Handler: gwmux,
	}

	return nil
}

// startup and run grpc server
func (s *Server) RunGrpc(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
		return fmt.Errorf("cannot setup tcp connection: %w", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterUserActionServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("grpc server failure: %w", err)
	}
	return nil
}

// startup and run grpc-gateway server
func (s *Server) Close(ctx context.Context) error {
	ctx.Done()
	return nil
}

// UpdateUser updates a user
// curl localhost:8090/user -X -d '{"id": 1, "name": "SoyPete", "email": "captainnobody1@gmail", "password": "password"}'
func (s *Server) UpdateUser(ctx context.Context, in *gen.User) (*emptypb.Empty, error) {
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

// GetUserInfo gets a user
// curl localhost:8090/user/getinfo/1
func (s *Server) GetUserInfo(ctx context.Context, in *gen.UserIdentifier) (*gen.User, error) {
	// validate required fields, ID default is 0. 0 is not a valid ID
	if in.UserId == 0 {
		return nil, errors.New("required ID fields")
	}

	// NOTE: This is where we would call the database to get the user
	// Database index would start at 1
	return &gen.User{
		Id:    in.UserId,
		Name:  "SoyPete",
		Email: "captainnobody1@email.com",
	}, nil
}

// DeleteUser deletes a user
// curl localhost:8090/user/delete/1 -X DELETE
func (s *Server) DeleteUser(ctx context.Context, in *gen.UserIdentifier) (*emptypb.Empty, error) {
	// validate required fields, ID default is 0. 0 is not a valid ID
	if in.UserId == 0 {
		return nil, errors.New("required ID fields")
	}
	return &emptypb.Empty{}, nil
}

// SendUserAction sends a user action
// curl localhost:8090/user/action -X -d '{"action": 1, "userId": 1}'
func (s *Server) SendUserAction(ctx context.Context, in *gen.Action) (*emptypb.Empty, error) {
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
