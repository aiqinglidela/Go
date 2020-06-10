package entity

import "TianCheng/model"

/*题目表*/
type Timu struct {
	Id int32 `json:"Id"`
	Title string `json:"Title"`
	Optiona string `json:"OptionA"`
	Optionb string `json:"OptionB"`
	Optionc string `json:"OptionC"`
	Optiond string `json:"OptionD"`
	Optione string `json:"OptionE"`
	Grade int32 `json:"Grade"`
	Leixing string `json:"Leixing"`
	Answer string `json:"Answer"`
	Shijuanid int32 `json:"ShijuanId"`
}
func Gettimu(id int32)(timu []Timu){
	model.DB.Debug().Where("Shijuanid=?",id).Find(&timu)
	return
}
func Createtime(timu *Timu)(err error){
	err = model.DB.Create(&timu).Error
	return
}
func Getjianda(id int32)(timu []Timu){
	model.DB.Debug().Where("Shijuanid=? and leixing='jianda' ",id).Find(&timu)
	return
}
