package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

//var timeNow = time.Now
var timeNow = func() time.Time {
	return time.Date(2020, 11, 12, 0, 0, 0, 0, time.UTC)
}

func GenerateTimestamp() int64 {
	now := timeNow()  // current local time
	sec := now.Unix() // number of seconds since January 1, 1970 UTC

	return sec // int 64
}

func main() {
	//var Age int
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(LoggerByTime()) //中间件：日志
	router.Use(gin.Recovery())
	router.Use(Cors()) //中间件:跨域

	api := router.Group("/pc").Use(AuthRequiredPc) //pc开头的路由中间件:验证
	{
		api.POST("/add", func(c *gin.Context) { //参数方式：body
			var (
				user User
				rtn  Rtn
			)
			err := c.ShouldBindJSON(&user)
			if err != nil {
				// 获取validator.ValidationErrors类型的errors
				if errs, ok := err.(validator.ValidationErrors); ok {
					// validator.ValidationErrors类型错误直接返回
					c.JSON(http.StatusOK, gin.H{
						"msg": errs.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
				return
				//fmt.Println("Add：" + err.Error())
				//rtn.R = 0
				//rtn.Err = err.Error()
				//c.JSON(http.StatusInternalServerError, rtn)
				//panic(err)
				//return
			}

			fmt.Println(GetLogFileName(), "添加了一个用户，名="+user.Name+",年龄="+strconv.Itoa(user.Age))
			rtn.R = 1
			c.JSON(http.StatusOK, rtn)
		})
		api.PUT("/edit", func(c *gin.Context) { //参数方式：body
			var (
				user User
				rtn  Rtn
			)
			err := c.ShouldBindJSON(&user)
			if err != nil {
				fmt.Println(GetLogFileName(), "Edit："+err.Error())
				rtn.R = 0
				rtn.Err = err.Error()
				c.JSON(http.StatusInternalServerError, rtn)
				panic(err)
				return
			}
			fmt.Println(GetLogFileName(), "编辑了用户，名="+user.Name+",年龄="+strconv.Itoa(user.Age))
			rtn.R = 1
			c.JSON(http.StatusOK, rtn)
		})
		api.DELETE("/del", func(c *gin.Context) { //参数方式：post form
			var rtn Rtn
			name := c.PostForm("name")
			fmt.Println(GetLogFileName(), "删除用户名="+name)
			rtn.R = 1
			c.JSON(http.StatusOK, rtn)
		})
	}
	apiOut := router.Group("/mobile").Use() //mobile开头的路由中间件:验证
	{
		apiOut.GET("/detail", func(c *gin.Context) { //参数方式：query
			var rtn Rtn
			name := c.Query("name")
			fmt.Println(GetLogFileName(), "详情："+"该用户名="+name)
			rtn.R = 1
			rtn.Data = name + strconv.FormatInt(GenerateTimestamp(), 10)
			c.JSON(http.StatusOK, rtn)
		})
	}

	_ = router.Run(":8888")
}

// GetLogFileName 中间件：日志。参考gin自带的日志中间件实现日志按照天存储，详见gin/logger.go部分
func GetLogFileName() *os.File {
	filename := fmt.Sprintf("./%s", time.Now().Format("20060102"))
	fileObj, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	return fileObj
}

type logger struct {
	Time       string // Time shows the time after the server returns a response.
	StatusCode int    // StatusCode is HTTP response code.
	Latency    string // Latency is how much time the server cost to process a certain request.
	ClientIP   string // ClientIP equals Context's ClientIP method.
	Method     string // Method is the HTTP method given to the request.
	Path       string // Path is a path the client requests.
}

func LoggerByTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		param := logger{}
		// Stop timer
		param.Time = time.Now().Format("2006-01-02 15:04:05")
		param.Latency = fmt.Sprintf("%dms", time.Now().Sub(start).Nanoseconds()/1e6)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path
		paramStr, _ := json.Marshal(param)
		fmt.Println(GetLogFileName(), string(paramStr))
	}
}

// Cors 中间件：跨域。可根据自己需求进行调整
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Origin,Authorization,Content-Length,Content-Type,Date,DataField")
			c.Header("Access-Control-Max-Age", "3600")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

// Rtn 中间件：拦截校验
//拦截校验
type Rtn struct {
	R    int         `json:"r"`
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

func AuthRequiredPc(c *gin.Context) {
	var rtn Rtn
	var err error
	authorization := c.Request.Header.Get("Authorization")
	if authorization != "pc" {
		rtn.Data = 0
		err = errors.New("authorization error")
		fmt.Println(GetLogFileName(), "AuthRequired："+err.Error())
		rtn.Err = err.Error()
		c.JSON(http.StatusUnauthorized, rtn)
		panic(err)
		return
	}
}

// AuthRequiredMobile
// Test:curl -v -H 'Authorization:mobile' http://localhost:8888/mobile/detail
func AuthRequiredMobile(c *gin.Context) {
	var rtn Rtn
	var err error
	authorization := c.Request.Header.Get("Authorization")
	if authorization != "mobile" {
		rtn.Data = 0
		err = errors.New("authorization error")
		fmt.Println(GetLogFileName(), "AuthRequired："+err.Error())
		rtn.Err = err.Error()
		c.JSON(http.StatusUnauthorized, rtn)
		panic(err)
		return
	}
}
