package crud_mysql

import (
	"ProgrammingGo/log"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateT(table string) {
	// Configure the database connection (always check errors)
	db := InitDB()

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err := db.Ping(); err != nil {
		log.Debug.Println("something went wrong")
	}
	query := `
    CREATE TABLE t_users (
       id INT(11) AUTO_INCREMENT,
       username VARCHAR(11) NOT NULL,
       password VARCHAR(11) NOT NULL,
       PRIMARY KEY (id)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if ret, err := db.Exec(query); err != nil {
		log.Debug.Println(err)
	} else {
		log.Info.Println(ret)
	}

}
