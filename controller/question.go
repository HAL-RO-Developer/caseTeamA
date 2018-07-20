package controller

import (
	"github.com/HAL-RO-Developer/caseTeamA/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA/service"
	"github.com/gin-gonic/gin"
)

var Question = questionimpl{}

type questionimpl struct {
}

type GenreInfo struct {
	GenreId   int    `json:"genre_id"`
	GenreName string `json:"genre_name"`
}

func (q *questionimpl) CreateQuestion(c *gin.Context) {
	req, success := validation.QuestionValidation(c)
	if !success {
		return
	}

	genreData := service.GetGenreData(req.GenreId)
	if len(genreData) == 0 {
		response.BadRequest(gin.H{"error": "そのジャンルIDは存在しません。"}, c)
		return
	}
	findBook := service.GetBookInfo(req.BookId, req.GenreId)
	findQuestion := service.GetQuestionInfo(req.BookId, req.QuestionNo)
	if findQuestion {
		response.BadRequest(gin.H{"error": "その問題番号は登録済みです。"}, c)
		return
	}

	findTagId := service.GetSentenceInfo(req.Sentence[0].TagId)
	if findTagId {
		response.BadRequest(gin.H{"error": "そのタグIDは登録済みです。"}, c)
		return
	}

	for i := 0; i < len(req.Answer); i++ {
		if service.GetSentenceInfo(req.Answer[i].TagId) {
			response.BadRequest(gin.H{"error": "回答登録に失敗しました。"}, c)
			return
		}
	}

	if !findBook {
		service.BookGenerate(req.BookId, req.GenreId)
	}
	service.QuestionGenerate(req.BookId, req.QuestionNo, req.GenreId, req.Sentence[0].TagId, req.Correct)
	service.SentenceGenerate(req.Sentence[0].TagId, req.BookId, req.QuestionNo, req.Sentence[0].Text, req.Sentence[0].Phonetic)
	service.CorrectGenerate(req)
	response.Json(gin.H{"success": "問題情報を登録しました。"}, c)
}

// ジャンル作成
func (q *questionimpl) CreateGenre(c *gin.Context) {
	req, success := validation.GenreValidation(c)
	if !success {
		return
	}
	success = service.GenreGenerate(req.Genre)
	if !success {
		response.BadRequest(gin.H{"error": "ジャンル登録に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "ジャンル情報を登録しました。"}, c)
}

// ジャンル情報取得
func (q *questionimpl) GetGenre(c *gin.Context) {
	var buf GenreInfo
	var res []GenreInfo

	genres := service.GetGenreNum()
	for i := 0; i < len(genres); i++ {
		buf.GenreId = genres[i].GenreId
		buf.GenreName = genres[i].GenreName
		res = append(res, buf)
		buf = GenreInfo{}
	}

	response.Json(gin.H{"genre": res}, c)
}
