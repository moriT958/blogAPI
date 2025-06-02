package services

import (
	"fmt"

	"github.com/moriT958/go-api/models"
)

func (s *BlogService) PostComment(comment models.Comment) (models.Comment, error) {
	comment, err := s.Cr.Create(comment)
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to record comment: %w", err)
	}

	return comment, nil
}
