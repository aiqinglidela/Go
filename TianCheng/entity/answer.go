package entity

import "TianCheng/model"

/*考生答案表*/
type Answer struct {
	Id int32 `json:"Id"`
	Kaoshiid int32 `json:"KaoshiId"`
	Userid int32 `json:"UserId"`
	Answer string `json:"Answer"`
}
func Creatanswer(answer *Answer)(err error){
	err = model.DB.Create(&answer).Error
	return
}
func Getans(id,li string)(ans Answer){
	model.DB.Debug().Limit(1).Offset(li).Where("kaoshiid=?",id).Find(&ans)
	return
}
