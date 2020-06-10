package entity

import (
	"TianCheng/model"
	"fmt"
)

/*成绩表*/
type Chengji struct {
	Id int32 `json:"Id"`
	Kaoshiid int32 `json:"KaoshiID"`
	Defen int32 `json:"DeFen"`
	Userphone string `json:"UserId"`
	Time int32 `json:"Time"`
	Status int32 `json:"Status"`
}
/*统计用户通过的考试数量*/
func Pass(phone string)(count int){
	user,_:= Getuser2(phone)
	model.DB.Where("userphone = ? and defen >=60", user.Phone).Find(&Chengji{}).Count(&count)
	return count
}
/*统计用户未通过的考试数量*/
func Fail(phone string)(count int){
	user,_:= Getuser2(phone)
	model.DB.Where("userphone = ? and defen < 60 ", user.Phone).Find(&Chengji{}).Count(&count)
	return
}
func Getchengji(testid int32)(chengjis []Chengji){
	model.DB.Debug().Where("kaoshiid=?",testid).Order("defen desc").Find(&chengjis)
	return
}
func Getcount(testid int32)(count int){
	model.DB.Where("kaoshiid = ?", testid).Find(&Chengji{}).Count(&count)
	return
}
func Getpass(testid,pass int32)(count int){
	model.DB.Where("kaoshiid = ? AND defen>=?", testid,pass).Find(&Chengji{}).Count(&count)
	return
}
func Createchengji(cj *Chengji)  {
	model.DB.Create(&cj)
}
func Getchengji2(phone,testid string)(cj Chengji){
	model.DB.Where("kaoshiid = ? AND userphone=?", testid,phone).Find(&cj)
	return
}
func Updatachengji(phone,testid string,fs int32)(err error){
	fmt.Println(phone,testid,fs,"2222222222222")
	err=model.DB.Debug().Model(&Chengji{}).Where("userphone=? and kaoshiid=? ",phone,testid).Update("defen",fs).Error
	return
}