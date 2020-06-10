package entity

import (
	"TianCheng/model"
)

/*考场表*/
type Kaochang struct {
	Id int32 `json:"Id"`
	Kaochangname string `json:"KaoChangName"`
	Fangzuobi int8 `json:"FangZuoBi"`
	Qiehuanmax int32 `json:"QieHuanMax"`
	Fz int8 `json:"FZ"`
	Number int32 `json:"Number"`
	Fabuid int32 `json:"Fabuid"`
}
func Getkaochang(id string)(kaochang Kaochang){
	model.DB.Debug().Where("id=?",id).Find(&kaochang)
	return
}
func Getkaochang2(id int32)(kaochang []Kaochang){
	model.DB.Debug().Where("fabuid=?",id).Find(&kaochang)
	return
}
func Getkaochang3(id int32)(kaochang Kaochang){
	model.DB.Debug().Where("fabuid=?",id).Find(&kaochang)
	return
}
func Deletekaochang(kc Kaochang)(err error){
	err=model.DB.Delete(&kc).Error
	return
}
func Createkaochang(kc *Kaochang){
	model.DB.Create(&kc)
}