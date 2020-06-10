package util

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
	"os/exec"
	"strings"
	"time"
)

var (
	appid           = "2016102200739729"
	aliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnXGp2mOyjqSU0znPat9meHiWYs3sSgEejTStAS4G90oBkouYGZF/IjAPQ8RV/nUwo4pSBLqh1o5PMbBu20ipGZqPDiOkyhSIC+s6cu3eRqKirRCoF+5UNXqrWTkMYpa6HgI3OljhauAwbSvBIqrpqWrc+VX2caUvIDu20WT2Q6PBVDYUnhH+MM+BHApQo0Rgh8RhtseQSV2vDLw+j7/sx8658frbKfyKqwgJBppuNWT3S2pT+BX/ok59vLS34y/9ZrI3upoDOnZJX9MoIeLrrNLMIGywZXhGjiH0bC7UxdEyXTXVD4K8yHxIQIY5/Q0NwA0wj3m/R5WLW5XID6aduwIDAQAB"
	privateKey      = "MIIEowIBAAKCAQEAnXGp2mOyjqSU0znPat9meHiWYs3sSgEejTStAS4G90oBkouYGZF/IjAPQ8RV/nUwo4pSBLqh1o5PMbBu20ipGZqPDiOkyhSIC+s6cu3eRqKirRCoF+5UNXqrWTkMYpa6HgI3OljhauAwbSvBIqrpqWrc+VX2caUvIDu20WT2Q6PBVDYUnhH+MM+BHApQo0Rgh8RhtseQSV2vDLw+j7/sx8658frbKfyKqwgJBppuNWT3S2pT+BX/ok59vLS34y/9ZrI3upoDOnZJX9MoIeLrrNLMIGywZXhGjiH0bC7UxdEyXTXVD4K8yHxIQIY5/Q0NwA0wj3m/R5WLW5XID6aduwIDAQABAoIBAGLxCB2/I3nywbgQD2IJLFsu5MsIv6IuG+X/IW2pYCVNc7mFUjtXLP7uFtmua6AyQzOy/CIOqx+rZOdEChZyDhqR/66SBcAcsPZUh/GrvFDFade6rPgFBRRql33FlRvrvuCW4eAf9OD2LFie84t2vsqB5oYmWG5IsVnfTN9ARGZyaPOwYSEcFBto+s7HeJnGX8dqwWWVhtG9z/PgSeO49bGkfigoF5QP4ajA09TgEjRZN0Wh67KFp+N80rzP+sV7JKBMzq/P9npFVIsT3I7vQg73VtA3siPg+6sh2towCrvzDbJQ0uLhtmBoK1Fe0SO2CfiHmuLYWh/TAuQqCPog0WkCgYEAy6CMaWrzSNTNPeyde5aLtESggDK45LPzfVgR09AI5NrNZGfLs16nu/fgSsPTTcO1PQ3W3Ut+qqglcxS1e0Gcogt8L2sSJQE2OXc63YWHcHoVhzerQBQAZgOFV9ijc2mngcEG0CgWJsqeCcThma+R4d+9ghnTUYr76NozD/Cbh10CgYEAxfBFgfEn8Ai3QR3rZ0BU76GJevCPzvARs3BTu2ylr6aUErWGm6NhDexua4kveMf9sGoJbqcpmdROu7E5HYFoulTxWoLxun5Sb/dDBLViLpYdr0cVkwe8KtUnv1djWC96d5Vi8fW8TXoNI2DlOXIUOFw+4J/8HUo/By0i6Z+R3/cCgYBEXIgHNCKtBtRIv9E3FpmcUZXYtC6iiU9re0+w3py9yD9AQCCA/XwufM9OCmQ8LUIBV04VSS0jOgQomIE8+hANzvFIMhgLWFbOABtBF30a48GqfaP+hTiFBxWTImXtb0EPLLCO3YZiS4+3E+PxZqIBRM4oiWcKbzCXQfn2lf6A+QKBgAnBY8uZtRwxcca10MpP81+0GHCEtrG+R0EOwjG0kx3rGdUqOS4miTwEcRS/uCU9xhURxNE9T+GTAyZg/62imSTF1vddCMjmPO3jw9vRjLs7Ds96s2eBtBZwVvspNW49OM34AN5n/BXddBRQTDvekXyNn1O2ztgqtpyF7viQssHVAoGBAJpIURFxAPngoIgv/s1aYYAaWlexGQ7qmN7JpuBLH9TwQQ02BT/q58U70lO6C6rU/vW3lpOtecSnjWW26P8MgsX830caflohkFqqFd9Ms37JPKxNSBts/RWqMr9m+rCLysfJN/qkY74GS0mK6OzlzHtqtCY4SVxJwuVlHX0H5Hzz"
	Client = alipay.NewClient(appid,privateKey,false)
	Bm = make(gopay.BodyMap)

)

func Pay(token string){
	// 请求参数
	Bm.Set("subject", "天秤平台发布次数")
	Bm.Set("out_trade_no", time.Now().String())
	Bm.Set("total_amount", "5.00")
	Bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	Client.SetReturnUrl("http://127.0.0.1:8082/return?Authorization="+token)         // 设置返回URL
	Client.SetNotifyUrl("http://127.0.0.1:8082/alipaynotify?Authorization="+token) // 设置异步通知URL
	// 电脑网站支付请求
	payUrl, err := Client.TradePagePay(Bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//打开默认浏览器
	payUrl = strings.Replace(payUrl,"&","^&",-1)
	exec.Command("cmd", "/c", "start",payUrl).Start()

}

func Pay2(token string){
	// 请求参数
	Bm.Set("subject", "天秤平台VIP")
	Bm.Set("out_trade_no", time.Now().String())
	Bm.Set("total_amount", "50.00")
	Bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	Client.SetReturnUrl("http://127.0.0.1:8082/return2?Authorization="+token)         // 设置返回URL
	Client.SetNotifyUrl("http://127.0.0.1:8082/alipaynotify?Authorization="+token) // 设置异步通知URL
	// 电脑网站支付请求
	payUrl, err := Client.TradePagePay(Bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//打开默认浏览器
	payUrl = strings.Replace(payUrl,"&","^&",-1)
	exec.Command("cmd", "/c", "start",payUrl).Start()

}