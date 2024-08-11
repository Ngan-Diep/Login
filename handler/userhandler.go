package handler
import(
	"net/http"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/echo/v4/jwt"

)

type LoginRequest struct{
	Username string  `json:username` `validate:username`
	Password string  `json:password` `validate:password`
}

func Login(c echo.Context) error {
 var LogReq LoginRequest
 err := c.Bind(&LogReq), if err!=nil{
	return c.JSON(http.StatusUnauthorized,echo.Map{message:"request error"})
 }

 err := c.Validate(&LogReq),if err!=nil{
	return c.JSON(http.StatusUnauthorized,echo.Map(message:"validate error"))
 }
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": userID,
	"exp":     time.Now().Add(time.Hour * 72).Unix(),
})

tokenString, err := token.SignedString([]byte("your_secret_key"))
if err != nil {
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
}

return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}

}
func HomeHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", nil)
}