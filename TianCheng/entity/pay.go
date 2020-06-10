package entity

import (
	"TianCheng/model"
	"strings"
)
/*企业表*/
type Pay struct {
	Id int32 `json:"Id"`
	Userphone string `json:"userphone"`
	Money int `json:"money"`
	Date string `json:"date"`
}

func Createpay(pay *Pay){
	model.DB.Create(&pay)
}
func Getpay(date string)(pays []Pay){
	da:=strings.Split(date," ")
	da1:=da[0]+" 00:00:00"
	model.DB.Debug().Where("date>? ",da1).Find(&pays)
	return
}
func Countpay(d1,d2 string)(pays []Pay,err error){
	err=model.DB.Debug().Where("date<=? and date >= ? ",d1,d2).Find(&pays).Error
	return
}