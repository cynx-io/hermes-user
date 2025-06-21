package usermodule

import (
	"context"
	"errors"
	pb "github.com/cynxees/hermes-user/api/proto/gen/hermes"
	"github.com/cynxees/hermes-user/internal/constant"
	"github.com/cynxees/hermes-user/internal/model/entity"
	"github.com/cynxees/hermes-user/internal/model/response"
	"github.com/cynxees/hermes-user/internal/repository/database"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

const (
	ResponseCodeSuccess    = "00"
	ResponseCodeDBError    = "DB"
	ResponseCodeNotFound   = "NF"
	ResponseCodeValidation = "VD"
)

type UserService struct {
	tblUser *database.TblUser
}

func NewUserService(tblUser *database.TblUser) *UserService {
	return &UserService{
		tblUser: tblUser,
	}
}

func (service *UserService) CheckUsername(ctx context.Context, req *pb.UsernameRequest, resp *pb.CheckUsernameResponse) (err error) {
	exists, err := service.tblUser.CheckUserExists(ctx, "username", req.Username)
	if err != nil {
		response.ErrorDbUser(resp)
	}

	response.Success(resp)
	resp.Exists = exists
	return
}

func (service *UserService) GetUser(ctx context.Context, req *pb.UsernameRequest, resp *pb.GetUserResponse) (err error) {
	user, err := service.tblUser.GetUser(ctx, "username", req.Username)
	if err != nil {
		if errors.Is(err, constant.ErrDatabaseNotFound) {
			response.ErrorNotFound(resp)
			return
		}
		response.ErrorDbUser(resp)
		return
	}

	response.Success(resp)
	resp.User = user.ToUserResponse()
	return
}

func (service *UserService) CreateUser(ctx context.Context, req *pb.UsernamePasswordRequest, resp *pb.CreateUserResponse) (err error) {
	// Check if username exists
	exists, err := service.tblUser.CheckUserExists(ctx, "username", req.Username)
	if err != nil {
		response.ErrorDbUser(resp)
		return
	}
	if exists {
		response.ErrorNotAllowed(resp)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ErrorInternal(resp)
		return
	}

	// Create user
	user := &entity.TblUser{
		Username: req.Username,
		Password: string(hashedPassword),
		Coin:     0,
	}

	id, err := service.tblUser.InsertUser(ctx, *user)
	if err != nil {
		response.ErrorDbUser(resp)
		return
	}

	// Get the created user
	createdUser, err := service.tblUser.GetUser(ctx, "id", strconv.Itoa(id))
	if err != nil {
		response.ErrorDbUser(resp)
		return
	}

	response.Success(resp)
	resp.User = createdUser.ToUserResponse()
	return
}

func (service *UserService) PaginateUsers(ctx context.Context, req *pb.PaginateRequest, resp *pb.PaginateUsersResponse) (err error) {
	users, _, err := service.tblUser.PaginateUser(ctx, req)
	if err != nil {
		response.ErrorDbUser(resp)
		return
	}

	if len(users) == 0 {
		response.ErrorNotFound(resp)
		return
	}

	usersResponse := make([]*pb.User, len(users))
	for i, user := range users {
		usersResponse[i] = user.ToUserResponse()
	}

	response.Success(resp)
	resp.Users = usersResponse
	return
}

func (service *UserService) ValidatePassword(ctx context.Context, req *pb.UsernamePasswordRequest, resp *pb.ValidatePasswordResponse) (err error) {

	// Get user by username
	user, err := service.tblUser.GetUser(ctx, "username", req.Username)
	if err != nil {
		if errors.Is(err, constant.ErrDatabaseNotFound) {
			response.ErrorNotFound(resp)
			return
		}
		response.ErrorDbUser(resp)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		response.ErrorValidation(resp)
		resp.Base.Code = ResponseCodeValidation
		resp.Base.Desc = "Invalid password"
		return
	}

	response.Success(resp)
	resp.User = user.ToUserResponse()
	return
}
