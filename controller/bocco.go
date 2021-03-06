package controller

import (
	"github.com/HAL-RO-Developer/caseTeamA/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA/service"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/config"
)

var Bocco = boccoimpl{}

type boccoimpl struct {
}

// BOCCOAPI設定登録
func (b *boccoimpl) RegistBocco(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.BoccoValidation(c)
	if !ok {
		return
	}

	_, find := service.ExisByBoccoAPI(name)
	if find {
		response.Json(gin.H{"success": "BOCCOAPI設定登録済みです。"}, c)
		return
	}

	_, ok = service.GetBoccoToken(req.Email, config.Env("apikey"), req.Password)
	if !ok {
		response.BadRequest(gin.H{"error": "BOCCOAPIのアクセストークンが取得できませんでした。"}, c)
		return
	}

	err := service.RegistrationBoccoInfo(name, req.Email, req.Password)

	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "メールアドレスとパスワードを登録しました。"}, c)
}

// BOCCOAPI設定取得
func (b *boccoimpl) GetBoccoInfo(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	info, find := service.ExisByBoccoAPI(name)
	if !find {
		response.Json(gin.H{"email": ""}, c)
		return
	}

	response.Json(gin.H{"email": info[0].Email}, c)
}

// BOCCOAPI設定削除
func (b *boccoimpl) DeleteBoccoInfo(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}
	_, find := service.ExisByBoccoAPI(name)
	if !find {
		response.BadRequest(gin.H{"error": "BOCCOAPI設定が見つかりませんでした。"}, c)
		return
	}
	service.DeleteBoccoInfo(name)
	response.Json(gin.H{"success": "メールアドレスとパスワードを削除しました。"}, c)
}
