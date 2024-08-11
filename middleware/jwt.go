package middleware
import(
	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func JWTAuth(e echo.Echo){
	return middleware.JWTAuth(middleware.JWTauth{
		SigninKey:[]byte("secret")
	})

}