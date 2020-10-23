package controllers

import (
	"f2fpay/system"
	"fmt"
	"github.com/labstack/echo"
	"github.com/smartwalle/alipay/v3"
	"net/http"
)

func Query(c echo.Context) error {
	if c.Request().Method == "GET" {
		//GET

		return c.Render(http.StatusOK, "query_test.html", "")
	} else {
		//POST
		out_trade_no := c.FormValue("out_trade_no")

		var p = alipay.TradeQuery{}
		p.OutTradeNo = out_trade_no
		//p.QueryOptions = []string{"TRADE_SETTLE_INFO"}

		rsp, err := system.Client.TradeQuery(p)
		if err != nil {
			fmt.Print(err,"\n")
		}
		if rsp.Content.Code != alipay.CodeSuccess {
			fmt.Print(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg,"\n")
		}
		//fmt.Printf("%+v",rsp.Content)
		return c.JSON(http.StatusOK, rsp.Content)
	}
}
