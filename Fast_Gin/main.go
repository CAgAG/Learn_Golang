package main

import (
	"Gin_go/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// 自定义 中间件
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("cur time", time.Now())
		context.Next() // 放行
		//context.Abort()  // 拦截
	}
}

func main() {
	Server := gin.Default()
	// 连接数据库
	Server_DB, err := NewDB()
	if err != nil {
		panic(err)
	}

	// 连接 Redis
	Server_RDB, err := NewRedis()
	if err != nil {
		panic(err)
	}

	// session
	session_store := cookie.NewStore([]byte("secret"))
	// 使用 redis 来保存 session
	//session_store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	Server.Use(sessions.Sessions("mysession", session_store))
	//Server.Use(myHandler())  // 全局使用中间件

	// GET 请求
	Server.GET("/GET", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	Server.POST("/POST", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})

	// GET 请求 + 参数
	// http://127.0.0.1:8080/User/Info?user_id=cag
	Server.GET("/User/Info", func(context *gin.Context) {
		user_id := context.Query("user_id")
		context.JSON(http.StatusOK, gin.H{
			"user_id": user_id,
		})
	})
	// http://127.0.0.1:8080/User/Info/cag
	Server.GET("/User/Info/:user_id", func(context *gin.Context) {
		user_id := context.Param("user_id")
		context.JSON(http.StatusOK, gin.H{
			"user_id": user_id,
		})
	})

	// POST 请求 + 参数
	// 接收 json 数据
	Server.POST("/json", func(context *gin.Context) {
		data, err := context.GetRawData()
		if err != nil {
			return
		}
		var mapper map[string]interface{}
		_ = json.Unmarshal(data, &mapper)
		context.JSON(http.StatusOK, mapper)
	})

	// POST 请求 + 参数
	// 接收表单
	Server.POST("/user/add", func(context *gin.Context) {
		username, update_flag := context.GetPostForm("username")
		password, update_flag := context.GetPostForm("password")
		if !update_flag {

		}
		new_user := models.User{Username: username, Password: password}
		var find_user models.User
		ret := Server_DB.Model(models.User{Username: username}).First(&find_user)
		if errors.Is(ret.Error, gorm.ErrRecordNotFound) || ret.RowsAffected == 0 { // 两种方式检测 用户是否存在
			Server_DB.Create(&new_user)
		}
		// 更新密码
		if ret.RowsAffected != 0 {
			//find_user.Password = password
			//Server_DB.Save(&find_user)
			Server_DB.Model(&find_user).Updates(models.User{Password: password})
		}

		// session
		session := sessions.Default(context)
		session.Set("username", username)
		session.Save()

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"session":  session.Get("username"),
		})

		// 删除
		//Server_DB.Where("name = ?", "jinzhu").Delete(&models.User{})
		// DELETE from User where id = 10 AND name = "jinzhu";

		// redis
		Server_RDB.Set("cur_user", username)
	})

	// 重定向
	Server.GET("/re", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/")
	})

	// 404页面
	Server.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "找不到页面",
		})
	})

	// 路由组
	op_group := Server.Group("/op")
	{
		// /op/test
		op_group.GET("/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "路由组",
			})
		})
	}

	// ===============================
	Server.LoadHTMLGlob("template/*")
	Server.GET("/", myHandler(), func(context *gin.Context) {
		cur_time, exist := context.Get("cur time")
		if !exist {
			cur_time = "can't get cur_time"
		}
		session := sessions.Default(context)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg":      session.Get("username"),
			"cur_time": cur_time,
		})
	})

	// ===============================
	Port := 8080
	IP := "127.0.0.1"
	Addr := fmt.Sprintf("%s:%d", IP, Port)
	fmt.Println(fmt.Sprintf("%s%s", "http://", Addr))
	Server.Run(Addr)
}
