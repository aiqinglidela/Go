package router

import (
	"TianCheng/controller"
	"TianCheng/controller4"
	"TianCheng/core"
	"TianCheng/entity"
	"TianCheng/router/middleware"
	"TianCheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)
func Router01() http.Handler {
	e := gin.Default()
	e.Static("/static1", "static1")
	e.LoadHTMLGlob("templates1/*")
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", nil)
	})
	e.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", nil)
	})
	e.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK,"about.html", nil)
	})
	e.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK,"blog.html", nil)
	})
	e.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK,"contact.html", nil)
	})
	e.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK,"service.html", nil)
	})
	e.GET("/dengchu",middleware.JWTAuthMiddleware(),func(c *gin.Context) {
		username := c.MustGet("username").(string)
		password := c.MustGet("password").(string)
		token := c.MustGet("token").(string)
		user, _ := entity.GetUser(username,password)
		tok:=new(entity.Tokeninvalid)
		tok.Token=token
		tok.Userphone=user.Phone
		entity.Createtoken(tok)
		c.HTML(http.StatusOK,"index.html", nil)
	})
	return e
}

func Router02() http.Handler {
	e := gin.Default()
	e.Static("/static2", "static2")
	e.LoadHTMLGlob("templates2/*")
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.GET("/sign1", controller.Sign1)
	e.GET("/sign2", controller.Sign2)
	e.POST("/zhuce", controller.CreateUser)
	e.GET("/login1", controller.Login1)
	e.GET("/login2", controller.Login2)
	e.GET("/login3", controller.Login3)
	e.GET("/wjmm", controller.Wjmm)
	e.GET("/czmm",controller.Czmm)
	e.GET("/czm",controller.Czm)
	e.PUT("/cm",controller.Cm)
	e.GET("/checkemail", controller.Checkemail)
	e.GET("/fsyzm", controller.Fsyzm)
	e.POST("/denglu", controller.GetUser)
	//e.POST("/auth", controller.AuthHandler)
	//e.GET("/home", middleware.JWTAuthMiddleware(), homeHandler)
	//e.GET("/index2",middleware.JWTAuthMiddleware(), controller.HomeHandler)
	e.GET("/ksgl", middleware.JWTAuthMiddleware(),controller.Ksgl)
	e.GET("/glks", middleware.JWTAuthMiddleware(),controller.Glks)
	e.DELETE("/scks",middleware.JWTAuthMiddleware(),controller.Scks)
	e.GET("/glks2", middleware.JWTAuthMiddleware(),controller.Glks2)
	e.POST("/tjks", middleware.JWTAuthMiddleware(),controller.Tjks)
	e.GET("/pysj", middleware.JWTAuthMiddleware(),controller.Pysj)
	e.GET("/pysj2", middleware.JWTAuthMiddleware(),controller.Pysj2)
	e.PUT("/jiaf",middleware.JWTAuthMiddleware(),controller.Jiaf)
	e.DELETE("/kssc",middleware.JWTAuthMiddleware(),controller.Kssc)
	e.POST("/ksxq", middleware.JWTAuthMiddleware(),controller.Ksxq)
	e.GET("/index", middleware.JWTAuthMiddleware(), controller.Home)
	e.POST("/index", middleware.JWTAuthMiddleware(), controller.Home)
	e.GET("/fabuks", middleware.JWTAuthMiddleware(), controller.Fabuks)
	e.GET("/fabuks2", middleware.JWTAuthMiddleware(), controller.Fabuks2)
	e.POST("/zj", middleware.JWTAuthMiddleware(), controller.Zj)
	e.GET("/ks", middleware.JWTAuthMiddleware(), controller.Ks)
	e.GET("/ksks", middleware.JWTAuthMiddleware(), controller.Ksks)
	e.POST("/jiaojuan", middleware.JWTAuthMiddleware(), controller.Jiaojuan)
	e.GET("/kcsz",middleware.JWTAuthMiddleware(), controller.Kcsz)
	e.GET("/kcxq",middleware.JWTAuthMiddleware(), controller.Kcxq)
	e.GET("/xjkc",middleware.JWTAuthMiddleware(), controller.Xjkc)
	e.PUT("tjkc",middleware.JWTAuthMiddleware(), controller.Tjkc)
	e.POST("/sckc",middleware.JWTAuthMiddleware(), controller.Sckc)
	e.GET("/glsj",middleware.JWTAuthMiddleware(), controller.Glsj)
	e.GET("/sjxq",middleware.JWTAuthMiddleware(), controller.Sjxq)
	e.POST("/scsj",middleware.JWTAuthMiddleware(), controller.Scsj)
	e.GET("/drsj",middleware.JWTAuthMiddleware(), controller.Drsj)
	e.GET("/ksbb",middleware.JWTAuthMiddleware(), controller.Ksbb)
	e.GET("/ckbb",middleware.JWTAuthMiddleware(), controller.Ckbb)
	e.GET("/qyyhgl",middleware.JWTAuthMiddleware(), controller.Qyyhgl)
	e.GET("/grzx",middleware.JWTAuthMiddleware(), controller.Grzx)
	e.GET("/gg",middleware.JWTAuthMiddleware(), controller.Gg)
	e.GET("/qysz",middleware.JWTAuthMiddleware(), controller.Qysz)
	e.POST("/xzbm",middleware.JWTAuthMiddleware(), controller.Xzbm)
	e.DELETE("/scbm",middleware.JWTAuthMiddleware(), controller.Scbm)
	e.GET("/cxyg",middleware.JWTAuthMiddleware(), controller.Cxyg)
	e.GET("/ygxx",middleware.JWTAuthMiddleware(), controller.Ygxx)
	e.PUT("/xxxg",middleware.JWTAuthMiddleware(), controller.Xxxg)
	e.GET("checkp",middleware.JWTAuthMiddleware(),controller.Checkp)
	e.PUT("/xgmm",middleware.JWTAuthMiddleware(),controller.Xgmm)
	e.GET("/zjcs",middleware.JWTAuthMiddleware(),controller.Zjcs)
	e.GET("/vip",middleware.JWTAuthMiddleware(),controller.Vip)
	e.GET("/return",middleware.JWTAuthMiddleware(),controller.Return)
	e.GET("/return2",middleware.JWTAuthMiddleware(),controller.Return2)
	e.POST("/alipaynotify",middleware.JWTAuthMiddleware(),controller.AliPayNotify)
	e.GET("scyg",middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		id:=c.Query("ygid")
		entity.Deleteuser(id)
		ss:=fmt.Sprint("/index")
		c.Request.URL.Path = ss
		e.HandleContext(c)
	})
	e.POST("/drks", middleware.JWTAuthMiddleware(), func (c *gin.Context){
		//单个文件
		file, err := c.FormFile("kaosheng")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		dst := fmt.Sprintf("E:/GoPath/src/TianCheng/proto/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		/*解析文档*/
		phones:=util.Drks(dst)
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"phones":phones,
		})

	})
	e.POST("/tjsj", middleware.JWTAuthMiddleware(), func (c *gin.Context){
		username:=c.MustGet("username").(string)
		password:=c.MustGet("password").(string)
		user,_:=entity.GetUser(username,password)
		examname:=c.PostForm("examname")
		//单个文件
			file, err := c.FormFile("sj")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			dst := fmt.Sprintf("E:/GoPath/src/TianCheng/proto/%s", file.Filename)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
			/*解析文档*/
			n1,n2,n3,n4,n5,n6,n7,n8,n9:=util.Exam(dst)
			var exam entity.Exam
			exam.Examname=examname
			exam.Danxuan=int32(n1)
			exam.Duoxuan=int32(n3)
			exam.Panduan=int32(n5)
			exam.Jianda=int32(n7)
			exam.Manfen=int32(n9)
			n10:=float32(n9)*0.6
			exam.Pass=int32(n10)
			exam.Fabuid=user.ID
			err=entity.Createexam(&exam)
			ex:=entity.Getexam3(examname)
			times:=util.Exam2(dst)
			j:=0
			if n1!=0 {
				for i:=0;i<n1;i++{
					times[j].Grade=int32(n2)
					times[j].Leixing="danxuan"
					times[j].Shijuanid=ex.Id
					j++
				}
			}
			if n3!=0{
				for i:=0;i<n3;i++{
					times[j].Grade=int32(n4)
					times[j].Leixing="duoxuan"
					times[j].Shijuanid=ex.Id
					j++
				}
			}
			if n5!=0 {
				for i:=0;i<n5;i++{
					times[j].Grade=int32(n6)
					times[j].Leixing="panduan"
					times[j].Shijuanid=ex.Id
					j++
				}
			}
			if n7!=0 {
				for i:=0;i<n7;i++{
					times[j].Grade=int32(n8)
					times[j].Leixing="jianda"
					times[j].Shijuanid=ex.Id
					j++
				}
			}
			for i:=0;i<len(times);i++{
				entity.Createtime(&times[i])
			}
		ss:=fmt.Sprint("/index")
		c.Request.URL.Path = ss
		e.HandleContext(c)

	})
	e.POST("/fb", middleware.JWTAuthMiddleware(), func (c *gin.Context){
		username:=c.MustGet("username").(string)
		password:=c.MustGet("password").(string)
		//token:=c.MustGet("token")
		//s:=fmt.Sprint(token)
		name:=c.PostForm("kaoshiname")
		sjname:=c.PostForm("sj")
		kcname:=c.PostForm("kc")
		starttime:=c.PostForm("kssj")
		endtime:=c.PostForm("jssj")
		lp:=c.PostForm("lp")
		yz:=c.PostForm("sf")
		strings.Replace(starttime,"T"," ",1)
		strings.Replace(endtime,"T"," ",1)
		fmt.Println("name=",name)
		fmt.Println("sjname=",sjname)
		fmt.Println("starttime=",starttime)
		fmt.Println("endtime=",endtime)
		fmt.Println("kcname=",kcname)
		fmt.Println("lp=",lp)
		fmt.Println("yz=",yz)
		var phones string
		user,_:=entity.GetUser(username,password)
		if yz=="1"{
			/*取出企业所有员工*/
			users,_:=entity.Getuser5(user.Qiyeid)
			if len(users)!=0{
				for _,v:=range users {
					phones=phones+","+v.Phone
				}
			}
		}else {
			phones=c.PostForm("phones")
		}
		var test entity.Test
		test.Kaoshiname=name
		id1,_:=strconv.Atoi(kcname)
		test.Kaochangid=int32(id1)
		id2,_:=strconv.Atoi(sjname)
		test.Shijuanid=int32(id2)
		test.Starttime=starttime
		test.Endtime=endtime
		test.Token=lp
		test.Phones=phones
		test.Fabuid=user.ID
		err:=entity.Creattest(test)
		if err!=nil {
			fmt.Println(err)
		}
		ss:=fmt.Sprint("/index")
		c.Request.URL.Path = ss
		e.HandleContext(c)
	})
	e.POST("/fb2", middleware.JWTAuthMiddleware(), func (c *gin.Context){
		username:=c.MustGet("username").(string)
		password:=c.MustGet("password").(string)
		//token:=c.MustGet("token")
		//s:=fmt.Sprint(token)
		name:=c.PostForm("kaoshiname")
		sjname:=c.PostForm("sj")
		kcname:=c.PostForm("kc")
		starttime:=c.PostForm("kssj")
		endtime:=c.PostForm("jssj")
		lp:=c.PostForm("lp")
		starttime=strings.Replace(starttime,"T"," ",1)
		starttime=starttime+":00"
		endtime=strings.Replace(endtime,"T"," ",1)
		endtime=endtime+":00"
		fmt.Println("name=",name)
		fmt.Println("sjname=",sjname)
		fmt.Println("starttime=",starttime)
		fmt.Println("endtime=",endtime)
		fmt.Println("kcname=",kcname)
		fmt.Println("lp=",lp)
		var phones string
		user,_:=entity.GetUser(username,password)
		phones=c.PostForm("phones")
		var test entity.Test
		test.Kaoshiname=name
		id1,_:=strconv.Atoi(kcname)
		test.Kaochangid=int32(id1)
		id2,_:=strconv.Atoi(sjname)
		test.Shijuanid=int32(id2)
		test.Starttime=starttime
		test.Endtime=endtime
		test.Token=lp
		test.Phones=phones
		test.Fabuid=user.ID
		err:=entity.Creattest(test)
		if err!=nil {
			fmt.Println(err)
		}
		ss:=fmt.Sprint("/index")
		c.Request.URL.Path = ss
		e.HandleContext(c)
	})
	return e
}
var Room = core.NewRoom()
func Router03() http.Handler {
	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.Static("/static3", "static3")
	e.StaticFile("/", "templates3/index.html")
	e.StaticFile("/refresh", "./templates3/refresh.html")
	e.StaticFile("/polling", "./templates3/polling.html")
	e.StaticFile("/ws", "./templates3/ws.html")
	//{
	//	// refresh
	//	e.GET("/refresh/archive", Refresh.Archive())
	//	e.POST("/refresh/msg", Refresh.Msg())
	//	e.GET("/refresh/leave", Refresh.Leave())
	//}
	//
	//{
	//	// polling
	//	e.GET("/polling/archive", LongPolling.Archive())
	//	e.POST("/polling/msg", LongPolling.Msg())
	//	e.GET("/polling/leave", LongPolling.Leave())
	//
	//}

	{
		// websocket
		e.GET("/ws/socket", Websocket.Handle())
	}
	return e
}
func Router04() http.Handler {
	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.Static("/static4", "static4")
	e.LoadHTMLGlob("templates4/*")
	e.GET("/", controller4.Login)
	e.GET("/login",controller4.Checkup)
	e.GET("/index",middleware.JWTAuthMiddleware(),controller4.Index)
	e.GET("/yhgl",middleware.JWTAuthMiddleware(),controller4.Yhgl)
	e.GET("/findusers",middleware.JWTAuthMiddleware(),controller4.Findusers)
	e.GET("/qygl",middleware.JWTAuthMiddleware(),controller4.Qygl)
	e.GET("/xzqycheck",middleware.JWTAuthMiddleware(),controller4.Xzqycheck)
	e.GET("/xzqy",middleware.JWTAuthMiddleware(),controller4.Xzqy)
	e.GET("/tzgl",middleware.JWTAuthMiddleware(),controller4.Tzgl)
	e.GET("/tznr",middleware.JWTAuthMiddleware(),controller4.Tznr)
	e.GET("/bjtz",middleware.JWTAuthMiddleware(),controller4.Bjtz)
	e.GET("/xztz",middleware.JWTAuthMiddleware(),controller4.Xztz)
	e.GET("/dengchu",middleware.JWTAuthMiddleware(),controller4.Dengchu)
	e.GET("/lh",middleware.JWTAuthMiddleware(),controller4.Lh)
	e.GET("/jc",middleware.JWTAuthMiddleware(),controller4.Jc)
	e.GET("/sctz",middleware.JWTAuthMiddleware(),controller4.Sctz)
	return e
}
