package database
import(
	"log"
	"gologin/config"
	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"

)

var DB *sqlx.DB



func InitDB(dbConfig config.DbConfig) error{
	var err error
	dsn := fmt.Println("%s:%s@tcp(%s:%d)/%s",&dbconfig.User,&dbconfig.Password,&dbconfig.Host,&dbconfig.Port,&dbcongig.Name)
	Db,err := sqlx.Open("ngandiep",dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

}

func GetUserByName(username string)(models.User string){
	var user models.User
	err:=DB.QueryRow(SELECT *username,password FROM users WHERE username =?,username).Scan(&user.Username,&user.Password)
	return user,err

}


// RunMigrations runs the necessary SQL migrations
func RunMigrations() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Database migrations executed successfully")
	return nil
}