package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"user_service/internal/hash"
	"user_service/internal/sms"
	tokens "user_service/internal/token"
	"user_service/models"
	"user_service/proto"
	"user_service/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	DB *mongo.Collection
	LG *log.Logger
}

func NewUserServer(db *mongo.Collection, lg *log.Logger) *Server {
	return &Server{DB: db, LG: lg}
}

func (s *Server) RegisterUser(ctx context.Context, req *proto.User) (*proto.UserResponse, error) {
	// Email to'g'riligini tekshirish
	if !sms.IsValidEmail(req.Email) {
		s.LG.Println("RegisterUser ERROR: Invalid email address")
		log.Println("Invalid email address")
		return nil, errors.New("ERROR: invalid email address")
	}

	// Parolni xesh qilish
	pass, err := hash.HashPassword(req.PasswordHash)
	if err != nil {
		s.LG.Printf("RegisterUser ERROR: %v", err)
		log.Println(err)
	}

	// id, _ := primitive.ObjectIDFromHex(req.UserId)

	user := &models.User{
		UserId:       primitive.NewObjectID(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: pass,
		Profile: models.Profile{
			Name:    req.Profile.Name,
			Address: req.Profile.Address,
		},
	}

	// Foydalanuvchini ma'lumotlar bazasiga qo'shish
	err = repo.AddUserToDB(user)
	if err != nil {
		s.LG.Printf("RegisterUser ERROR: failed to add user to DB: %v", err)
		return nil, fmt.Errorf("ERROR: failed to add user to DB: %v", err)
	}

	// Foydalanuvchini Redis'ga saqlash
	err = repo.CacheUser(user)
	if err != nil {
		s.LG.Printf("RegisterUser ERROR: failed to cache user: %v", err)
		return nil, fmt.Errorf("ERROR: failed to cache user: %v", err)
	}

	response := &proto.UserResponse{User: &proto.User{
		UserId:       user.UserId.Hex(),
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Profile: &proto.Profile{
			Name:    user.Profile.Name,
			Address: user.Profile.Address,
		},
	}}

	s.LG.Println("LoginUser INFO: Succecc returned")
	log.Println("Successfully registere user")
	return response, nil
}

func (s *Server) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	email := req.Email
	password := req.Password

	// MongoDB'dan foydalanuvchi ma'lumotlarini olish
	user, err := repo.GetUserFromDB(email)
	if err != nil {
		s.LG.Printf("RegisterUser ERROR: failed to get user from DB: %v", err)
		return nil, fmt.Errorf("ERROR: failed to get user from DB: %v", err)
	}
	// Parolni tekshirish
	if !hash.CheckPasswordHash(password, user.PasswordHash) {
		s.LG.Println("RegisterUser ERROR: invalid password")
		return nil, fmt.Errorf("ERROR: invalid password")
	}

	// JWT Token yaratish
	token, err := tokens.CreateToken(user.UserId)
	if err != nil {
		s.LG.Printf("RegisterUser ERROR: failed to generate token: %v", err)
		return nil, fmt.Errorf("ERROR: failed to generate token: %v", err)
	}

	jwttoken := &proto.LoginResponse{
		UserId: user.UserId.Hex(),
		Token:  token,
	}

	s.LG.Println("LoginUser INFO: Succecc returned")
	log.Println("Successfully login user")
	return jwttoken, nil
}

func (s *Server) GetUserProfile(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	userid := req.UserId
	token := req.Token

	if repo.ValidateToken(token) {
		s.LG.Println("GetUserProfile ERROR: invalid token")
		return nil, fmt.Errorf("ERROR: invalid token")
	}

	// Foydalanuvchi IDsi orqali topish
	user, err := repo.GetUserWithID(userid)
	if err != nil {
		s.LG.Printf("GetUserProfile ERROR: error: %v", err)
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	userResponse := &proto.UserResponse{
		User: &proto.User{
			UserId:       user.UserId,
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Profile: &proto.Profile{
				Name:    user.Profile.Name,
				Address: user.Profile.Address,
			},
		},
	}

	s.LG.Println("GetUserProfile INFO: Succecc returned")
	log.Print("Successfully get user profile")
	return userResponse, nil
}
