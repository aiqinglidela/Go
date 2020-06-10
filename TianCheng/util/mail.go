package util

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"net/smtp"
)


//使用第三方库gomail实现邮件的发送
//更多了解，请前往https://godoc.org/gopkg.in/gomail.v2
func SendMailByGomail(mailto,body string ){
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "aiqinglideyali@126.com", "wangkang")
	m.SetHeader("To", mailto)
	m.SetHeader("Subject", "天秤平台验证码")
	//m.Embed("c://1.jpg")
	m.SetBody("text/html", "验证码为以下六位："+body)

	d := gomail.NewDialer("smtp.126.com",465, "aiqinglideyali@126.com", "KJWAHURASUXQNLMB")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
//使用go自带的net/smtp包实现邮件的发送
//更多了解，请前往https://godoc.org/net/smtp
func SendMailByNetSmtp(){
	auth := smtp.PlainAuth("", "xxx@163.com", "***", "smtp.163.com")

	to := []string{"xxx@163.com"}
	image,_:=ioutil.ReadFile("c://1.jpg")
	imageBase64:=base64.StdEncoding.EncodeToString(image)
	msg := []byte("from:xxx@163.com\r\n"+
		"to: xxx@163.com\r\n" +
		"Subject: hello,subject!\r\n"+
		"Content-Type:multipart/mixed;boundary=a\r\n"+
		"Mime-Version:1.0\r\n"+
		"\r\n" +
		"--a\r\n"+
		"Content-type:text/plain;charset=utf-8\r\n"+
		"Content-Transfer-Encoding:quoted-printable\r\n"+
		"\r\n"+
		"此处为正文内容!\r\n"+
		"--a\r\n"+
		"Content-type:image/jpg;name=1.jpg\r\n"+
		"Content-Transfer-Encoding:base64\r\n"+
		"\r\n"+
		imageBase64+"\r\n"+
		"--a--\r\n")
	err := smtp.SendMail("smtp.163.com:25", auth, "xxx@163.com", to, msg)
	if err != nil {
		fmt.Println(err)
	}
}
