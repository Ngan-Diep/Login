package routes
import (
	"github.com/labstack/echo/v4"
    "github.com/Login/gologin/handler"
    "github.com/Login/gologin/middleware"

)
func InitRoutes(){
	e.POST("/",handler.HomeHandler)
	e.POST("/login",handler.Login)
	r:=e.Group("/restricted")
	r.Use(middleware.JWTAuth())
	r.GET("",handler.Restricted)
}