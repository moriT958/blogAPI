package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/moriT958/go-api/models"
)

func (s *BlogService) PostArticle(article models.Article) (models.Article, error) {
	article, err := s.Ar.Create(article)
	if err != nil {
		return models.Article{}, fmt.Errorf("failed to create post: %w", err)
	}
	return article, nil
}

func (s *BlogService) GetArticles(page int) ([]models.Article, error) {
	ariticles, err := s.Ar.FindAll(page)
	if err != nil {
		return nil, fmt.Errorf("failed to get articles: %w", err)
	}

	return ariticles, nil
}

func (s *BlogService) GetArticle(id int) (models.Article, error) {

	// get article concurrently.
	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)
	go func(ch chan<- articleResult) {
		article, err := s.Ar.FindByID(id)
		ch <- articleResult{article: article, err: err}
	}(articleChan)

	// get comments concurrently.
	type commentResult struct {
		comments *[]models.Comment
		err      error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)
	go func(ch chan<- commentResult) {
		comments, err := s.Cr.FindAll(id)
		ch <- commentResult{comments: &comments, err: err}
	}(commentChan)

	var article models.Article
	var comments []models.Comment
	var articleErr, commentErr error

	// check two channels.
	for range 2 {
		select {
		case ar := <-articleChan:
			article, articleErr = ar.article, ar.err
		case cr := <-commentChan:
			comments, commentErr = *cr.comments, cr.err
		}
	}

	if articleErr != nil {
		if errors.Is(articleErr, sql.ErrNoRows) {
			return models.Article{}, fmt.Errorf("No data found: %w", articleErr)
		}
		return models.Article{}, fmt.Errorf("failed to load data: %w", articleErr)
	}

	if commentErr != nil {
		if errors.Is(commentErr, sql.ErrNoRows) {
			return models.Article{}, fmt.Errorf("No data found: %w", commentErr)
		}
		return models.Article{}, fmt.Errorf("failed to load data: %w", commentErr)
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil
}

func (s *BlogService) AddNice(articleId int) (models.Article, error) {
	if err := s.Ar.AddNice(articleId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Article{}, fmt.Errorf("does not exit target article: %w", err)
		}
		return models.Article{}, fmt.Errorf("failed to update nice count: %w", err)
	}

	article, err := s.Ar.FindByID(articleId)
	if err != nil {
		return models.Article{}, fmt.Errorf("failed to add nice num: %w", err)
	}

	return article, nil
}
