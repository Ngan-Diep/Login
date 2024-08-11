package utils
import(
	"time"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
)


func HashPassword(password string) string{
	bytes,err:=bcrypt.GenerateFromPassword([]byte (passwword),14)
	return string(bytes),err
}

func CheckPassword(password,hash string) bool {
	err:=bcrypt.ComparePasswordAndHash([]byte (password),[]byte (hash))
	return err == nil
}

func GenerateJWT(userID int, secretKey string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    return token.SignedString([]byte(secretKey))
}
}