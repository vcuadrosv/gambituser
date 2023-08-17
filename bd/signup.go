package bd

//Este codigo graba dentro de la base de datos
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vcuadrosv/gambituser/models"
	"github.com/vcuadrosv/gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUDI, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)
	//Aqui envia un error en formato execel
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp Existosa")
	return err
}
