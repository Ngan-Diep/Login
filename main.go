package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "github.com/mattn/go-colorable"
    "github.com/mattn/go-isatty"
    "github.com/valyala/bytebufferpool"
    "github.com/valyala/fasttemplate"
    "golang.org/x/text"
    "golang.org/x/crypto"
    "golang.org/x/net"
    "golang.org/x/sys"
   "github.com/Login/gologin/database"
    "github.com/Login/gologin/handler"
    "github.com/Login/gologin/config"
    "github.com/Login/gologin/middleware"
    "github.com/Login/gologin/models"
    "github.com/Login/gologin/routes"
    "github.com/Login/gologin/static"
    "github.com/Login/gologin/templates"
    "github.com/Login/gologin/utils"
)


func main() {
	e := echo.New()
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	if err := database.RunMigrations(); err != nil {
		log.Fatal(err)
	}
	routes.InitRoutes()

	e.Static("/Static", "static")
	e.Renderer := &render.HTML{
		Templates: map[string]*template.Template{
			index.html: template.Must(template.ParseFiles("templates/index.html")),
		},
	}

	e.Logger.Fatal(e.Start(":8080"))

}
