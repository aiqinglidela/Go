package util

import (
	"TianCheng/entity"
	"baliance.com/gooxml/document"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Exam(path string)(n1,n2,n3,n4,n5,n6,n7,n8,n9 int){
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	var line string
	/*行*/
	for j:=0;j< len(doc.Paragraphs());j++{
		line=""
		/*行中片段*/
		for _, run := range doc.Paragraphs()[j].Runs() {
			line=line+run.Text()
		}
		if strings.HasPrefix(line, "一.") {
			i:=strings.Index(line,"(")
			j:=strings.Index(line,")")
			arr:=strings.Split(line[i+1:j],",")
			n1,_=strconv.Atoi(arr[0])
			n2,_=strconv.Atoi(arr[1])
			n9=n9+n1*n2
		}
		if strings.HasPrefix(line, "二.") {
			i:=strings.Index(line,"(")
			j:=strings.Index(line,")")
			arr:=strings.Split(line[i+1:j],",")
			n3,_=strconv.Atoi(arr[0])
			n4,_=strconv.Atoi(arr[1])
			n9=n9+n3*n4
		}
		if strings.HasPrefix(line, "三.") {
			i:=strings.Index(line,"(")
			j:=strings.Index(line,")")
			arr:=strings.Split(line[i+1:j],",")
			n5,_=strconv.Atoi(arr[0])
			n6,_=strconv.Atoi(arr[1])
			n9=n9+n5*n6
		}
		if strings.HasPrefix(line, "四.") {
			i:=strings.Index(line,"(")
			j:=strings.Index(line,")")
			arr:=strings.Split(line[i+1:j],",")
			n7,_=strconv.Atoi(arr[0])
			n8,_=strconv.Atoi(arr[1])
			n9=n9+n7*n8
		}

	}
return
}
func Exam2(path string) (timu [100]entity.Timu){
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	var line string
	k:=0
	/*行*/
	for j:=0;j< len(doc.Paragraphs());j++ {
		line = ""
		/*行中片段*/
		for _, run := range doc.Paragraphs()[j].Runs() {
			line = line + run.Text()
		}
		a:=strings.HasPrefix(line, "A.")
		b:=strings.HasPrefix(line, "B.")
		c:=strings.HasPrefix(line, "C.")
		d:=strings.HasPrefix(line, "D.")
		e:=strings.HasPrefix(line, "答案:")
		f:=strings.HasPrefix(line, " ")
		g:=strings.HasPrefix(line, "A.")
		h1:=strings.HasPrefix(line, "一.")
		h2:=strings.HasPrefix(line, "二.")
		h3:=strings.HasPrefix(line, "三.")
		h4:=strings.HasPrefix(line, "四.")
		if a||b||c||d||e||f||g||h1||h2||h3||h4{
			if a{
				timu[k].Optiona=fmt.Sprintf(line[2:])
			}
			if b{
				timu[k].Optionb=fmt.Sprintf(line[2:])
			}
			if c{
				timu[k].Optionc=fmt.Sprintf(line[2:])
			}
			if d{
				timu[k].Optiond=fmt.Sprintf(line[2:])
			}
			if e {
				timu[k].Answer=fmt.Sprintf(line[7:])
			}
			if f {
				k++
			}
		}else {
			timu[k].Title=fmt.Sprintf(line[2:])
		}
	}
	return
}
func Drks(path string) (ks string) {
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	var line string
	line = ""
	/*行*/
	for j:=0;j< len(doc.Paragraphs());j++ {
		/*行中片段*/
		for _, run := range doc.Paragraphs()[j].Runs() {
			line = line + run.Text()
		}
		line=line+","
	}
	return line
}