package service

import (
	"belajar-auth/domain"
	"belajar-auth/dto"
	"belajar-auth/internal/util"
	"context"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Authenticate implements domain.UserService.
func (u *UserService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	token := util.GeneratorRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)

	return dto.AuthRes{
		Token: token,
	}, nil
}

// ValidateToken implements domain.UserService.
func (u *UserService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}
