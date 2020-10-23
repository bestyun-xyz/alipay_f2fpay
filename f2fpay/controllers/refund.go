package controllers

import (
	"f2fpay/system"
	"fmt"
	"github.com/labstack/echo"
	"github.com/smartwalle/alipay/v3"
	"net/http"
)

func Refund(c echo.Context) error {
	if c.Request().Method == "GET" {
		//GET

		return c.Render(http.StatusOK, "refund_test.html", "")
	} else {
		//POST
		out_trade_no := c.FormValue("out_trade_no")
		refund_amount := c.FormValue("refund_amount")
		out_request_no := c.FormValue("out_request_no")

		var p = alipay.TradeRefund{}
		p.RefundAmount = refund_amount
		p.OutTradeNo = out_trade_no
		p.OutRequestNo = out_request_no
		rsp, err := system.Client.TradeRefund(p)
		if err != nil {
			fmt.Print(err,"\n")
		}
		//fmt.Printf("%v", rsp.Content)

		return c.JSON(http.StatusOK, rsp.Content)
	}
}

