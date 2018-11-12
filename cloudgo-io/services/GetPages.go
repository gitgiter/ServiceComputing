package services

import (
	"fmt"

	"github.com/kataras/iris"
)

// LoadResources register views and load static resources
func LoadResources(app *iris.Application) {
	// register views from ./templates folder
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	// favicon
	app.Favicon("./static/img/favicon.ico")

	// load static files
	// first parameter is the href request url, second is the real system path
	app.StaticWeb("/public", "./static")
}

// GetLoginPage load html and static web
func GetLoginPage(app *iris.Application) {

	app.Get("/login", func(ctx iris.Context) {
		ctx.View("login.html")
	})
}

// User contains the form data
// all field need to be exported (upper case)
type User struct {
	Username string
	Password string
}

// GetInfoPage load html and static web
func GetInfoPage(app *iris.Application) {

	app.Post("/info", func(ctx iris.Context) {
		// get the form data
		form := User{}
		err := ctx.ReadForm(&form)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}

		fmt.Println(form)

		// bind data by passing key-value pair
		username := form.Username
		password := form.Password
		ctx.ViewData("username", username)
		ctx.ViewData("password", password)
		ctx.View("info.html")
	})
}

// NotImplement returns 501 error to client
func NotImplement(app *iris.Application) {

	app.Get("/unknown", func(ctx iris.Context) {
		ctx.StatusCode(501)
		ctx.JSON(iris.Map{
			"error": "501 not implement error",
		})
	})
}

// GetStaticPage show all the static files
func GetStaticPage(app *iris.Application) {

	app.Get("/public", func(ctx iris.Context) {
		ctx.HTML(`<a href='/public/css/main.css'>/public/css/main.css</a><br/><br/>
			<a href='/public/img/bg.jpg'>/public/img/bg.jpg</a><br/><br/>
			<a href='/public/img/favicon.ico'>/public/img/favicon.ico</a><br/><br/>
			<a href='/public/js/showStatic.js'>/public/js/showStatic.js</a>`)
	})
}
