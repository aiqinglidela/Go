package entity

import "TianCheng/model"

/*部门表*/
type Bumeng struct {
	Qiyeid int32 `json:"QiyeId"`
	Bumengid int32 `json:"BuMengId"`
	Bumengname string `json:"BuMengName"`
}
func Getbumeng(id int32)(bumeng []Bumeng){
	model.DB.Debug().Where("qiyeid=?",id).Find(&bumeng)
	return
}
func Getbumeng2(id string)(bumeng Bumeng){
	model.DB.Debug().Where("bumengid=?",id).Find(&bumeng)
	return
}
func Createbumeng(bm *Bumeng)(err error)  {
	err=model.DB.Create(&bm).Error
	return
}
func Deletebumeng(id,name string)(err error){
	err=model.DB.Where("qiyeid=? and bumengname=?",id,name).Delete(Bumeng{}).Error
	return
}