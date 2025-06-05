package usermodule

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "hermes/api/proto/user"
	"hermes/internal/constant"
	"hermes/internal/model/entity"
	"hermes/internal/repository/database"
	"math"
	"strconv"
)

const (
	ResponseCodeSuccess    = "00"
	ResponseCodeDBError    = "DBU"
	ResponseCodeNotFound   = "NFU"
	ResponseCodeValidation = "VDU"
)

type UserService struct {
	tblUser *database.TblUser
}

func NewUserService(tblUser *database.TblUser) *UserService {
	return &UserService{
		tblUser: tblUser,
	}
}

func (service *UserService) CheckUsername(ctx context.Context, username string) (*pb.CheckUsernameResponse, error) {
	exists, err := service.tblUser.CheckUserExists(ctx, "username", username)
	if err != nil {
		return &pb.CheckUsernameResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while checking username",
			},
		}, err
	}

	return &pb.CheckUsernameResponse{
		Base: &pb.BaseResponse{
			Code: ResponseCodeSuccess,
			Desc: "Success",
		},
		Exists: exists,
	}, nil
}

func (service *UserService) GetUser(ctx context.Context, username string) (*pb.GetUserResponse, error) {
	user, err := service.tblUser.GetUser(ctx, "username", username)
	if err != nil {
		if errors.Is(err, constant.ErrDatabaseNotFound) {
			return &pb.GetUserResponse{
				Base: &pb.BaseResponse{
					Code: ResponseCodeNotFound,
					Desc: "User not found",
				},
			}, err
		}
		return &pb.GetUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while getting user",
			},
		}, err
	}

	return &pb.GetUserResponse{
		Base: &pb.BaseResponse{
			Code: ResponseCodeSuccess,
			Desc: "Success",
		},
		User: &pb.UserData{
			Id:          int32(user.Id),
			Username:    user.Username,
			Coin:        int32(user.Coin),
			CreatedDate: timestamppb.New(user.CreatedDate),
			UpdatedDate: timestamppb.New(user.UpdatedDate),
		},
	}, nil
}

func (service *UserService) CreateUser(ctx context.Context, username, password string) (*pb.CreateUserResponse, error) {
	// Check if username exists
	exists, err := service.CheckUsername(ctx, username)
	if err != nil {
		return &pb.CreateUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while checking username",
			},
		}, err
	}
	if exists.Exists {
		return &pb.CreateUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeValidation,
				Desc: "Username already exists",
			},
		}, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.CreateUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeValidation,
				Desc: "Error hashing password",
			},
		}, err
	}

	// Create user
	user := &entity.TblUser{
		Username: username,
		Password: string(hashedPassword),
		Coin:     0,
	}

	id, err := service.tblUser.InsertUser(ctx, *user)
	if err != nil {
		return &pb.CreateUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while creating user",
			},
		}, err
	}

	// Get the created user
	createdUser, err := service.tblUser.GetUser(ctx, "id", strconv.Itoa(id))
	if err != nil {
		return &pb.CreateUserResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while getting created user",
			},
		}, err
	}

	return &pb.CreateUserResponse{
		Base: &pb.BaseResponse{
			Code: ResponseCodeSuccess,
			Desc: "Success",
		},
		User: &pb.UserData{
			Id:          int32(createdUser.Id),
			Username:    createdUser.Username,
			Coin:        int32(createdUser.Coin),
			CreatedDate: timestamppb.New(createdUser.CreatedDate),
			UpdatedDate: timestamppb.New(createdUser.UpdatedDate),
		},
	}, nil
}

func (service *UserService) PaginateUsers(ctx context.Context, req *pb.PaginateRequest) (*pb.PaginateUsersResponse, error) {
	users, total, err := service.tblUser.PaginateUser(ctx, req)
	if err != nil {
		return &pb.PaginateUsersResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeDBError,
				Desc: "Database error while paginating users",
			},
		}, err
	}

	if len(users) == 0 {
		return &pb.PaginateUsersResponse{
			Base: &pb.BaseResponse{
				Code: ResponseCodeNotFound,
				Desc: "No users found",
			},
		}, errors.New("no users found")
	}

	usersResponse := make([]*pb.UserData, len(users))
	for i, user := range users {
		usersResponse[i] = &pb.UserData{
			Id:          int32(user.Id),
			Username:    user.Username,
			Coin:        int32(user.Coin),
			CreatedDate: timestamppb.New(user.CreatedDate),
			UpdatedDate: timestamppb.New(user.UpdatedDate),
		}
	}

	totalPages := int32(math.Ceil(float64(total) / float64(req.Limit)))

	return &pb.PaginateUsersResponse{
		Base: &pb.BaseResponse{
			Code: ResponseCodeSuccess,
			Desc: "Success",
		},
		Users:      usersResponse,
		Total:      int32(total),
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}, nil
}
