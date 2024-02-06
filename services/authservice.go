package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jn0x/reddigo/domain"
	"github.com/jn0x/reddigo/http/requests"
	"github.com/jn0x/reddigo/repositories"
	"github.com/jn0x/reddigo/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(u requests.AuthReq) []error
	Login(u requests.AuthReq) (string, []error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() AuthService {
	return AuthService(&authService{userRepo: repositories.NewUserRepository()})
}

func (s *authService) Signup(u requests.AuthReq) []error {
	if err := utils.ValidateAuth(u.Username, u.Password); err != nil {
		return err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	user := &domain.User{Username: u.Username, Password: string(pass)}
	err_ := s.userRepo.CreateUser(user)
	if err_ != nil {
		var errs []error
		errs = append(errs, err_)
		return errs

	}
	return nil
}

func (s *authService) Login(u requests.AuthReq) (string, []error) {
	user := s.userRepo.GetUserByUsername(u.Username)
	if user.Username == "" {
		var errs []error
		errs = append(errs, errors.New("invalid credentials"))
		return "", errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {

		var errs []error
		errs = append(errs, errors.New("invalid credentials"))
		return "", errs

	}
	token, err := GenerateToken(&user)
	if err != nil {
		var errs []error
		errs = append(errs, errors.New("Something wrong happened"))

		return "", errs
	}

	return token, nil
}

func GenerateToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("this is the secret key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("this is the secret key"), nil
	})
	if err != nil {
		return nil
	}
	if !token.Valid {
		return nil
	}

	return token
}
