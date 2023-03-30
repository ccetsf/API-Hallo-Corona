package handlers

import (
	result_dto "hallo-corona/dto/result"
	user_dto "hallo-corona/dto/user"
	"hallo-corona/models"
	"hallo-corona/pkg/bcrypt"
	jwtToken "hallo-corona/pkg/jwt"
	repo "hallo-corona/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type handlerAuth struct {
	AuthRepository repo.AuthRepository
}

func HandlerAuth(AuthRepository repo.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(ctx echo.Context) error {

	//create new register request
	request := new(user_dto.RegisterRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Failed binding request DTO", Error: err.Error()})
	}

	//create validation struct
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Failed validation request DTO", Error: err.Error()})
	}

	//hashing password
	passwordHash, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Failed to hashing password", Error: err.Error()})
	}

	//create new data user from request to model user
	dataUser := models.User{
		FullName: request.FullName,
		Username: request.Username,
		Email:    request.Email,
		Password: passwordHash,
		Role:     request.Role,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Photo:    "http://localhost:5000/uploads/profile.png",
	}

	//register dataUser
	newUser, err := h.AuthRepository.Register(dataUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed to register new account", Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Register Success", Data: newUser})

}

func (h *handlerAuth) Login(ctx echo.Context) error {

	//create new login request
	request := new(user_dto.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Failed binding request DTO", Error: err.Error()})
	}

	//create new data login from request to model user
	data := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	//login data user
	userLogin, err := h.AuthRepository.Login(data.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Incorrect username", Error: err.Error()})
	}

	//check password
	isValid := bcrypt.CheckPasswordHash(request.Password, userLogin.Password)
	if !isValid {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Incorrect password", Error: err.Error()})
	}

	//set jwt map
	claims := jwt.MapClaims{}
	claims["id"] = userLogin.Id
	claims["role"] = userLogin.Role
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix() // 4 hours expired

	//generate token
	token, generateTokenErr := jwtToken.GenerateToken(&claims)
	if generateTokenErr != nil {
		log.Println(generateTokenErr)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	//create login response
	loginResponse := user_dto.LoginResponse{
		FullName: userLogin.FullName,
		Username: userLogin.Username,
		Email:    userLogin.Email,
		Role:     userLogin.Role,
		Token:    token,
	}

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Login Success", Data: loginResponse})

}

func (h *handlerAuth) CheckAuth(ctx echo.Context) error {
	user := ctx.Get("userLogin")
	userId := user.(jwt.MapClaims)["id"].(float64)

	userLogin, _ := h.AuthRepository.CheckAuth(int(userId))

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "CheckAuth Success", Data: convertResponseCheckAuth(userLogin)})
}

func convertResponseCheckAuth(u models.User) models.CheckAuthResponse {
	return models.CheckAuthResponse{
		FullName: u.FullName,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}
