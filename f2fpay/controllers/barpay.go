package controllers

import (
	"f2fpay/system"
	"fmt"
	"github.com/labstack/echo"
	"github.com/smartwalle/alipay/v3"
	"net/http"
)

func Barpay(c echo.Context) error {
	if c.Request().Method == "GET" {
		//GET

		return c.Render(http.StatusOK, "barpay_test.html", "")
	} else {
		//POST

		out_trade_no := c.FormValue("out_trade_no")
		subject := c.FormValue("subject")
		total_amount := c.FormValue("total_amount")
		auth_code := c.FormValue("auth_code")

		var p = alipay.TradePay{}
		p.OutTradeNo = out_trade_no
		p.Subject = subject
		p.TotalAmount = total_amount
		p.Scene = "bar_code"
		p.AuthCode = auth_code

		rsp, err := system.Client.TradePay(p)
		if err != nil {
			fmt.Print(err,"\n")
		}
		if rsp.Content.Code != alipay.CodeSuccess {
			fmt.Print(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg,"\n")
		}
		//fmt.Printf("%v", rsp.Content)

		return c.JSON(http.StatusOK, rsp.Content)
	}
}
