package services

import (
	"github.com/moriT958/go-api/models"
	"github.com/moriT958/go-api/repositories"
)

var _ IBlogService = (*BlogService)(nil)

type BlogService struct {
	Ar repositories.IArticleRepository
	Cr repositories.ICommentRepositoy
}

func NewBlogService(ar repositories.IArticleRepository, cr repositories.ICommentRepositoy) *BlogService {
	return &BlogService{
		Ar: ar,
		Cr: cr,
	}
}

type IBlogService interface {
	PostComment(models.Comment) (models.Comment, error)
	PostArticle(models.Article) (models.Article, error)
	GetArticles(int) ([]models.Article, error)
	GetArticle(int) (models.Article, error)
	AddNice(int) (models.Article, error)
}
