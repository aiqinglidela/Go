package entity

import (
	"TianCheng/model"
	"time"
)

/*考试表*/
type Test struct {
	Id int32 `json:"Id"`
	Shijuanid int32 `json:"ShijuanId"`
	Kaochangid int32 `json:"KaochangId"`
	Starttime string `json:"StartTime"`
	Endtime string `json:"EndTime"`
	Token string `json:"Token"`
	Phones string `json:"Phones"`
	Kaoshiname string `json:"KaoshiName"`
	Fabuid int32 `json:"Fabuid"`
}
/*统计用户考试数量*/
func Countuser(phone string)(tests []Test,count int,err error){
	model.DB.Debug().Where("phones like ?","%"+phone+"%").Find(&tests).Count(&count)
	return
}
/*用户还未开始的考试*/
func Nostart(phone string)(tests []Test,count int,err error){
	date:=time.Now()
	model.DB.Debug().Where("phones like ? and starttime > ?","%"+phone+"%",date).Find(&tests).Count(&count)
	return
}
/*用户开始没有结束的考试*/
func Startnoend(phone string)(tests []Test,count int,err error){
	date:=time.Now()
	model.DB.Debug().Where("phones like ? and starttime < ? and endtime > ?","%"+phone+"%",date,date).Find(&tests).Count(&count)
	return
}
/*结束的考试*/
func End(id int32)(test []Test)  {
	date:=time.Now()
	model.DB.Debug().Where("fabuid=? and endtime < ?",id,date).Find(&test)
	return
}
func Gettest(id string)(tests Test){
	model.DB.Debug().Where("id=?",id).Find(&tests)
	return
}
func Creattest(test Test) (err error){
	err = model.DB.Create(&test).Error
	return
}
func Gettest2(id int32)(test []Test){
	model.DB.Debug().Where("fabuid=? ",id).Find(&test)
	return
}
func Updatetest(id,phone string){
	model.DB.Debug().Model(&Test{}).Where("id=?",id).Update("phones",phone)
	return
}
func Deletetest(testid string){
	model.DB.Debug().Where("id=?",testid).Delete(&Test{})
}
func Gettest3()(count int){
	model.DB.Debug().Table("tests").Count(&count)
	return
}