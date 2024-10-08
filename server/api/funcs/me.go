package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-auth/api/errors"
	"github.com/sohamjaiswal/grpc-auth/api/meta"
	"github.com/sohamjaiswal/grpc-auth/global"
	"github.com/sohamjaiswal/grpc-auth/models"
	"github.com/sohamjaiswal/grpc-auth/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Me(ctx context.Context, req *pb.NoParams) (*pb.MeResponse, error) {
	userAuth, err := meta.AuthorizeUser(ctx)
	if err != nil {
		return nil, err
	}
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}
	var user models.User
	if err = db.Where("username = ?", userAuth.Username).First(&user).Error; err != nil {
		return nil, errors.UserNotFound()
	}
	return &pb.MeResponse{
		User: &pb.User{
			Username: *user.Username,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}