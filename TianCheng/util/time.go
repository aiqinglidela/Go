package util

import (
	"strconv"
	"strings"
	"time"
)

func Timediffer(s1,s2 string) (h,m int) {
	var ss,s string
	s=""
	ss=""
	for i:=0;i<len(s1);i++{
		if s1[i]<='9'&&s1[i]>='0'{
			s=s+string(s1[i])
		}
	}
	for i:=0;i<len(s2);i++{
		if s2[i]<='9'&&s2[i]>='0'{
			ss=ss+string(s2[i])
		}
	}
	s11,_:=strconv.Atoi(s)
	s22,_:=strconv.Atoi(ss)
	sss:=s22-s11
	sss=sss/100
	h=sss/60
	m=sss-h*60
	return h,m
}
func Timediffer2(h,m int,tm string)(t int32)  {
	tm1:=tm[:2]
	tm2:=tm[3:5]
	tm11,_:=strconv.Atoi(tm1)
	tm22,_:=strconv.Atoi(tm2)
	s1:=h*60+m
	s2:=tm11*60+tm22
	s3:=s1-s2
	s4:=int32(s3)
	return s4
}
func Timediffer3(s1 string) (h int) {
	stringTime1 := s1
	data:=time.Now().String()
	stringTime2 := data[:19]
	loc, _ := time.LoadLocation("Local")
	the_time1, err1 := time.ParseInLocation("2006-01-02 15:04:05", stringTime1, loc)
	the_time2, err2 := time.ParseInLocation("2006-01-02 15:04:05", stringTime2, loc)
	if err1 == nil && err2==nil {
		unix_time1 := the_time1.Unix() //1504082441
		unix_time2 := the_time2.Unix() //1504082441
		t:=unix_time2-unix_time1
		if t/60>5 {
			//超过五分钟
			return 1
		}else {
			//没有超过5分钟
			return 0
		}
	}
	return 1
}
func Timediffer4(s1 string) (h int) {
	stringTime1 := s1
	data:=time.Now().String()
	stringTime2 := data[:19]
	loc, _ := time.LoadLocation("Local")
	the_time1, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime1, loc)
	the_time2, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime2, loc)
	unix_time1 := the_time1.Unix() //1504082441
	unix_time2 := the_time2.Unix() //1504082441
	t:=unix_time2-unix_time1
	if t>0 {
		return 1//已经开始 //已经结束
	}
	return 0//还未开始
}
func Timediffer5(s1 string) (d1,d2,d3,d4,d5,d6,d7 string) {
	da:=strings.Split(s1," ")
	da1:=da[0]+" 00:00:00"
	loc, _ := time.LoadLocation("Local")
	the_time1, _ := time.ParseInLocation("2006-01-02 15:04:05", da1, loc)
	unix_time1 := the_time1.Unix() //1504082441
	daa1:=unix_time1-86400
	daa2:=unix_time1-86400*2
	daa3:=unix_time1-86400*3
	daa4:=unix_time1-86400*4
	daa5:=unix_time1-86400*5
	daa6:=unix_time1-86400*6
	daa7:=unix_time1-86400*7
	d1 = time.Unix(daa1, 0).Format("2006-01-02 15:04:05")
	d2 = time.Unix(daa2, 0).Format("2006-01-02 15:04:05")
	d3 = time.Unix(daa3, 0).Format("2006-01-02 15:04:05")
	d4 = time.Unix(daa4, 0).Format("2006-01-02 15:04:05")
	d5 = time.Unix(daa5, 0).Format("2006-01-02 15:04:05")
	d6 = time.Unix(daa6, 0).Format("2006-01-02 15:04:05")
	d7 = time.Unix(daa7, 0).Format("2006-01-02 15:04:05")
	return
}