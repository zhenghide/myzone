package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func SetSession(c *gin.Context, k string, o interface{}) {
	session := sessions.Default(c)
	session.Set(k, o)
	session.Save()
}

func GetSession(c *gin.Context, k string) interface{} {
	session := sessions.Default(c)
	return session.Get(k)
}

func DeleteSession(c *gin.Context, k string)  {
	session := sessions.Default(c)
	session.Delete(k)
	session.Save()
}
