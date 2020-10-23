package controllers

import (
	"bytes"
	"f2fpay/system"
	"fmt"
	"github.com/labstack/echo"
	"github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay/v3"
	"net/http"
)

func Qrpay(c echo.Context) error {
	if c.Request().Method == "GET" {
		//GET
		return c.Render(http.StatusOK, "qrpay_test.html", "")
	} else {
		//POST

		out_trade_no := c.FormValue("out_trade_no")
		subject := c.FormValue("subject")
		total_amount := c.FormValue("total_amount")

		var p = alipay.TradePreCreate{}
		p.OutTradeNo = out_trade_no
		p.Subject = subject
		p.TotalAmount = total_amount
		p.TimeoutExpress="5m"//超时5分钟
		rsp, err := system.Client.TradePreCreate(p)
		if err != nil {
			fmt.Print(err,"\n")
		}
		if rsp.Content.Code != alipay.CodeSuccess {
			fmt.Print(rsp.Content.Msg, rsp.Content.SubMsg,"\n")
		}

		qrbytes, err := qrcode.Encode(rsp.Content.QRCode, qrcode.Medium,256)
		return c.Stream(200,"image/png",bytes.NewReader(qrbytes))
	}
}