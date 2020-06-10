package controller4

import (
	"TianCheng/entity"
	"TianCheng/router/middleware"
	"TianCheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Login(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",nil)
}
func Checkup(c *gin.Context)  {
		username:=c.Query("username")
		password:=c.Query("password")
		user,err:=entity.GetUser(username,password)
		if user.Phonestatus!=-1 && user.Emailstatus!=-1{
			c.JSON(http.StatusOK, gin.H{
				"code": 201,
			})
			return
		}
		entity.Deletetoken(user.Phone)
		if err==nil{
			// 生成Token
			tokenString, _ := middleware.GenToken(user.Username,user.Password)
			c.JSON(http.StatusOK,gin.H{
				"code": 200,
				"token":tokenString,
				"data": gin.H{
					"username":user.Username,
				},
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"code": 201,
			})
		}
}
func Index(c *gin.Context){
	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	formatStr := "2006-01-02 15:04:05"
	Date:=fmt.Sprintln(time.Now().Format(formatStr))
	pays:=entity.Getpay(Date)
	money:=0
	for  _,v := range pays{
		money=money+v.Money
	}
	usercount:=entity.Getuser10()
	testcount:=entity.Gettest3()
	d1,d2,d3,d4,d5,d6,d7:=util.Timediffer5(Date)
	da:=strings.Split(Date," ")
	d0:=da[0]+" 00:00:00"
	_,n1:=entity.Countuser2(d0,d1)
	_,n2:=entity.Countuser2(d1,d2)
	_,n3:=entity.Countuser2(d2,d3)
	_,n4:=entity.Countuser2(d3,d4)
	_,n5:=entity.Countuser2(d4,d5)
	_,n6:=entity.Countuser2(d5,d6)
	_,n7:=entity.Countuser2(d6,d7)

	m11:=0
	m22:=0
	m33:=0
	m44:=0
	m55:=0
	m66:=0
	m77:=0
	m1,err1:=entity.Countpay(d0,d1)
	if err1==nil {
		for _,v := range m1{
			m11=m11+v.Money
		}
	}
	m2,err2:=entity.Countpay(d1,d2)
	if err2==nil {
		for _,v := range m2{
			m22=m22+v.Money
		}
	}
	m3,err3:=entity.Countpay(d2,d3)
	if err3==nil {
		for _,v := range m3{
			m33=m33+v.Money
		}
	}
	m4,err4:=entity.Countpay(d3,d4)
	if err4==nil {
		for _,v := range m4{
			m44=m44+v.Money
		}
	}
	m5,err5:=entity.Countpay(d4,d5)
	if err5==nil {
		for _,v := range m5{
			m55=m55+v.Money
		}
	}
	m6,err6:=entity.Countpay(d5,d6)
	if err6==nil {
		for _,v := range m6{
			m66=m66+v.Money
		}
	}
	m7,err7:=entity.Countpay(d6,d7)
	if err7==nil {
		for _,v := range m7{
			m77=m77+v.Money
		}
	}

	fmt.Println(m11,m22,m33,m44,m55,m66,m77)
	d11:=strings.Split(d1," ")
	d22:=strings.Split(d2," ")
	d33:=strings.Split(d3," ")
	d44:=strings.Split(d4," ")
	d55:=strings.Split(d5," ")
	d66:=strings.Split(d6," ")
	d77:=strings.Split(d7," ")
	c.HTML(http.StatusOK,"index.html",gin.H{
		"username":username,
		"token":token,
		"money":money,
		"usercount":usercount,
		"testcount":testcount,
		"n1":n1,
		"n2":n2,
		"n3":n3,
		"n4":n4,
		"n5":n5,
		"n6":n6,
		"n7":n7,
		"m1":m11,
		"m2":m22,
		"m3":m33,
		"m4":m44,
		"m5":m55,
		"m6":m66,
		"m7":m77,
		"d1":d11[0],
		"d2":d22[0],
		"d3":d33[0],
		"d4":d44[0],
		"d5":d55[0],
		"d6":d66[0],
		"d7":d77[0],
	})
}
func Yhgl(c *gin.Context){
	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	users:=entity.Getuser11()
	c.HTML(http.StatusOK,"yonghuguanli.html",gin.H{
		"token":token,
		"username":username,
		"users":users,
	})
}
func Findusers(c *gin.Context){
	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	name:=c.Query("name")
	ysy:=c.Query("ysy")
	users:=entity.Getuser12(name,ysy)
	c.HTML(http.StatusOK,"yonghuguanli.html",gin.H{
		"token":token,
		"username":username,
		"users":users,
	})
}
type QU struct {
	Qiye entity.Qiye
	User entity.User
}
func Qygl(c *gin.Context)  {
	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	qyname:=c.Query("qyname")
	if qyname != "" {
		qiyes:=entity.Getqiye3(qyname)
		qiyeids:=make([]int32,len(qiyes))
		for i:=0;i<len(qiyes);i++{
			qiyeids[i]=qiyes[i].Qiyeid
		}
		users:=entity.Getuser14(qiyeids)
		qu:=make([]QU, len(qiyes))
		for i:=0;i<len(qiyes);i++{
			qu[i].Qiye=qiyes[i]
			qu[i].User=users[i]
		}
		c.HTML(http.StatusOK,"qiyeguanli.html",gin.H{
			"username":username,
			"token":token,
			"qu":qu,
		})
		return
	}
	qiyes:=entity.Getqiye2()
	users:=entity.Getuser13()
	qu:=make([]QU, len(qiyes))
	for i:=0;i<len(qiyes);i++{
		qu[i].Qiye=qiyes[i]
		qu[i].User=users[i]
	}
	c.HTML(http.StatusOK,"qiyeguanli.html",gin.H{
		"username":username,
		"token":token,
		"qu":qu,
	})
}
func Xzqycheck(c *gin.Context){
	qymc:=c.Query("qymc")
	glyname:=c.Query("glyname")
	glyphone:=c.Query("glyphone")
	glyemail:=c.Query("glyemail")
	glynumber:=c.Query("glynumber")
	qiye:=entity.Getqiye4(qymc)
	user:=entity.Getuser15(glyname,glyphone,glyemail,glynumber)
	if qiye.Id==0 && user.ID==0{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
		return
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
	}
}
func Xzqy(c *gin.Context){
	qymc:=c.Query("qymc")
	glyname:=c.Query("glyname")
	glyphone:=c.Query("glyphone")
	glyemail:=c.Query("glyemail")
	glymm:=c.Query("glymm")
	glynumber:=c.Query("glynumber")
	qiyes:=entity.Getqiye2()
	qy:=new(entity.Qiye)
	qy.Qiyename=qymc
	id1:=qiyes[len(qiyes)-1].Qiyeid+1
	qy.Qiyeid=id1
	entity.Createqiye(qy)
	us:=new(entity.User)
	us.Username=glyname
	us.Name=glyname
	us.Password=glymm
	us.Email=glyemail
	us.Phone=glyphone
	us.Number=glynumber
	us.Qiyeid=id1
	us.Qiyepower=1
	us.Vip=0
	us.Fabunumber=5
	us.Sex="男"
	us.Emailstatus=0
	us.Phonestatus=0
	formatStr := "2006-01-02 15:04:05"
	us.Ruzhidate=fmt.Sprintln(time.Now().Format(formatStr))
	entity.Createuser2(us)

	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	qiyess:=entity.Getqiye2()
	users:=entity.Getuser13()
	qu:=make([]QU, len(qiyess))
	for i:=0;i<len(qiyess);i++{
		qu[i].Qiye=qiyess[i]
		qu[i].User=users[i]
	}
	c.HTML(http.StatusOK,"qiyeguanli.html",gin.H{
		"username":username,
		"token":token,
		"qu":qu,
	})
}
func Tzgl(c *gin.Context){
	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	title:=c.Query("title")
	if title!="" {
		tz:=entity.Gettz3(title)
		c.HTML(http.StatusOK,"tongzhigonggao.html",gin.H{
			"username":username,
			"token":token,
			"tz":tz,
		})
		return
	}
	tz:=entity.Gettz()
	c.HTML(http.StatusOK,"tongzhigonggao.html",gin.H{
		"username":username,
		"token":token,
		"tz":tz,
	})
}
func Tznr(c *gin.Context){
	id:=c.Query("id")
	tz:=entity.Gettz2(id)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"tz":tz,
	})
}
func Bjtz(c *gin.Context)  {
	id:=c.Query("tzid")
	title:=c.Query("title")
	text:=c.Query("text")
	entity.Updatatz(id,title,text)

	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	tz:=entity.Gettz()
	c.HTML(http.StatusOK,"tongzhigonggao.html",gin.H{
		"username":username,
		"token":token,
		"tz":tz,
	})
}
func Xztz(c *gin.Context){
	title:=c.Query("title")
	text:=c.Query("text")
	tz:=new(entity.Gonggao)
	tz.Title=title
	tz.Text=text
	entity.Createtz(tz)

	username:=c.MustGet("username").(string)
	token:=c.MustGet("token").(string)
	tz1:=entity.Gettz()
	c.HTML(http.StatusOK,"tongzhigonggao.html",gin.H{
		"username":username,
		"token":token,
		"tz":tz1,
	})
}
func Dengchu(c *gin.Context){
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	tok:=new(entity.Tokeninvalid)
	tok.Token=token
	tok.Userphone=user.Phone
	entity.Createtoken(tok)
	c.HTML(http.StatusOK,"login.html", nil)
}
func Lh(c *gin.Context)  {
	id:=c.Query("id")
	entity.Deleteuser2(id)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}
func Jc(c *gin.Context)  {
	id:=c.Query("id")
	entity.Jiechu(id)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}
func Sctz(c *gin.Context)  {
	id:=c.Query("id")
	entity.Deletetz(id)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}