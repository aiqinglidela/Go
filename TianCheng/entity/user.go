package entity

import (
	"TianCheng/model"
	"fmt"
)

/*用户表*/
type User struct {
	ID int32 `json:"Id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
	Name string `json:"Name"`
	Sex string `json:"Sex"`
	Number string `json:"Number"`
	Qiyeid int32 `json:"Qiyeid"`
	Qiyepower int32 `json:"QiyePower"`
	Vip int8 `json:"Vip"`
	Bumengname string `json:"BuMengName"`
	Fabunumber int32 `json:"FaBuNumber"`
	Ruzhidate string `json:"RuZhiDate"`
	Emailstatus int `json:"Emailstatus"`
	Phonestatus int `json:"Phonestatus"`
}
/*检查用户名*/
func Checkusername(username string) (err error){
	user := new(User)
	err = model.DB.Debug().Where("username=?", username).First(user).Error
	fmt.Println(user)
	return err
}
/*检查邮箱*/
func Checkemail(email string) (err error){
	user := new(User)
	err = model.DB.Debug().Where("email=? and emailstatus=0", email).First(user).Error
	fmt.Println(user)
	return err
}
/*检查手机号*/
func Checkphone(phone string) (err error){
	user := new(User)
	err = model.DB.Debug().Where("phone=? and phonestatus=0", phone).First(user).Error
	fmt.Println(user)
	return err
}
/*登录*/
func GetUser(username,password string)(user User, err error){
	err=model.DB.Debug().Where("(username=? and password=? ) or ( email=? and password=?)", username,password,username,password).First(&user).Error
	return
}
/*登录*/
func Getuser2(phone string)(user User, err error){
	err = model.DB.Debug().Where("phone=?",phone).First(&user).Error
	return
}
/*注册*/
func CreateUser(user User) (err error){
	err = model.DB.Create(&user).Error
	return
}
/*根据id获取user*/
func Getuser3(id string) (user User, err error){
	err = model.DB.Debug().Where("id=?",id).First(&user).Error
	return
}
/*根据企业id和部门名称获取user*/
func Getuser4(id int32,name string) (user []User, err error){
	if err = model.DB.Debug().Where("qiyeid=? AND bumengname=?",id,name).Find(&user).Error; err!=nil{
		return nil, err
	}
	return
}
/*根据企业id获取user*/
func Getuser5(id int32) (user []User, err error){
	if err = model.DB.Debug().Where("qiyeid=?",id).Find(&user).Error; err!=nil{
		return nil, err
	}
	return
}
func Getuser6(id int32,bmname string)(user []User, err error){
	if err = model.DB.Debug().Where("qiyeid=? and bumengname=?",id,bmname).Find(&user).Error; err!=nil{
		return nil, err
	}
	return
}
func Getuser7(id int32,bmname,name string)(user []User, err error){
	if err = model.DB.Debug().Where("qiyeid=? and bumengname=? and username like ?",id,bmname,"%"+name+"%").Find(&user).Error; err!=nil{
		return nil, err
	}
	return
}
func Updatauser(id,bm string,qx int)(err error)  {
	err=model.DB.Debug().Model(&User{}).Where("id=?",id).Update("bumengname",bm).Error
	err=model.DB.Debug().Model(&User{}).Where("id=?",id).Update("qiyepower",qx).Error
	return
}
func Deleteuser(id string){
	model.DB.Debug().Where("id=?",id).Delete(&User{})
}
func Updatapassword(id int32,password string)(err error){
	err=model.DB.Debug().Model(&User{}).Where("id=?",id).Update("password",password).Error
	return
}
func Updatafabunumber(id,num int32)(err error){
	err=model.DB.Debug().Model(&User{}).Where("id=?",id).Update("fabunumber",num).Error
	return
}
func Updatavip(id int32)(err error){
	err=model.DB.Debug().Model(&User{}).Where("id=?",id).Update("vip",1).Error
	return
}
/*检查邮箱*/
func Updatepassword2(email,password string) (err error){
	err=model.DB.Debug().Model(&User{}).Where("email=?",email).Update("password",password).Error
	return
}
func Getuser9(phones []string)(users []User){
	model.DB.Debug().Where("phone in (?)", phones).Find(&users)
	return
}
func Getuser10()(count int){
	model.DB.Debug().Table("users").Count(&count)
	return
}
func Countuser2(d1,d2 string)(users []User,count int){
	model.DB.Debug().Where("ruzhidate <= ? and ruzhidate >= ? ",d1,d2).Find(&users).Count(&count)
	return
}
func Getuser11()(users []User){
	model.DB.Debug().Find(&users)
	return
}
func Getuser12(name,ysy string)(users []User){
	if ysy =="1" {
		model.DB.Debug().Where("username like ?","%"+name+"%").Find(&users)
		return
	}
	if ysy=="2" {
		model.DB.Debug().Where("phone like ?","%"+name+"%").Find(&users)
		return
	}
	if ysy=="3" {
		model.DB.Debug().Where("email like ?","%"+name+"%").Find(&users)
		return
	}
	return
}
func Getuser13()(users []User){
	model.DB.Debug().Where("qiyepower=1").Find(&users)
	return
}
func Getuser14(qiyeids []int32)(users []User){
	model.DB.Debug().Where("qiyeid in (?)",qiyeids).Find(&users)
	return
}
func Getuser15(glyname,glyphone,glyemail,glynumber string)(users User){
	model.DB.Debug().Where("username = ? or phone = ? or email = ? or number=?",glyname,glyphone,glyemail,glynumber).First(&users)
	return
}
func Createuser2(user *User){
	model.DB.Create(&user)
	return
}
func Deleteuser2(id string){
	model.DB.Debug().Model(&User{}).Where("id=?",id).Update("emailstatus",1)
	model.DB.Debug().Model(&User{}).Where("id=?",id).Update("phonestatus",1)
}
func Jiechu(id string){
	model.DB.Debug().Model(&User{}).Where("id=?",id).Update("emailstatus",0)
	model.DB.Debug().Model(&User{}).Where("id=?",id).Update("phonestatus",0)
}