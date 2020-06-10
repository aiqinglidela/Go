package util

import (
	"TianCheng/entity"
	//"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func Export(c *gin.Context,cj []entity.Chengji){
	// 列标题
	//排名	姓名	成绩	计时	作弊
	titles := []string{
		"排名","姓名","成绩","计时","作弊",
	}
	// 数据源

	data := []map[string]interface{}{}

	for i:=0;i<len(cj);i++{
		m1:= map[string]interface{}{}
		m1["No"]=i+1
		m1["name"]=cj[i].Userphone
		m1["cj"]=cj[i].Defen
		m1["time"]=cj[i].Time
		m1["zb"]=cj[i].Status
		data=append(data,m1)
	}
	//fmt.Println(data[0])

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	for clumnNum,v := range titles {
		sheetPosition := Div(clumnNum+1)+"1"
		f.SetCellValue("Sheet1", sheetPosition,v)
	}
	for vv,v := range data {
		r:=strconv.Itoa(vv+2)
		f.SetCellValue("Sheet1","A"+r,v["No"])
		f.SetCellValue("Sheet1","B"+r,v["name"])
		f.SetCellValue("Sheet1","C"+r,v["cj"])
		f.SetCellValue("Sheet1","D"+r,v["time"])
		f.SetCellValue("Sheet1","E"+r,v["zb"])
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	name:="File/"
	n1:=rand.Intn(100)
	name=name+strconv.Itoa(n1)+".xlsx"
	// Save xlsx file by the given path.
	if err := f.SaveAs(name); err != nil {
		println(err.Error())
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"成绩分析.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(name)
}
// Div 数字转字母
func Div(Num int)  string{
	var(
		Str string = ""
		k int
		temp []int   //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O",
		"P","Q","R","S","T","U","V","W","X","Y","Z"}

	if Num >26 {  //数据大于26需要进行拆分
		for {
			k = Num % 26  //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26{   //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	}else{
		return Slice[Num]
	}
	for _,value := range temp{
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}