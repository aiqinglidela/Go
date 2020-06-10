package entity

import "TianCheng/model"

/*消息表*/
type Mail struct {
	Id int32 `json:"Id"`
	Aceptemail string `json:"AceptId"`
	Title string `json:"Title"`
	Message string `json:"Message"`
	Date string `json:"Date"`
}
func Createmail(em *Mail)(err error){
	err = model.DB.Create(&em).Error
	return
}
func Deletemail(email string){
	model.DB.Debug().Where("aceptemail=?",email).Delete(&Mail{})
}
func Getmail(email string)(mail Mail){
	model.DB.Debug().Where("aceptemail=?",email).Find(&mail)
	return
}
func Getmail2(email,yzm string)(mail Mail)  {
	model.DB.Debug().Where("aceptemail=? and message= ?",email,yzm).Find(&mail)
	return
}