package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"net/http"
	"strconv"
)

// 处理增加答疑记录请求
func HandleCreateChat(c *gin.Context) {
	QuserId := c.PostForm("from")
	AuserId := c.PostForm("to")
	token, err := strconv.Atoi(c.PostForm("token"))

	if err != nil {
		fmt.Println("wrong token number")
	}

	chat := model.CreateChat(model.Db, QuserId, AuserId, token)
	c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
}

// 处理一条答疑已完成的请求
func HandleEndChat(c *gin.Context) {
	hash := c.Query("hash")
	chat := model.EndChat(model.Db, hash)
	if chat.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}
