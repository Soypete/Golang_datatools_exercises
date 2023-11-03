package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	gengo "github.com/soypete/Golang_datatools_exercises/ex-4-data-contracts/gen/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer_UpdateUser(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		in      *gengo.User
		want    *emptypb.Empty
		wantErr bool
	}{
		{
			name: "Update User:Success",
			ctx:  context.Background(),
			in: &gengo.User{
				Id:       1,
				Name:     "SoyPete",
				Password: "password",
				Email:    "captainnobody1@gmail.com",
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name: "Update User:Fail - Invalid Email",
			ctx:  context.Background(),
			in: &gengo.User{
				Id:       1,
				Name:     "SoyPete",
				Email:    "captainnobody1",
				Password: "password",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Update User:Fail - No Password",
			ctx:  context.Background(),
			in: &gengo.User{
				Id:    1,
				Name:  "SoyPete",
				Email: "captainnobody1@email.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Update User:Fail - No Email",
			ctx:  context.Background(),
			in: &gengo.User{
				Id:       1,
				Name:     "SoyPete",
				Password: "password",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupGrpc()
			got, err := s.UpdateUser(tt.ctx, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUserInfo(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		in      *gengo.UserIdentifier
		want    *gengo.User
		wantErr bool
	}{
		{
			name: "Get User:Success",
			ctx:  context.Background(),
			in: &gengo.UserIdentifier{
				UserId: 1,
			},
			want: &gengo.User{
				Id:    1,
				Name:  "SoyPete",
				Email: "captainnobody1@email.com",
				// Do not want to return password
			},
			wantErr: false,
		},
		{
			name:    "Get User:Fail - Invalid ID",
			ctx:     context.Background(),
			in:      &gengo.UserIdentifier{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupGrpc()
			got, err := s.GetUserInfo(tt.ctx, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		in      *gengo.UserIdentifier
		want    *emptypb.Empty
		wantErr bool
	}{
		{
			name: "Delete User:Success",
			ctx:  context.Background(),
			in: &gengo.UserIdentifier{
				UserId: 1,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name:    "Delete User:Fail - Invalid ID",
			ctx:     context.Background(),
			in:      &gengo.UserIdentifier{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupGrpc()
			got, err := s.DeleteUser(tt.ctx, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_SendUserAction(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		in      *gengo.Action
		want    *emptypb.Empty
		wantErr bool
	}{
		{
			name: "Send User Action:Success",
			ctx:  context.Background(),
			in: &gengo.Action{
				UserId: 1,
				Action: gengo.ActionEnum_LOGIN,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name: "Send User Action:Success - full payload",
			ctx:  context.Background(),
			in: &gengo.Action{
				UserId:    1,
				Action:    gengo.ActionEnum_LOGOUT,
				Timestamp: timestamppb.New(time.Now()),
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name: "Send User Action:Success - fail payload",
			ctx:  context.Background(),
			in: &gengo.Action{
				UserId:    1,
				Action:    gengo.ActionEnum_VIEW_CART,
				Success:   false,
				Error:     "timeout on request",
				Timestamp: timestamppb.New(time.Now()),
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name: "Send User Action:Fail",
			ctx:  context.Background(),
			in: &gengo.Action{
				UserId: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupGrpc()
			got, err := s.SendUserAction(tt.ctx, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.SendUserAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.SendUserAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
