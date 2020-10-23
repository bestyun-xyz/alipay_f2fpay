package main

import (
	"f2fpay/controllers"
	"f2fpay/system"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {

	system.InitPara()

	e := echo.New()


	e.GET("/", controllers.Index)

	e.GET("/barpay_test", controllers.Barpay)
	e.POST("/barpay_test", controllers.Barpay)
	e.GET("/qrpay_test", controllers.Qrpay)
	e.POST("/qrpay_test", controllers.Qrpay)
	e.GET("/query_test", controllers.Query)
	e.POST("/query_test", controllers.Query)
	e.GET("/refund_test", controllers.Refund)
	e.POST("/refund_test", controllers.Refund)


	e.Static("static", "web/static")
	//e.File("/favicon.ico", "web/static/img/favicon.ico")
	//初始化模版引擎
	t := &Template{
		//模版引擎支持提前编译模版, 这里对views目录下以html结尾的模版文件进行预编译处理
		//预编译处理的目的是为了优化后期渲染模版文件的速度
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}

	//向echo实例注册模版引擎
	e.Renderer = t

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(system.ListenPort))

}
