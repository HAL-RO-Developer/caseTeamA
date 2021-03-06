package validation

import (
	"fmt"

	"github.com/HAL-RO-Developer/caseTeamA/controller/response"
	"github.com/gin-gonic/gin"
)

type Question struct {
	BookId     int       `json:"book_id"`
	QuestionNo int       `json:"question_no"`
	Sentence   []TagInfo `json:"sentence"`
	Answer     []TagInfo `json:"answer"`
	Correct    string    `json:"correct"`
	GenreId    int       `json:"genre_id"`
}

type TagInfo struct {
	TagId    string `json:"tag_id"`
	Text     string `json:"text"`
	Phonetic string `json:"phonetic"`
}

type GenreInfo struct {
	Genre string `json:"genre_name"`
}

func QuestionValidation(c *gin.Context) (Question, bool) {
	var req Question
	err := c.BindJSON(&req)
	fmt.Println(req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}

func GenreValidation(c *gin.Context) (GenreInfo, bool) {
	var req GenreInfo
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}
