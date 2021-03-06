package controller

import (
	"fmt"
	"github.com/HAL-RO-Developer/caseTeamA/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA/service"
	"github.com/gin-gonic/gin"
)

var Reader = readerimpl{}

type readerimpl struct {
}

// リーダーで読み取った情報の送信
func (r *readerimpl) SendTag(c *gin.Context) {
	var msg string
	req, ok := validation.ReaderValidation(c)
	if !ok {
		response.BadRequest(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	device, find := service.GetDeviceInfoFromDeviceId(req.DeviceId)

	if !find {
		response.BadRequest(gin.H{"error": "デバイスが見つかりませんでした"}, c)
		return
	}
	_, result := service.SendUserAnswer(req.DeviceId, req.Uuid, req.OldUuid)
	if result < 0 {
		if result == -3 {
			service.TalkBocco("タグが登録されてないよ。", device[0].Name)
		}
		response.BadRequest(gin.H{"error": result}, c)
		return
	}

	// boccoAPI実行
	_, find_bocco := service.ExisByBoccoAPI(device[0].Name)
	// 回答記録取得
	userRecord, _ := service.GetByRecordFromChild(device[0].Name, device[0].ChildId)
	// 連続正解数取得
	continueAns := service.GetContinueCorrect(userRecord)
	// メッセージ取得
	message, find := service.GetMessageInfoFromSame(device[0].Name, device[0].ChildId, 3, continueAns)
	if !find {
		message, find = service.GetMessageInfoFromTrue(device[0].Name, device[0].ChildId, result)
	}
	// boccoAPIを実行しないとき
	if !find_bocco {
		if result == 2 || result == 3 {
			response.Json(gin.H{"success": false}, c)
		} else if result == 1 || result == 4 {
			response.Json(gin.H{"success": true}, c)
		}
		return
	}

	if result == 3 {
		message[0].Message = "前回の問題を回答してね"
	}
	
	// ユーザーがメッセージを登録していなかったとき
	if !find {
		childInfo, _ := service.GetByChildInfo(device[0].Name, device[0].ChildId)
		childRecord, _ := service.GetByRecordFromChild(device[0].Name, device[0].ChildId)
		tagInfo := service.GetTagDataFromUuid(req.Uuid)
		continueAns := service.GetContinueCorrect(childRecord)
		defMsg, con, _ := service.GetDefaultMessage(childInfo[0].BirthDay, continueAns, tagInfo.BookId, result)
		
		switch defMsg.MsgCondition {
		case 1:
			msg = fmt.Sprintf(defMsg.Message, childInfo[0].NickName, con.ContinueAns)
		case 2:
			msg = fmt.Sprintf(defMsg.Message, con.ContinueAns)
		default:
			msg = defMsg.Message
		}
	} else {
		msg = message[service.GetRandomNo(len(message))].Message
	}
	if result == 4 {
		response.Json(gin.H{"success": true}, c)
		return
	}
	service.TalkBocco(msg, device[0].Name)
	if result == 2 || result == 3 {
		response.Json(gin.H{"success": false}, c)
		return
	}
	response.Json(gin.H{"success": true}, c)
}
