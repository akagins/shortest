package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type account struct {
	Accountid string `form:"accountid" json:"accountid" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

/*
　サインイン処理
*/
func Signin(db *gorm.DB, g *gin.Context) {

	var url = ""

	g.Request.ParseForm()

	accountEx := account{}
	db.Where("(accountid=? or email=?) and password=?", g.PostForm("accountid"), g.PostForm("accountid"), g.PostForm("password")).Find(&accountEx)
	if accountEx.Accountid != "" {
		url = "/home"
	} else {
		url = "/"
	}
	g.JSON(200, gin.H{"url": url})
}

/*
　サインアップ処理
*/
func Signup(db *gorm.DB, g *gin.Context) {
	g.Request.ParseForm()

	accountEx := account{}
	accountEx.Accountid = g.PostForm("accountid")
	accountEx.Email = g.PostForm("email")
	accountEx.Password = g.PostForm("password")

	if err := g.Bind(&accountEx); err != nil {
		log.Println(err)
		g.JSON(http.StatusOK, gin.H{"url": "", "message": "ユーザ名かパスワードが間違っています"})
		return
	}

	db.Create(&accountEx)
	g.JSON(http.StatusOK, gin.H{"url": "/"})
}
