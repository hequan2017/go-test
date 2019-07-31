package internal

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func jsonError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": msg})
}
func jsonAuthError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"ok": false, "msg": msg})
}

func jsonData(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": true, "data": data})
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		//logrus.WithError(err).Error("gin context http handler error")
		jsonError(c, err.Error())
		return true
	}
	return false
}
func handlerAuthMiddlewareError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"msg": err.Error()})
		return true
	}
	return false
}

func wshandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		logrus.WithError(err).Error("handler ws ERROR:")
		dt := time.Now().Add(time.Second)
		if err := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); err != nil {
			logrus.WithError(err).Error("websocket writes control message failed:")
		}
		return true
	}
	return false
}
