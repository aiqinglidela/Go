package entity

import "TianCheng/model"

type Tokeninvalid struct {
	Id int32 `json:"Id"`
	Token string `json:"Token"`
	Userphone string `json:"Token"`
}

func Gettoken(token1 string)(token2 Tokeninvalid)  {
	model.DB.Debug().Where("token=?",token1).Find(&token2)
	return
}
func Deletetoken(userphone string)  {
	model.DB.Debug().Where("userphone=?",userphone).Delete(Tokeninvalid{})
}
func Createtoken(token *Tokeninvalid)  {
	model.DB.Create(&token)
}