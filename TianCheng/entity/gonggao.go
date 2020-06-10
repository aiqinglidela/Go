package entity

import "TianCheng/model"

/*公告表*/
type Gonggao struct {
	Id int32 `json:"Id"`
	Title string `json:"Title"`
	Text string `json:"Text"`
}
func Getgg()(gg []Gonggao){
	model.DB.Debug().Limit(5).Find(&gg)
	return
}
func Gettz()(gg []Gonggao){
	model.DB.Debug().Find(&gg)
	return
}
func Gettz2(id string)(gg Gonggao){
	model.DB.Debug().Where("id = ?",id).Find(&gg)
	return
}
func Updatatz(id,title,text string){
	model.DB.Debug().Model(&Gonggao{}).Where("id=?",id).Update("title",title)
	model.DB.Debug().Model(&Gonggao{}).Where("id=?",id).Update("text",text)
}
func Createtz(tz *Gonggao)  {
	model.DB.Create(&tz)
}
func Gettz3(title string)(ggs []Gonggao){
	model.DB.Debug().Where("title like ? or text like ?","%"+title+"%","%"+title+"%").Find(&ggs)
	return
}
func Deletetz(id string){
	model.DB.Debug().Where("id=?",id).Delete(&Gonggao{})
}