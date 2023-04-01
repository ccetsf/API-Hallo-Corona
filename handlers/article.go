package handlers

import (
	article_dto "hallo-corona/dto/article"
	result_dto "hallo-corona/dto/result"
	"hallo-corona/models"
	repo "hallo-corona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

// create url server for path files
var path_file = "http://localhost:5000/"

type handlerArticle struct {
	ArticleRepository repo.ArticleRepository
}

func HandlerArticle(ArticleRepository repo.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

// method create article
func (h *handlerArticle) CreateArticle(ctx echo.Context) error {

	//get user role
	userLogin := ctx.Get("userLogin")
	userRole := userLogin.(jwt.MapClaims)["role"].(string)

	//check user access
	if userRole != "Doctor" {
		// return error access denied
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "access denied", Error: "Error: Access denied"})
	}

	// get user id
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// get datafile
	dataFile := ctx.Get("dataFile").(string)

	// create new request object article
	request := article_dto.CreateArticleRequest{
		UserId:      int(userId),
		Title:       ctx.FormValue("title"),
		Attache:     dataFile,
		Description: ctx.FormValue("description"),
	}

	// validate the request object
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Validasi form error", Error: err.Error()})
	}

	// create new model article object
	article := models.Article{
		UserId:      request.UserId,
		Title:       request.Title,
		Attache:     path_file + request.Attache,
		Description: request.Description,
	}

	// store data to database
	article, err = h.ArticleRepository.CreateArticle(article)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Gagal menambahkan data", Error: err.Error()})
	}

	// get article by id
	article, _ = h.ArticleRepository.GetArticle(article.Id)

	// return success result
	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Article berhasil ditambahkan", Data: article})

}

// method get article by id
func (h *handlerArticle) GetArticle(ctx echo.Context) error {

	//get article id
	id, _ := strconv.Atoi(ctx.Param("id"))

	//get article by id
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Article tidak ditemukan", Error: err.Error()})
	}

	//return success result
	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Article ditemukan", Data: article})
}

// method find all article
func (h *handlerArticle) FindArticles(ctx echo.Context) error {

	//find all articles
	articles, err := h.ArticleRepository.FindArticles()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Gagal mendapatkan semua article", Error: err.Error()})
	}

	if len(articles) <= 0 {
		return ctx.JSON(http.StatusOK, result_dto.ErrorResult{Code: http.StatusOK, Message: "Record not found", Error: ""})
	}

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Semua article berhasil didapatkan", Data: articles})
}

// method delete article
func (h *handlerArticle) DeleteArticle(ctx echo.Context) error {

	//get user role
	userLogin := ctx.Get("userLogin")
	userRole := userLogin.(jwt.MapClaims)["role"].(string)

	//check user access
	if userRole != "Doctor" {
		//return error access denied
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Access denied", Error: "Error: Access denied"})
	}

	// get article id from url parameter
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get article by id
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "article not found", Error: err.Error()})
	}

	// check user access to article
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	if article.UserId != int(userId) {
		// return error access denied
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Access denied", Error: "Error: Access denied"})
	}

	//delete article
	articleDelete, err := h.ArticleRepository.DeleteArticle(article, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Delete article failed", Error: err.Error()})
	}

	//return result
	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Delete article succeeded", Data: articleDelete})

}

func (h *handlerArticle) MyArticles(ctx echo.Context) error {

	//get user id
	userLogin := ctx.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	//find all user articles
	articles, err := h.ArticleRepository.MyArticles(int(userId))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Article Not Found", Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Article didapatkan", Data: articles})
}

func (h *handlerArticle) UpdateArticle(ctx echo.Context) error {

	userLogin := ctx.Get("userLogin")
	userRole := userLogin.(jwt.MapClaims)["role"].(string)
	if userRole != "Doctor" {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Access denied", Error: "Access denied"})
	}

	dataFile := ctx.Get("dataFile").(string)

	request := article_dto.UpdateArticleRequest{
		Title:       ctx.FormValue("title"),
		Attache:     dataFile,
		Description: ctx.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Validation failed", Error: err.Error()})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Article Not Found", Error: err.Error()})
	}

	//check user access to article
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	if int(userId) != article.UserId {
		return ctx.JSON(http.StatusBadRequest, result_dto.ErrorResult{Code: http.StatusBadRequest, Message: "Access Denied", Error: "Access Denied"})
	}

	if request.Title != "" {
		article.Title = request.Title
	}

	if request.Attache != "" {
		article.Attache = path_file + request.Attache
	}

	if request.Description != "" {
		article.Description = request.Description
	}

	articleUpdated, err := h.ArticleRepository.UpdateArticle(article)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, result_dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Update Article Failed", Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, result_dto.SuccessResult{Code: http.StatusOK, Message: "Article successfully updated", Data: articleUpdated})

}
