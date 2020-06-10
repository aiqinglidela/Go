package entity

import "TianCheng/model"

/*企业表*/
type Qiye struct {
	Id int32 `json:"Id"`
	Qiyeid int32 `json:"QiyeId"`
	Qiyename string `json:"QiyeName"`
}
func Getqiye(id int32)(qiye Qiye){
	model.DB.Debug().Where("qiyeid=?",id).Find(&qiye)
	return
}
func Getqiye2()(qys []Qiye){
	model.DB.Debug().Find(&qys)
	return
}
func Getqiye3(qyname string)(qys []Qiye){
	model.DB.Debug().Where("qiyename like ?","%"+qyname+"%").Find(&qys)
	return
}
func Getqiye4(qyname string)(qys Qiye){
	model.DB.Debug().Where("qiyename = ?",qyname).First(&qys)
	return
}
func Createqiye(qy *Qiye){
	model.DB.Create(&qy)
	return
}

