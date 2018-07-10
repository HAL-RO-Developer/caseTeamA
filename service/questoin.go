package service

import (
	"github.com/HAL-RO-Developer/caseTeamA/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA/model"
)

// ドリル情報作成
func BookGenerate(bookId int, genreId int) {
	book := model.Book{
		BookId:  bookId,
		GenreId: genreId,
	}
	db.Create(&book)
}

func GetBookInfo(bookId int, genreId int) bool {
	var books []model.Book
	db.Where("book_id = ? and genre_id = ?", bookId, genreId).Find(&books)
	return len(books) != 0
}

// 問題作成
func QuestionGenerate(bookId int, questionNo int, sentece string, correct string) {
	question := model.Question{
		BookId:     bookId,
		QuestionNo: questionNo,
		Sentence:   sentece,
		Correct:    correct,
	}

	db.Create(&question)
}

func GetQuestionInfo(bookId int, questionNo int) bool {
	var questions []model.Question
	db.Where("book_id = ? and question_no = ?", bookId, questionNo).Find(&questions)
	return len(questions) != 0
}

// 問題タグ
func SentenceGenerate(tagId string, uuid string, bookId int, questionNo int, sentence string) {
	tag := model.Tag{
		TagId:      tagId,      // 任意のタグ番号
		Uuid:       "",         // tagのuuid
		BookId:     bookId,     // ドリルId
		QuestionNo: questionNo, // ドリルの問題番号
		Sentence:   sentence,   // 問題文
	}

	db.Create(&tag) // カラム作成
}

func GetSentenceInfo(tagId string) bool {
	var questions []model.Tag
	db.Where("tag_id = ?", tagId).Find(&questions)
	return len(questions) != 0
}

// 回答タグ
func CorrectGenerate(tagInfo validation.Question) {
	var tag model.Tag
	for i := 0; i < len(tagInfo.Answer); i++ {
		tag = model.Tag{
			TagId:      tagInfo.Answer[i].TagId, // 任意のタグ番号
			Uuid:       "",                  // タグのuuid
			BookId:     tagInfo.BookId,          // ドリルid
			QuestionNo: tagInfo.QuestionNo,      // ドリルの問題番号
			Answer:     tagInfo.Answer[i].Text,  // タグの回答
		}
		db.Create(&tag)
		tag = model.Tag{}
	}
}

// ジャンル作成
func GenreGenerate(genreName string) bool {
	if getGenreInfo(genreName) {
		return false
	}
	genre := model.Genre{
		GenreId:   len(GetGenreNum()) + 1,
		GenreName: genreName,
	}

	db.Create(&genre)
	return true
}

// 本ジャンル取得
func getGenreInfo(genre string) bool {
	var genres []model.Genre
	db.Where("genre_name = ?", genre).Find(&genres)
	return len(genres) != 0
}

func GetGenreNum() []model.Genre {
	var genres []model.Genre
	db.Where("").Find(&genres)
	return genres
}
