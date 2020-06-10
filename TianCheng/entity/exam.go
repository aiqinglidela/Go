package entity

import (
	"TianCheng/model"
)

/*试卷表*/
type Exam struct {
	Id int32 `json:"Id"`
	Examname string `json:"ExamName"`
	Danxuan int32 `json:"DanXuan"`
	Duoxuan int32 `json:"DuoXuan"`
	Panduan int32 `json:"PanDuan"`
	Jianda int32 `json:"JianDan"`
	Manfen int32 `json:"ManFen"`
	Pass int32 `json:"Pass"`
	Fabuid int32 `json:"Fabuid"`
}

func Getexam(id int32)(exam Exam)  {
	model.DB.Debug().Where("id=?",id).Find(&exam)
	return
}
func Getexam2(id int32)(exam []Exam)  {
	model.DB.Debug().Where("fabuid=?",id).Find(&exam)
	return
}
func Deleteexam(ex Exam)(err error){
	err=model.DB.Delete(&ex).Error
	return
}
func Createexam(exam *Exam)(err error){
	err = model.DB.Create(&exam).Error
	return
}
func Getexam3(name string)(exam Exam){
	model.DB.Debug().Where("examname=?",name).Find(&exam)
	return
}
