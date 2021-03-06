package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA/model"
	_ "github.com/satori/go.uuid"
)

// ユーザーの回答データ送信
func SendUserAnswer(deviceId string, tagUuid string, oldUuid string) (model.Record, int) {
	var record model.Record
	var oldTag *model.Tag
	deviceInfo, find := GetDeviceInfoFromDeviceId(deviceId)
	if !find {
		return model.Record{}, -1
	}

	if oldUuid != "" {
		oldTag = GetTagDataFromUuid(tagUuid)
		if oldTag == nil {
			return model.Record{}, -2
		}
	}

	tagInfo := GetTagDataFromUuid(tagUuid)
	if tagInfo == nil {
		return model.Record{}, -3
	}

	// 問題タグ時
	if tagInfo.Sentence != "" {
		if oldUuid != "" {
			if judgeSentenceTag(oldTag.Uuid) {
				tagInfo.Sentence = "前回の問題を回答してね"
			}
		}

		bocco, find := GetDeviceInfoFromDeviceId(deviceId)
		if !find {
			return model.Record{}, -4
		}
		TalkBocco(tagInfo.Sentence,bocco[0].Name)
		return record, 4
	} else {
		if oldUuid == "" {
			record = createRecord(deviceInfo[0].Name, deviceInfo[0].ChildId, tagInfo.BookId, tagInfo.QuestionNo, tagUuid)
			err := db.Create(&record).Error
			if err != nil {
				return model.Record{}, -6
			}
			if !record.Correct {
				return record, 2
			}
		} else if oldTag.Answer != "" || ((oldTag.BookId == tagInfo.BookId && oldTag.QuestionNo == tagInfo.QuestionNo) && oldTag.Answer == "") {
			record = createRecord(deviceInfo[0].Name, deviceInfo[0].ChildId, tagInfo.BookId, tagInfo.QuestionNo, tagUuid)
			err := db.Create(&record).Error
			if err != nil {
				return model.Record{}, -6
			}
			if !record.Correct {
				return record, 2
			}
		} else {
			return model.Record{}, 3
		}
	}
	return record, 1
}

// 問題タグか?
func judgeSentenceTag(tagId string) bool {
	var tags []model.Question
	db.Where("sentence = ?", tagId).Find(&tags)
	return len(tags) != 0

}

// 回答記録作成
func createRecord(name string, childId int, bookId int, questionNo int, tagUuid string) model.Record {
	var correct bool
	var record model.Record
	
	genreId := GetBookData(bookId)
	correctTag := GetByCorrect(bookId, questionNo)
	correctId := GetTagDataFromTagId(correctTag)
	answerId := GetTagDataFromUuid(tagUuid)
	if correctTag == "" {
		return model.Record{}
	}
	if correctId.Uuid == tagUuid {
		correct = true
	} else {
		correct = false
	}
	record = model.Record{
		Name:       name,
		ChildId:    childId,
		AnswerDay:  time.Now(),
		BookId:     bookId,
		QuestionNo: questionNo,
		GenreId:    genreId[0].GenreId,
		UserAnswer: answerId.TagId,
		Correct:    correct,
	}
	return record
}
