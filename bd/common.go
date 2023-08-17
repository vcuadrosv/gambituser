package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vcuadrosv/gambituser/models"
	"github.com/vcuadrosv/gambituser/secretm"
)

var SecretModel models.SecretRDSJSON
var err error
var Db *sql.DB //Todo lo que es conexion de base de datos debe ser de tipo puntero por velocidad

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// Aqui vamos a conectarnos a la base de datos y la conexion va a quedar abierta
// el drive de mysql no es estandar de mysql toca importarlo go get github.com/go-sql-driver/mysql
// conexion basica de la base de datos
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//un ping para ver que la conexion esta abierta
	err := Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion existosa de la base de datos")
	return nil
}

// Esta funcion establece los parametros que se ejecutaran como permisos dentro de la base de datos
// y AWS
func ConnStr(claves models.SecretRDSJSON) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPassword=true",
		dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
