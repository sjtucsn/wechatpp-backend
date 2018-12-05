package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"net/http"
	"strconv"
)

// 处理增加答疑记录请求
func HandleCreateChat(c *gin.Context) {
	QuserId := c.PostForm("from")
	AuserId := c.PostForm("to")
	if !utils.VerifyParams(c, map[string]string{"from": QuserId, "to": AuserId}) {
		return
	}

	chat := model.CreateChat(model.Db, QuserId, AuserId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
}

// 处理一条答疑已完成的请求
func HandleEndChat(c *gin.Context) {
	hash := c.PostForm("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	chat := model.EndChat(model.Db, hash)
	if chat.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

// 处理给答疑打分的请求
func HandleScoreChat(c *gin.Context) {
	hash := c.PostForm("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	grade, err := strconv.Atoi(c.PostForm("grade"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "wrong grade number"})
		return
	}

	chat := model.ScoreChat(model.Db, hash, grade)
	if chat.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "msg": "no such chat"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
	}
}
