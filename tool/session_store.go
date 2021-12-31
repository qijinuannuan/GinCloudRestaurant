package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

func InitSession(engine *gin.Engine) {
	config := GetConfig().Redis
	store, err := redis.NewStore(10, "tcp", config.Addr + ":" + config.Port, config.Password, []byte("secret"))
	if err != nil {
		log.Fatal(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

func SetSession(ctx *gin.Context, key, value interface{}) error {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

func GetSession(ctx *gin.Context, key interface{}) interface{} {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}
	return session.Get(key)
}