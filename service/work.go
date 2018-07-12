package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA/model"
)

type Question struct {
	BookId     int
	QuestionNo int
}

// ユーザーごとの回答情報取得
func ExisByRecord(name string) ([]model.Record, bool) {
	var records []model.Record
	db.Where("name = ?", name).Find(&records)
	return records, len(records) != 0
}

// 子どもごとの回答情報取得
func GetByRecordFromChild(name string, childId int) ([]model.Record, bool) {
	var records []model.Record
	db.Where("name = ? and child_id = ?", name, childId).Find(&records)
	return records, len(records) != 0
}

// 同一回答日の回答情報を取得
func GetByRecordFromDay(name string, childId int, day time.Time) ([]model.Record, bool) {
	var records []model.Record
	date := day.String()
	db.Where("name = ? and child_id = ? and answer_day = ?", name, childId, date).Find(&records)
	return records, len(records) != 0
}

// 同一ジャンルの回答情報を取得
func GetByRecordFromGenre(name string, childId int, genreId int) ([]model.Record, bool) {
	var records []model.Record
	db.Where("name = ? and child_id = ? and genre_id = ?", name, childId, genreId).Find(&records)
	return records, len(records) != 0
}

// 同一回答日&同一ジャンルの回答情報を取得
func GetByRecordFromGenreDate(name string, childId int, day time.Time, genreId int) ([]model.Record, bool) {
	var records []model.Record
	date := day.String()
	db.Where("name = ? and child_id = ? and answer_day = ? and genre_id = ?", name, childId, date, genreId).Find(&records)
	return records, len(records) != 0
}

// 問題番号から答えを取得
func GetByCorrect(bookId int, questionNo int) string {
	var question []model.Question
	err := db.Where("book_id = ? and question_no = ?", bookId, questionNo).Find(&question).Error

	if err != nil {
		return ""
	}
	return question[0].Correct
}

// BookIdの問題数を取得
func GetByQuestion(genreId int) int {
	var question []model.Question
	db.Where("genre_id = ?", genreId).Find(&question)

	return len(question)
}

// ジャンル名称の取得
func GetGenreName(bookId int) string {
	book := GetBookData(bookId)
	if book == nil {
		return ""
	}
	genre := GetGenreData(book[0].GenreId)
	if genre == nil {
		return ""
	}

	return genre[0].GenreName
}

// 本情報の取得
func GetBookData(bookId int) []model.Book {
	var book []model.Book
	err := db.Where("book_id = ?", bookId).Find(&book).Error
	if err != nil {
		return nil
	}
	return book
}

// ジャンル情報の取得
func GetGenreData(genreId int) []model.Genre {
	var genre []model.Genre
	err := db.Where("genre_id = ?", genreId).Find(&genre).Error
	if err != nil {
		return nil
	}
	return genre
}

// ジャンルの個数取得
func GetGenreNumber() int {
	var genre []model.Genre
	db.Where("").Find(&genre)
	return len(genre)
}

// 問題情報の取得
func GetQuestionData(bookId int, qNo int) []model.Question {
	var question []model.Question
	err := db.Where("book_id = ? and question_no = ?", bookId, qNo).Find(&question).Error
	if err != nil {
		return nil
	}
	return question
}

// タグ情報の取得(タグIDから)
func GetTagDataFromTagId(tagId string) *model.Tag {
	var tag model.Tag
	err := db.Where("tag_id = ?", tagId).First(&tag).Error
	if err != nil {
		return nil
	}
	return &tag
}

// タグ情報の取得(bookId&questionNoから)
func GetTagDataFromBookId(bookId int, questionId int) []model.Tag {
	var tag []model.Tag
	err := db.Where("book_id = ? and question_no = ?", bookId, questionId).Find(&tag).Error
	if err != nil {
		return nil
	}
	return tag
}

// タグ情報の取得(uuidから)
func GetTagDataFromUuid(uuid string) *model.Tag {
	var tag model.Tag
	err := db.Where("uuid = ?", uuid).Find(&tag).Error
	if err != nil {
		return nil
	}
	return &tag
}

// 回答情報の削除
func DeleteUserAnswer(deviceId string) bool {
	var records []model.Record
	err := db.Where("device_id = ?", deviceId).First(&records).Error
	if err != nil {
		return false
	}
	db.Delete(records)
	return true
}


// 回答数&正解数カウント
func CountAnswerNum(record []model.Record) (int, int) {
	var already []Question
	var now Question
	var correctNum int
	var i int
	
	recordNum := len(record)
	answerNum := len(record)
	for i = recordNum; i > 0; i-- {
		now.BookId = record[i - 1].BookId
		now.QuestionNo = record[i - 1].QuestionNo
		// 1問目の時
		if i == len(record) {
			if record[i - 1].Correct {
				correctNum++
			}
			already = append(already, now)
		} else {
			if judgeQuestion(record[i - 1].BookId, record[i - 1].QuestionNo, already) {
				answerNum--
			} else {
				if record[i - 1].Correct {
					correctNum++
				}
				already = append(already, now)
			}
		}
		now = Question{}
	}
	return answerNum, correctNum
}

// 同一問題を回答済みか？
func judgeQuestion(bookId int, questionNo int, already []Question) bool {
	for i := 0; i < len(already); i++ {
		// 同一問題を回答済みの時
		if (bookId == already[i].BookId) && (questionNo == already[i].QuestionNo) {
			return true
		}
	}
	return false
}