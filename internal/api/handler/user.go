package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hanzala211/instagram/internal/api/dto"
	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/internal/cache"
	"github.com/hanzala211/instagram/internal/services"
	"github.com/hanzala211/instagram/utils"
)

type UserHandler struct {
	userService *services.UserService
	rdRepo *cache.RedisRepo
}

var validate = validator.New()

func NewUserHandler(userService *services.UserService, rdRepo *cache.RedisRepo) *UserHandler {
	return &UserHandler{
		userService: userService,
		rdRepo: rdRepo,
	}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	req := &dto.SignupRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
	}
	err = h.userService.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return	
	}
	token, err := utils.CreateToken(user.ID, utils.GetEnv("JWT_SECRET", "secret"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmtKey := fmt.Sprintf("user-%s", user.ID)
	jso, err := json.Marshal(&user)
	err = h.rdRepo.Set(fmtKey, jso, time.Hour * 24)
	if err != nil {
		fmt.Println(err)
	}
	utils.WriteResponse(w, 200, map[string]any{"token": token, "user": user})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &dto.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validate.Struct(req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	existingUser, err := h.userService.Login(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := utils.CreateToken(existingUser.ID, utils.GetEnv("JWT_SECRET", "secret"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteResponse(w, 200, map[string]any{"token": token, "user": existingUser})
}

func (h *UserHandler) ME(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	utils.WriteResponse(w, 200, map[string]any{"user": user})
}