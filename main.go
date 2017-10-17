package main

import (
	"html/template"
	"net/http"

	"shortest/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var server *gin.Engine
var templates map[string]*template.Template

func init() {
	loadTemplates()
}

func main() {
	db := gormConnect()
	server = gin.Default()

	server.GET("/", indexRouter)
	server.GET("/signup", signupRoute)

	server.POST("/signin", func(g *gin.Context) { controller.Signin(db, g) })
	server.POST("/signup/regist", func(g *gin.Context) { controller.Signup(db, g) })

	server.GET("/home", homeRoute)

	server.Static("/assets", "./assets")
	server.Run(":8080")
}

func indexRouter(g *gin.Context) {
	server.SetHTMLTemplate(templates["index"])
	g.HTML(http.StatusOK, "_base_index.html", nil)
}

func signupRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["signup"])
	g.HTML(http.StatusOK, "_base.html", nil)
}

func homeRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["home"])
	g.HTML(http.StatusOK, "_base.html", nil)
}

/*
　テンプレート設定
*/
func loadTemplates() {
	var baseTemplate = "templates/common/_base.html"
	var indexTemplate = "templates/common/_base_index.html"
	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(template.ParseFiles(indexTemplate, "templates/back/index.html"))
	templates["signup"] = template.Must(template.ParseFiles(baseTemplate, "templates/back/signup.html"))
	templates["home"] = template.Must(template.ParseFiles(baseTemplate, "templates/back/home.html"))
}

/*
　GORM設定
*/
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "mariadb"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "shortest_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
