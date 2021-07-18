package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		c.JSON(200, data)
	})

	r.GET("/ping1", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(200, data)
	})

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 36000, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
		c.JSON(200, gin.H{"a": "b"})
	})

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// in this case proper binding will be automatically selected
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	r.GET("/someJson", func(c *gin.Context) {
		names := []string{"lenan", "sfls", "ddd"}
		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(200, names)
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		c.String(200, "Hello %s %s", firstname, lastname)
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(200, gin.H{"a": "b"})
	})
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	p()
	r.Run(":9999")
}

func p() {
	s := "200"
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Printf("type : %T value : %#v \n", num, num)
	}
	s1 := 10
	str := strconv.Itoa(s1)

	fmt.Printf("type: %T, value: %#v \n", str, str)

	num1, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("can not convert int")
	} else {
		fmt.Printf("type : %T value : %#v \n", num1, num1)
	}
	num2, err := strconv.ParseInt("10", 10, 32)
	if err != nil {
		fmt.Println("can not convert int")
	} else {
		fmt.Printf("type : %T value : %#v \n", num2, num2)
	}

	str2 := strconv.FormatInt(5, 2)
	fmt.Printf("type : %T value : %#v \n", str2, str2)
	str3 := strconv.FormatFloat(5.334344, 'f', 2, 32)
	fmt.Printf("type : %T value : %#v \n", str3, str3)

}

func fprintln() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件具柄写入数据
	fmt.Fprintf(fileObj, "向文件总写入信息: %s", name)
	fileObj.Close()
}

func sprint() {
	s1 := fmt.Sprint("沙河")
	name := "王子"
	age := 12
	s2 := fmt.Sprintf("name:%s, age: %d", name, age)
	s3 := fmt.Sprintln("hello")
	fmt.Println(s1, s2, s3)
}

func errorf() {
	err := fmt.Errorf("这是个错误")
	fmt.Print(err)
}
func currentTime() {
	now := time.Now()
	fmt.Printf("%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("timestamp1: %v, timestamp2: %v", timestamp1, timestamp2)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func timeAdd() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}
func formatDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
func logP() {
	log.Println("log...")
	v := "普通"
	log.Printf("这是一条%s的日志", v)
	//log.Panicln("这是一条触发panic的日志")
	//log.Fatalln("这是一条触发fatal的日志")

}
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
