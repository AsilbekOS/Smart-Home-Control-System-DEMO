package repo

import (
	"context"
	"fmt"
	"log"
	"time"
	mongodb "user_service/database/MongoDB"
	redisDB "user_service/database/Redis"
	"user_service/models"
	"user_service/proto"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToDB(user *models.User) error {
	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}

	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "user_id", Value: user.UserId},
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password_hash", Value: user.PasswordHash},
		{Key: "profile", Value: bson.D{
			{Key: "name", Value: user.Profile.Name},
			{Key: "address", Value: user.Profile.Address},
		}},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// _, err = collection.InsertOne(context.Background(), user)
	return nil
}

// Foydalanuvchini Redis'ga saqlash
func CacheUser(user *models.User) error {
	redisClient, err := redisDB.ConnectRedis()
	if err != nil {
		log.Println(err)
	}

	_, err = redisClient.HSet(context.Background(), "user:"+user.Email, map[string]interface{}{
		"user_id":         user.UserId.Hex(),
		"username":        user.Username,
		"email":           user.Email,
		"password_hash":   user.PasswordHash,
		"profile_name":    user.Profile.Name,
		"profile_address": user.Profile.Address,
	}).Result()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Verifikatsiya codini redisga saqlash
func AddRedisVercy(code string) string {
	redisClient, err := redisDB.ConnectRedis()
	if err != nil {
		log.Println(err)
	}

	err = redisClient.Set(context.Background(), "VerifyCode", code, 3*time.Minute).Err()
	if err != nil {
		log.Fatalf("Error saving to Redis: %v", err)
	}

	fmt.Println("Data saved to Redis successfully.")

	val, err := redisClient.Get(context.Background(), "VerifyCode").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Kalayotgan kalit mavjud emas.")
		} else {
			log.Fatalf("Redisdan olishda xatolik: %v", err)
		}
	}

	return val
}

// Foydalanuvchi ma'lumotlarini MongoDB'dan olish
func GetUserFromDB(email string) (*models.User, error) {
	// fmt.Println("Mongodan ma'lumotni oldi.")
	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}
	var user *models.User
	err = collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Foydalanuvchi ma'lumotlarini MongoDB'dan olish
func GetUserWithID(id string) (*proto.User, error) {
	// fmt.Println("Mongodan ma'lumotni oldi.")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error changing id type: %v", err)
	}

	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}
	var user *proto.User
	err = collection.FindOne(context.Background(), bson.D{{Key: "user_id", Value: userId}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Tokenni tekshirish funksiyasi
var secretKey = []byte("Asilbek")

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}
