package article_dto

type CreateArticleRequest struct {
	UserId      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Attache     string `json:"attache" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateArticleRequest struct {
	Title       string `json:"title" validate:"required"`
	Attache     string `json:"attache" validate:"required"`
	Description string `json:"description" validate:"required"`
}
