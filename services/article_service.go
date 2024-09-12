package services

import (
	"database/sql"
	"errors"

	"github.com/moriT958/go-api/apperrors"
	"github.com/moriT958/go-api/models"
	"github.com/moriT958/go-api/repositories"
)

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	// 独自エラーを起点としてラップする。
	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// チャネルで送受信できるデータは一つのみなので構造体としてまとめる
	// articleを受信するチャネル
	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)
	go func(ch chan<- articleResult) {
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan)

	// commentを受信するチャネル
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)
	go func(ch chan<- commentResult) {
		commentList, err := repositories.SelectCommentList(s.db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan)

	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// チャネル2つ分の走査が必要なため
	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		// sqlドライバ側のエラーなのかScan関数で起きたエラーなのか区別するため、条件分岐する。
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "No data found")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to load data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		if errors.Is(commentGetErr, sql.ErrNoRows) { // sqlはレコードが見つからない場合ErrNoRowsを返す。
			err := apperrors.NAData.Wrap(commentGetErr, "No data found")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to load data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
