package service

import (
	"github.com/HAL-RO-Developer/caseTeamA/model"
	"math/rand"
	"time"
)

const CONTINUES = 5 // 連続正解数の条件
const PER = 0.8 // 正答率＆解答率の条件

type Condition struct {
	Name        string
	ChildId     int
	Birthday    time.Time
	ContinueAns int
	BookId      int
	Result      int
}

type MessageVariable struct {
	ContinueAns int
	AnsPer  float64
	CorrPer float64
}

// デフォルトメッセージの取得
func GetDefaultMessage(birthday time.Time, continueAns int, bookId int, result int) (model.DefaultMessage, MessageVariable, bool) {
	var condition Condition
	var msgVariable MessageVariable
	var useCondition []int

	condition.Birthday = birthday
	condition.ContinueAns = continueAns
	condition.BookId = bookId
	condition.Result = result

	// 満たしている条件を取得
	useCondition, msgVariable = getCondition(condition)
	if len(useCondition) != 0 {
		messages, find := getMessage(useCondition[getRandomNo(len(useCondition))])
		if !find {
			return model.DefaultMessage{}, MessageVariable{}, false
		}
		return messages[getRandomNo(len(useCondition) - 1)], msgVariable, true
	}

	return model.DefaultMessage{}, MessageVariable{}, false
}

func getMessage(condition int) ([]model.DefaultMessage, bool) {
	var message []model.DefaultMessage
	db.Where("msg_condition = ?", condition).Find(&message)
	return message, len(message) != 0
}

func getRandomNo(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func getCondition(condition Condition) ([]int, MessageVariable) {
	var useCondition []int
	var msgVariable MessageVariable

	if condition.ContinueAns > CONTINUES {
		msgVariable.ContinueAns = condition.ContinueAns
		useCondition = append(useCondition, 1) // 連続正解時(名前+連続正解数)
		useCondition = append(useCondition, 2) // 連続正解時(連続正解数)
	}
	if condition.Result == 1 || condition.Result == 4 {
		useCondition = append(useCondition, 3) // 正解時
	} else {
		useCondition = append(useCondition, 4) // 不正解時
	}
	// 科目達成率
	childRecord, _ := GetByRecordFromGenre(condition.Name, condition.ChildId, condition.BookId)
	numProblems := GetByQuestion(condition.BookId)
	numAns, numCorr := CountAnswerNum(childRecord)
	ansPer := float64(numAns) / float64(numProblems)
	corrPer := float64(numCorr) / float64(numAns)
	if ansPer > PER {
		useCondition = append(useCondition, 5) // ジャンルごとの解答率
		msgVariable.AnsPer = ansPer
	}
	if corrPer > PER {
		useCondition = append(useCondition, 6) // ジャンルごとの正答率
		msgVariable.CorrPer = corrPer
	}
	
	return useCondition, msgVariable
}
