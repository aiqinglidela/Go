package controller

import (
	"TianCheng/entity"
	"TianCheng/router/middleware"
	"TianCheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func Sign1(c *gin.Context) {
	data := map[string]interface{}{
		"code":"200",
	}
	c.JSONP(http.StatusOK, data)
}
func Sign2(c *gin.Context) {
	c.HTML(http.StatusOK,"sign.html", nil)
}

func Login1(c *gin.Context) {
	data := map[string]interface{}{
		"code":"200",
	}
	c.JSONP(http.StatusOK, data)
}
func Login2(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html", nil)
}
func Login3(c *gin.Context) {
	c.HTML(http.StatusOK,"login2.html", nil)
}
func Wjmm(c *gin.Context){
	c.HTML(http.StatusOK,"wangjimima.html", nil)
}
func Fsyzm(c *gin.Context){
	email:=c.Query("email")
	em2:=entity.Getmail(email)
	st:=util.Timediffer3(em2.Date)
	if st==1 {
		entity.Deletemail(email)
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":202,
		})
		return
	}
	yzm:=rand.Intn(1000000)+100000
	yzm2:=strconv.Itoa(yzm)
	em:=new(entity.Mail)
	em.Title="天秤平台验证码"
	em.Aceptemail=email
	em.Message=yzm2
	data:=time.Now().String()
	em.Date=data[:19]
	err:=entity.Createmail(em)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
	}else {
		util.SendMailByGomail(email,yzm2)
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
	}
}
func Checkemail(c *gin.Context)  {
	email:=c.Query("email")
	err :=entity.Checkemail(email)
	if err ==nil {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
	}
}
func Checkuser(username,Password,Password2,Email,Phone string)(data map[string]interface{}){
	data = map[string]interface{}{}
	data["username"]="ok"
	data["email"]="ok"
	data["phone"]="ok"
	data["password"]="ok"
	err1 :=entity.Checkusername(username)
	if err1==nil {
		data["username"]="用户名已存在"
	}
	err2 :=entity.Checkemail(Email)
	if err2==nil {
		data["email"]="邮箱已存在"
	}
	err3 :=entity.Checkphone(Phone)
	if err3==nil {
		data["phone"]="手机号已存在"
	}
	if Password!=Password2{
		data["password"]="密码不一致"
	}
	return data
}
func CreateUser(c *gin.Context)  {
	var user entity.User
	username := c.PostForm("username")
	Password := c.PostForm("password")
	Password2 := c.PostForm("password2")
	Email := c.PostForm("email")
	Phone := c.PostForm("phone")
	data:=Checkuser(username,Password,Password2,Email,Phone)
	user.Username=username
	user.Password=Password
	user.Email=Email
	user.Phone=Phone
	user.Qiyeid=-1
	user.Qiyepower=-1
	user.Sex="男"
	user.Fabunumber=5
	user.Ruzhidate=time.Now().Format("2006-01-02 15:04:05")
	user.Phonestatus=0
	user.Emailstatus=0
	// 2. 存入数据库
	if data["username"]!="ok" || data["email"]!="ok" || data["phone"]!="ok" || data["password"]!="ok"{
		data["code"]="203"
		c.JSON(http.StatusOK,data)
	}else {
		err:=entity.CreateUser(user)
		if err != nil{
			data["code"]="201"
			c.JSON(http.StatusOK,data)
		}else{
			data["code"]="200"
			c.JSON(http.StatusOK, data)
		}
	}
}
func GetUser(c *gin.Context) {
	// 查询这个表里的所有数据
	username := c.PostForm("Username")
	password := c.PostForm("Password")
	user, err := entity.GetUser(username,password)
	entity.Deletetoken(user.Phone)
	if err==nil {
		// 生成Token
		tokenString, _ := middleware.GenToken(user.Username,user.Password)
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"token": tokenString,
				"username":username,
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 202,
		"msg":  "鉴权失败",
	})
}
func Home(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	//tests :=new([]entity.Test)
	//if user.Emailstatus!=0 || user.Phonestatus!=0 {
	//	c.HTML(http.StatusOK,"404.html", gin.H{
	//		"token":token,
	//		"data":"账号有误",
	//		"data1": gin.H{
	//			"username": username,
	//			"password": password,
	//			"Totle1":0,
	//			"Totle2":0,
	//			"Totle3":0,
	//			"Totle4":0,
	//			"Tests":tests,
	//		},
	//	})
	//	return
	//}
	//考试总数
	_,Total1,_:=entity.Countuser(user.Phone)
	//通过的考试
	Total2:=entity.Pass(user.Phone)
	//未通过的考试
	Total3:=entity.Fail(user.Phone)
	//未开始的考试
	test,Total4,_:=entity.Nostart(user.Phone)
	c.HTML(http.StatusOK,"index.html", gin.H{
		"code": 2000,
		"token":token,
		"msg":  "success",
		"data": gin.H{
			"username": username,
			"password": password,
			"Totle1":Total1,
			"Totle2":Total2,
			"Totle3":Total3,
			"Totle4":Total4,
			"Tests":test,
		},
	})
}
func Ksxq(c *gin.Context){
	id := c.Query("id")
	test:=entity.Gettest(id)
	exam:=entity.Getexam(test.Shijuanid)
	kcid2 := *(*int)(unsafe.Pointer(&test.Kaochangid))
	kcid:=strconv.Itoa(kcid2)
	kaochang:=entity.Getkaochang(kcid)
	timu:=entity.Gettimu(exam.Id)
	fbid2 := *(*int)(unsafe.Pointer(&test.Fabuid))
	fbid:=strconv.Itoa(fbid2)
	user,_:=entity.Getuser3(fbid)
	var sum int
	sum=0
	if exam.Danxuan>=1{
		sum++
	}
	if exam.Duoxuan>=1 {
		sum++
	}
	if exam.Jianda>=1 {
		sum++
	}
	if exam.Panduan>=1{
		sum++
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"kcname":test.Kaoshiname,
		"stime":test.Starttime,
		"etime":test.Endtime,
		"kdnumber":kaochang.Number,
		"dati":sum,
		"xiaoti": len(timu),
		"zongf":exam.Manfen,
		"pass":exam.Pass,
		"fabuid":user.Username,
	})
}
func Fabuks(c *gin.Context)  {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	if user.Qiyepower==1{
		kaochang:=entity.Getkaochang2(user.ID)
		exam:=entity.Getexam2(user.ID)
		bumeng:=entity.Getbumeng(user.Qiyeid)
		c.HTML(http.StatusOK,"qiyeyuangong.html",gin.H{
			"token":token,
			"kaochang":kaochang,
			"exam":exam,
			"bumeng":bumeng,
			"data":gin.H{
				"username":username,
			},
		})
	}else {
		c.HTML(http.StatusOK,"500.html",gin.H{
			"token":token,
			"data":"没有权限",
		})
	}
}
func Fabuks2(c *gin.Context)  {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	if user.Vip==1 || user.Fabunumber>=1{
		kaochang:=entity.Getkaochang2(user.ID)
		exam:=entity.Getexam2(user.ID)
		c.HTML(http.StatusOK,"bufenyonghu.html",gin.H{
			"token":token,
			"kaochang":kaochang,
			"exam":exam,
			"data":gin.H{
				"username":username,
			},
		})
	}else {
		c.HTML(http.StatusOK,"500.html",gin.H{
			"token":token,
			"data":"没有次数",
		})
	}
}
func Zj(c *gin.Context)  {
	id:=c.Query("id")
	bumeng:=entity.Getbumeng2(id)
	users,_:=entity.Getuser4(bumeng.Qiyeid,bumeng.Bumengname)
	var phones string
	phones=""
	for _,v :=range users  {
		phones=phones+v.Phone+","
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"phones":phones,
	})
}
func Ks(c *gin.Context)  {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	//开始但没有结束的考试
	test,_,_:=entity.Startnoend(user.Phone)
	c.HTML(http.StatusOK,"kaoshi.html",gin.H{
		"code":200,
		"token":token,
		"test":test,
		"data":gin.H{
			"username":username,
		},
	})
}
func Ksks(c *gin.Context)  {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user, _ := entity.GetUser(username,password)
	ksid:=c.Query("id")
	ksxq:=c.Query("ksxq")
	ks:=0
	if ksxq=="1" {
		ks=1
	}
	token2:=c.Query("token2")
	test:=entity.Gettest(ksid)
	if user.ID!=test.Fabuid {
		if token2!=test.Token {
			c.HTML(http.StatusOK,"500.html",gin.H{
				"data":"令牌错误",
				"token":token,
			})
			return
		}
	}
	t1,t2:=util.Timediffer(test.Starttime,test.Endtime)
	xq:=c.Query("xq")
	exam:=*new(entity.Exam)
	if xq=="1" {
		exid:=c.Query("id")
		exid2,_:=strconv.Atoi(exid)
		exid3:=int32(exid2)
		exam=entity.Getexam(exid3)
	}else {
		exam=entity.Getexam(test.Shijuanid)
	}

	timu:=entity.Gettimu(exam.Id)
	danx:=make([]entity.Timu,exam.Danxuan)
	duox:=make([]entity.Timu,exam.Duoxuan)
	pand:=make([]entity.Timu,exam.Panduan)
	jian:=make([]entity.Timu,exam.Jianda)
	var n1,n2,n3,n4 int
	var g1,g2,g3,g4 int
	if exam.Danxuan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="danxuan"{
				danx[n1]=timu[i]
				g1+=int(timu[i].Grade)
				n1=n1+1
			}
		}
	}
	if exam.Duoxuan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="duoxuan"{
				duox[n2]=timu[i]
				g2+=int(timu[i].Grade)
				n2++
			}
		}
	}
	if exam.Panduan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="panduan"{
				pand[n3]=timu[i]
				g3+=int(timu[i].Grade)
				n3++
			}
		}
	}
	if exam.Jianda>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="jianda"{
				jian[n4]=timu[i]
				g4+=int(timu[i].Grade)
				n4++
			}
		}
	}
	kcid:=int(test.Kaochangid)
	kcid2:=strconv.Itoa(kcid)
	kc:=entity.Getkaochang(kcid2)
	c.HTML(http.StatusOK,"exam.html",gin.H{
		"code":200,
		"token":token,
		"t1":t1,
		"t2":t2,
		"kc":kc,
		"danxuan":danx,
		"n1":n1,
		"g1":g1,
		"duoxuan":duox,
		"n2":n2,
		"g2":g2,
		"panduan":pand,
		"n3":n3,
		"g3":g3,
		"jianda":jian,
		"n4":n4,
		"g4":g4,
		"examid":exam.Id,
		"ksid":ksid,
		"ksxq":ks,
	})
}
func Jiaojuan(c *gin.Context)  {
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	user,_:=entity.GetUser(username,password)
	token := c.MustGet("token").(string)
	examid:=c.Query("examid")
	s, _ := strconv.Atoi(examid)
	exam:=entity.Getexam(int32(s))
	timu:=entity.Gettimu(exam.Id)
	var ans string
	ans=""
	/**/
	tm:=c.Query("tm")
	zuobinum:=c.PostForm("zuobinum")
	ksid:=c.PostForm("ksid")
	kc:=entity.Getkaochang(ksid)
	/**/
	var def int32
	def=0
	if exam.Danxuan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="danxuan"{
				id:=int(timu[i].Id)
				id2:=strconv.Itoa(id)
				k:=c.PostForm(id2)
				if k==timu[i].Answer{
					def=def+timu[i].Grade
				}
				ans=ans+k+","
			}
		}
	}
	if exam.Duoxuan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="duoxuan"{
				id:=int(timu[i].Id)
				id2:=strconv.Itoa(id)
				k:=c.PostFormArray(id2)
				k1:=""
				for _, v2 := range k {
					ans=ans+v2
					k1=k1+v2
				}
				if k1==timu[i].Answer{
					def=def+timu[i].Grade
				}
				ans=ans+","
			}
		}
	}
	if exam.Panduan>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="panduan"{
				id:=int(timu[i].Id)
				id2:=strconv.Itoa(id)
				k:=c.PostForm(id2)
				if k==timu[i].Answer{
					def=def+timu[i].Grade
				}
				ans=ans+k+","
			}
		}
	}
	if exam.Jianda>=1 {
		for i:=0;i< len(timu); i++ {
			if timu[i].Leixing =="jianda"{
				id:=int(timu[i].Id)
				id2:=strconv.Itoa(id)
				k:=c.PostForm(id2)
				ans=ans+k+","
			}
		}
	}
	answer:=entity.Answer{}
	answer.Answer=ans
	answer.Kaoshiid=exam.Id
	answer.Userid=user.ID
	err:=entity.Creatanswer(&answer)
	/**/
		cj:=new(entity.Chengji)
		cj.Defen=def
		cj.Userphone=user.Phone
		z1,_:=strconv.Atoi(zuobinum)
		z2:=int32(z1)
		if kc.Fangzuobi==1 && kc.Qiehuanmax< z2 {
			cj.Status=1
		}else {
			cj.Status=0
		}
		k1,_:=strconv.Atoi(ksid)
		k2:=int32(k1)
		cj.Kaoshiid=k2
		tes:=entity.Gettest(ksid)
		h,m:=util.Timediffer(tes.Starttime,tes.Endtime)
		ti:=util.Timediffer2(h,m,tm)
		cj.Time=ti
		entity.Createchengji(cj)
	/**/
	if err !=nil {

		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"token":token,
		})
	}

}
func Kcsz(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	user,_:=entity.GetUser(username,password)
	token := c.MustGet("token").(string)
	kc:=entity.Getkaochang2(user.ID)
	c.HTML(http.StatusOK,"shezhiguanli.html",gin.H{
		"code":200,
		"token":token,
		"kc":kc,
		"data":gin.H{
			"username":username,
		},
	})

}
func Kcxq(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	kcid:=c.Query("kcid")
	kc:=entity.Getkaochang(kcid)
	c.HTML(http.StatusOK,"shezhi.html",gin.H{
		"code":200,
		"token":token,
		"kc":kc,
		"data":gin.H{
			"username":username,
		},
	})
}
func Xjkc(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	kc:=entity.Kaochang{}
	c.HTML(http.StatusOK,"shezhi.html",gin.H{
		"code":200,
		"token":token,
		"kc":kc,
		"data":gin.H{
			"username":username,
		},
	})
}
func Tjkc(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	kcname:=c.Query("kcname")
	val1:=c.Query("val1")
	val2:=c.Query("val2")
	val4:=c.Query("val4")
	maxqh:=c.Query("maxqh")
	kc:=new(entity.Kaochang)
	kc.Kaochangname=kcname
	a,_:=strconv.Atoi(val1)
	kc.Fangzuobi=int8(a)
	if a==1 {
		b,_:=strconv.Atoi(maxqh)
		kc.Qiehuanmax=int32(b)
	}else {
		kc.Qiehuanmax=999
	}
	d,_:=strconv.Atoi(val2)
	kc.Fz=int8(d)
	f,_:=strconv.Atoi(val4)
	kc.Number=int32(f)
	kc.Fabuid=user.ID
	entity.Createkaochang(kc)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"token":token,
	})

}
func Sckc(c *gin.Context){
	token := c.MustGet("token").(string)
	kcid := c.Query("kcid")
	kc:=entity.Kaochang{}
	s, _ := strconv.Atoi(kcid)
	kc.Id=int32(s)
	err:=entity.Deletekaochang(kc)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}

}
func Glsj(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	user,_:=entity.GetUser(username,password)
	token := c.MustGet("token").(string)
	sj:=entity.Getexam2(user.ID)
	c.HTML(http.StatusOK,"shijuanguanli.html",gin.H{
		"code":200,
		"token":token,
		"sj":sj,
		"data":gin.H{
			"username":username,
		},
	})
}
func Sjxq(c *gin.Context)  {
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	c.HTML(http.StatusOK,"shijuanguanli.html",gin.H{
		"code":200,
		"token":token,
		"data":gin.H{
			"username":username,
		},
	})
}
func Scsj(c *gin.Context){
	token := c.MustGet("token").(string)
	sjid := c.Query("sjid")
	ex:=entity.Exam{}
	s, _ := strconv.Atoi(sjid)
	ex.Id=int32(s)
	err:=entity.Deleteexam(ex)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}
}
func Drsj(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	c.HTML(http.StatusOK,"daorushijuan.html",gin.H{
		"code":200,
		"token":token,
		"data":gin.H{
			"username":username,
		},
	})
}
func Ksbb(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	test:=entity.End(user.ID)
	c.HTML(http.StatusOK,"baobiao.html",gin.H{
		"code":200,
		"token":token,
		"test":test,
		"data":gin.H{
			"username":username,
		},
	})
}
func Ckbb(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	testid:=c.Query("testid")
	ex:=c.Query("ex")
	test:=entity.Gettest(testid)
	exam:=entity.Getexam(test.Shijuanid)
	cj:=entity.Getchengji(test.Id)
	if ex=="1" {
		util.Export(c,cj)
	}
	cjrs:=entity.Getcount(test.Id)
	tgrs:=entity.Getpass(test.Id,exam.Pass)
	var bb,b float32
	bb=float32(cjrs)
	b=float32(tgrs)
	bb=b/bb
	zhcj:=cj[0].Defen
	zccj:=cj[len(cj)-1].Defen
	var avg int32
	avg=0
	var i int
	i=0
	for ;i<len(cj);i++{
		avg=avg+cj[i].Defen
	}
	var aa,a float32
	aa=float32(avg)
	a=float32(cjrs)
	aa=aa/a

	c.HTML(http.StatusOK,"baobiao2.html",gin.H{
		"token":token,
		"cj":cj,
		"cjrs":cjrs,
		"tgrs":tgrs,
		"tgl":bb,
		"zhcj":zhcj,
		"zccj":zccj,
		"pjcj":aa,
		"testid":testid,
		"data":gin.H{
			"username":username,
		},
	})
}
func Qyyhgl(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	if user.Qiyepower==1{
		users,_:=entity.Getuser5(user.Qiyeid)
		bumengs:=entity.Getbumeng(user.Qiyeid)
		c.HTML(http.StatusOK,"qiyeyonghuguanli.html",gin.H{
			"code":200,
			"token":token,
			"users":users,
			"bumengs":bumengs,
			"qiyeid":user.Qiyeid,
			"data":gin.H{
				"username":username,
			},
		})
	}else {
		c.HTML(http.StatusOK,"500.html",gin.H{
			"token":token,
			"data":"没有权限",
		})
	}

}
func Grzx(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	qiye:=entity.Getqiye(user.Qiyeid)
	c.HTML(http.StatusOK,"gerenzhongxin.html",gin.H{
		"code":200,
		"token":token,
		"user":user,
		"qiye":qiye,
		"data":gin.H{
			"username":username,
		},
	})
}
func Gg(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	gg:=entity.Getgg()
	gg0:=gg[0]
	gg1:=gg[1]
	gg2:=gg[2]
	gg3:=gg[3]
	gg4:=gg[4]
	c.HTML(http.StatusOK,"tongzhi.html",gin.H{
		"code":200,
		"token":token,
		"gg0":gg0,
		"gg1":gg1,
		"gg2":gg2,
		"gg3":gg3,
		"gg4":gg4,
		"data":gin.H{
			"username":username,
		},
	})
}
func Qysz(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	qy:=entity.Getqiye(user.Qiyeid)
	bumeng:=entity.Getbumeng(qy.Qiyeid)
	c.HTML(http.StatusOK,"qiyeshezhi.html",gin.H{
		"code":200,
		"token":token,
		"qy":qy,
		"bm":bumeng,
		"data":gin.H{
			"username":username,
		},
	})
}
func Xzbm(c *gin.Context){
	token := c.MustGet("token").(string)
	qyid:=c.Query("qyid")
	name:=c.Query("name")
	id1,_:=strconv.Atoi(qyid)
	id:=int32(id1)
	bm:=entity.Getbumeng(id)
	i:=0
	var num int32
	for ;i< len(bm); i++ {
		if bm[i].Bumengname==name {
			break
		}
		if bm[i].Bumengid>num {
			num=bm[i].Bumengid
		}
	}
	if i==len(bm) {
		bumeng:=entity.Bumeng{}
		bumeng.Qiyeid=id
		bumeng.Bumengid=num+1
		bumeng.Bumengname=name
		err:=entity.Createbumeng(&bumeng)
		if err==nil {
			c.JSON(http.StatusOK,gin.H{
				"code":200,
				"token":token,
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"code":201,
				"token":token,
			})
		}
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":202,
			"token":token,
		})
	}

}
func Scbm(c *gin.Context){
	token := c.MustGet("token").(string)
	qyid:=c.Query("qyid")
	name:=c.Query("name")
	err:=entity.Deletebumeng(qyid,name)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
			"token":token,
		})
	}
}
func Cxyg(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	bumengs:=entity.Getbumeng(user.Qiyeid)
	bmname:=c.Query("bmname")
	name:=c.Query("name")
	if  name==""{
		users,_:=entity.Getuser6(user.Qiyeid,bmname)
		c.HTML(http.StatusOK,"qiyeyonghuguanli.html",gin.H{
			"code":200,
			"token":token,
			"users":users,
			"bumengs":bumengs,
			"qiyeid":user.Qiyeid,
			"data":gin.H{
				"username":username,
			},
		})
	}else {
		users,_:=entity.Getuser7(user.Qiyeid,bmname,name)
		c.HTML(http.StatusOK,"qiyeyonghuguanli.html",gin.H{
			"code":200,
			"token":token,
			"users":users,
			"bumengs":bumengs,
			"qiyeid":user.Qiyeid,
			"data":gin.H{
				"username":username,
			},
		})
	}

}
func Ygxx(c *gin.Context)  {
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	userid:=c.Query("userid")
	user,_:=entity.Getuser3(userid)
	qiye:=entity.Getqiye(user.Qiyeid)
	c.HTML(http.StatusOK,"yuangongxinxi.html",gin.H{
		"code":200,
		"token":token,
		"user":user,
		"qiye":qiye,
		"data":gin.H{
			"username":username,
		},
	})
}
func Xxxg(c *gin.Context){
	token := c.MustGet("token").(string)
	id:=c.Query("id")
	bm:=c.Query("bm")
	qx:=c.Query("qx")
	qx2,_:=strconv.Atoi(qx)
	err:=entity.Updatauser(id,bm,qx2)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
			"token":token,
		})
	}
}
func Checkp(c *gin.Context){
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	p2:=c.Query("password")
	if password !=p2{
		c.JSON(http.StatusOK,gin.H{
			"code":201,
			"token":token,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"token":token,
		})
	}
}
func Xgmm(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	p2:=c.Query("password")
	err:=entity.Updatapassword(user.ID,p2)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
			"token":token,
		})
	}
}
func Zjcs(c *gin.Context)  {
	token := c.MustGet("token").(string)
	util.Pay(token)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}
func Vip(c *gin.Context)  {
	token := c.MustGet("token").(string)
	util.Pay2(token)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}

func Return(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	out_trade_no:=c.Query("out_trade_no")
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no",out_trade_no)
	req,err:=util.Client.TradeQuery(bm)
	if err!=nil{
		fmt.Println(err)
	}else {
		if req.Response.Msg=="Success" {
			pay:=new(entity.Pay)
			pay.Userphone=user.Phone
			formatStr := "2006-01-02 15:04:05"
			pay.Date=fmt.Sprintln(time.Now().Format(formatStr))
			pay.Money=5
			entity.Createpay(pay)
			err=entity.Updatafabunumber(user.ID,user.Fabunumber+1)
			c.HTML(http.StatusOK,"500.html",gin.H{
				"token":token,
				"data":"请返回首页",
			})
			return
		}
	}
	c.HTML(http.StatusOK,"500.html",gin.H{
		"token":token,
		"data":"交易失败",
	})

}

func Return2(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	out_trade_no:=c.Query("out_trade_no")
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no",out_trade_no)
	req,err:=util.Client.TradeQuery(bm)
	if err!=nil{
		fmt.Println(err)
	}else {
		if req.Response.Msg=="Success" {
			pay:=new(entity.Pay)
			pay.Userphone=user.Phone
			formatStr := "2006-01-02 15:04:05"
			pay.Date=fmt.Sprintln(time.Now().Format(formatStr))
			pay.Money=50
			entity.Createpay(pay)
			err=entity.Updatavip(user.ID)
			c.HTML(http.StatusOK,"500.html",gin.H{
				"token":token,
				"data":"请返回首页",
			})
			return
		}
	}
	c.HTML(http.StatusOK,"500.html",gin.H{
		"token":token,
		"data":"交易失败",
	})

}
func AliPayNotify(c *gin.Context) {
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	//notifyReq, err := alipay.ParseNotifyResult(c.Request) 废弃
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		//数据解析出错
		c.HTML(http.StatusOK,"500.html",gin.H{
			"token":token,
			"data":"支付失败",
		})
		return
	}
	tradeStatus := notifyReq.Get("trade_status")
	if tradeStatus == "TRADE_SUCCESS" {
		//交易成功 处理订单
		entity.Updatafabunumber(user.ID,user.Fabunumber+1)
	}
}
func Czmm(c *gin.Context)  {
	email:=c.Query("email")
	yzm:=c.Query("yzm")
	em:=entity.Getmail2(email,yzm)
	if em.Id==0 {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
			"email":email,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"email":email,
	})
}
func Czm(c *gin.Context){
	email:=c.Query("email")
	c.HTML(http.StatusOK,"chongzhimima.html",gin.H{
		"email":email,
	})
}
func Cm(c *gin.Context){
	email:=c.Query("email")
	password:=c.Query("password")
	err :=entity.Updatepassword2(email,password)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
	}
}
func Ksgl(c *gin.Context){
	username:=c.MustGet("username").(string)
	password:=c.MustGet("password").(string)
	token := c.MustGet("token").(string)
	user,_:=entity.GetUser(username,password)
	test:=entity.Gettest2(user.ID)
	c.HTML(http.StatusOK,"kaoshiguanli.html",gin.H{
		"token":token,
		"test":test,
		"data":gin.H{
			"username":username,
		},
	})
}
func Glks(c *gin.Context){
	testid:=c.Query("id")
	test:=entity.Gettest(testid)
	st:=util.Timediffer4(test.Starttime)
	if st==1 {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
		return
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
	}

}
func Glks2(c *gin.Context){
	username:=c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	testid:=c.Query("id")
	test:=entity.Gettest(testid)
	usersphone:=test.Phones
	up:=strings.Split(usersphone,",")
	users:=entity.Getuser9(up)
	c.HTML(http.StatusOK,"kaoshengguanli.html",gin.H{
		"code":200,
		"token":token,
		"testid":testid,
		"users":users,
		"data":gin.H{
			"username":username,
		},
	})
}
type Login struct {
	Ksid     string `form:"ksid" json:"ksid"`
	Testid string `form:"testid" json:"testid"`
}

func Scks(c *gin.Context){
	var login Login
	if err := c.ShouldBindJSON(&login); err == nil {
		test:=entity.Gettest(login.Testid)
		test.Phones=strings.ReplaceAll(test.Phones,login.Ksid+",","")
		entity.Updatetest(login.Testid,test.Phones)
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"testid":login.Testid,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
		})
	}
}
type Login5 struct {
	Phones     string `form:"phones" json:"phones"`
	Testid string `form:"testid" json:"testid"`
}
func Tjks(c *gin.Context){
	var login Login5
	if err := c.ShouldBindJSON(&login); err == nil {
		fmt.Println(login)
		up:=strings.Split(login.Phones,",")
		users:=entity.Getuser9(up)
		var phones=""
		for _,v := range users{
			phones=phones+v.Phone+","
		}
		test:=entity.Gettest(login.Testid)
		test.Phones=test.Phones+phones
		entity.Updatetest(login.Testid,test.Phones)
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"testid":login.Testid,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
		})
	}
}
func Pysj(c *gin.Context){
	testid:=c.Query("id")
	test:=entity.Gettest(testid)
	st:=util.Timediffer4(test.Endtime)
	if st==1 { //已经结束
		c.JSON(http.StatusOK,gin.H{
			"code":200,
		})
		return
	} else { //还没有结束
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
	}
}
func Pysj2(c *gin.Context){
	username := c.MustGet("username").(string)
	token := c.MustGet("token").(string)
	testid:=c.Query("id")
	li:=c.Query("li")
	li2,_:=strconv.Atoi(li)
	test:=entity.Gettest(testid)
	exam:=entity.Getexam(test.Shijuanid)
	num1:=exam.Danxuan+exam.Duoxuan+exam.Panduan
	num:=exam.Jianda
	timus:=entity.Getjianda(exam.Id)
	uans:=entity.Getans(testid,li)
	uansss:=strings.Split(uans.Answer,",")
	uanss:=make([]string,num)
	for i:=int(num1);i<len(uansss)-1;i++ {
		uanss[i-int(num1)]=uansss[i]
	}
	c.HTML(http.StatusOK,"piyueshijuan.html",gin.H{
		"token":token,
		"testid":testid,
		"timus":timus,
		"uans":uans,
		"uanss":uanss,
		"li":li2+1,
		"data":gin.H{
			"username":username,
		},
	})
}
func Jiaf(c *gin.Context)  {
	//token := c.MustGet("token").(string)
	userid:=c.Query("userid")
	testid:=c.Query("testid")
	fs:=c.Query("fs")
	g1:=c.Query("g1")
	user,_:=entity.Getuser3(userid)
	cj:=entity.Getchengji2(user.Phone,testid)
	fs1,_:=strconv.Atoi(fs)
	fs2:=int32(fs1)
	g11,_:=strconv.Atoi(g1)
	g12:=int32(g11)
	if fs2<0 || fs2 >g12 {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
		return
	}
	newfs:=cj.Defen+fs2
	err:=entity.Updatachengji(user.Phone,testid,newfs)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":201,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
	})
}
func Kssc(c *gin.Context){
	token := c.MustGet("token").(string)
	testid:=c.Query("testid")
	entity.Deletetest(testid)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"token":token,
	})
}