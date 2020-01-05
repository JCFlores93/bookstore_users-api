package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

//MYSQL_USERS_USERNAME=root;
//MYSQL_USERS_PASSWORD=password;
//MYSQL_USERS_HOST=172.17.0.2:3306;
//MYSQL_USERS_SCHEMA=users_db

const (
	mysql_users_username = "MYSQL_USERS_USERNAME"
	mysql_users_password = "MYSQL_USERS_PASSWORD"
	mysql_users_host = "MYSQL_USERS_HOST"
	mysql_users_schema = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	if Client != nil {
		fmt.Println("Client not null")
	}
	//mysql.SetLogger(logger.GetLo)
 	log.Println("database successfully configured")
}
