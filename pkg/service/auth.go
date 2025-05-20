package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-demo/pkg/repository"
	"go-demo/types"
	"time"
)

const salt = "dsafljhh1231fsdjklhah"
const secret = "dsa#41fljhh1231fsdjkl@hAAHh"
const tokenTTL = time.Hour * 12

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthServiceImpl struct {
	repo repository.AuthorizationRepository
}

func NewAuthService(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) CreateUser(user types.User) (int, error) {
	user.Password = s.hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthServiceImpl) SignIn(username string, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	if user.Password != s.hashPassword(password) {
		return "", errors.New("incorrect username or password")
	}

	return s.generateToken(user.Id)
}

func (s *AuthServiceImpl) generateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(secret))
}

func (s *AuthServiceImpl) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
