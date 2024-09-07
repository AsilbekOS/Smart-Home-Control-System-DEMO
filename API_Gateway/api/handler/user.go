package handler

import (
	proto "api_gateway/proto/user"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
)

type UserClient struct {
	Client proto.UserServiceClient
	LG     *log.Logger
}

func NewUserClient(cl proto.UserServiceClient, lg *log.Logger) *UserClient {
	return &UserClient{Client: cl, LG: lg}
}

// @Router		/user/register [post]
// @Summary		CREATE User
// @Description Bu endpoint yangi foydalanuvchini yaratish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		User
// @Accept		json
// @Produce 	json
// @Param		body body proto.UserSwag true "UserSwag"
// @Success 	201 {object} map[string]interface{} "User successfully created"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (u *UserClient) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Register ichiga kirdi>>>>>>>>>>>>>")
	var userReq proto.User

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		u.LG.Printf("ERROR: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protojson.Unmarshal(bytes, &userReq)

	resp, err := u.Client.RegisterUser(r.Context(), &userReq)
	if err != nil {
		u.LG.Printf("ERROR: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		u.LG.Println("ERROR: Ma'lumotni encode qilishda xatolik...")
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}

// @Router		/user/login [post]
// @Summary		LOGIN User
// @Description Bu endpoint foydalanuvchi profilini olish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		User
// @Accept		json
// @Produce 	json
// @Param		body body proto.LoginRequest true "LoginRequest"
// @Success 	200 {object} map[string]interface{} "Login successful"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (u *UserClient) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userReq proto.LoginRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		u.LG.Println("ERROR: POST: bodydan o'qib olishda xatolik...")
		http.Error(w, "POST: bodydan o'qib olishda xatolik...", http.StatusBadRequest)
		return
	}

	err = protojson.Unmarshal(bytes, &userReq)
	if err != nil {
		u.LG.Println("ERROR: POST: unmarshal qilishda xatolik...")
		http.Error(w, "POST: unmarshal qilishda xatolik...", http.StatusBadRequest)
		return
	}

	resp, err := u.Client.LoginUser(r.Context(), &userReq)
	if err != nil {
		u.LG.Println("ERROR: POST: Serverdan ma'lumot olishda xatolik...")
		http.Error(w, "POST: Serverdan ma'lumot olishda xatolik...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		u.LG.Println("ERROR: Ma'lumotni encode qilishda xatolik...")
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}

// @Router		/user/profile [get]
// @Summary		PROFILE User
// @Description Bu endpoint foydalanuvchi profilini olish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		User
// @Accept		json
// @Produce 	json
// @Param 		token query string true "Token: "
// @Param 		user_id query string true "User ID: "
// @Success 	200 {object} map[string]interface{} "Successful response"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (u *UserClient) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	user_id := r.URL.Query().Get("user_id")

	if token == "" {
		u.LG.Println("Token kiritishingiz lozim")
		http.Error(w, "Token kiritishingiz lozim", http.StatusBadRequest)
		return
	}

	if user_id == "" {
		u.LG.Println("User ID kiritishingiz lozim")
		http.Error(w, "User ID kiritishingiz lozim", http.StatusBadRequest)
		return
	}

	userReq := proto.UserRequest{
		UserId: user_id,
	}

	resp, err := u.Client.GetUserProfile(r.Context(), &userReq)
	if err != nil {
		u.LG.Println("ERROR: GET: Serverdan ma'lumot olishda xatolik...")
		http.Error(w, "GET: Serverdan ma'lumot olishda xatolik...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		u.LG.Println("ERROR: Ma'lumotni encode qilishda xatolik...")
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}
