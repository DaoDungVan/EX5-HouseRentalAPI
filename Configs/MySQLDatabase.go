package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MySQLDB *sql.DB

// đọc biến môi trường, chưa set thì lấy giá trị mặc định
func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func ConnectMySQL() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getenv("MYSQL_USER", "root"),
		getenv("MYSQL_PASSWORD", "root"),
		getenv("MYSQL_HOST", "127.0.0.1"),
		getenv("MYSQL_PORT", "3306"),
		getenv("MYSQL_DB", "house_rental_db"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("cannot open MySQL: %w", err)
	}

	// sql.Open chưa kết nối thật nên phải Ping thử
	if err := db.Ping(); err != nil {
		return fmt.Errorf("cannot connect to MySQL: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	MySQLDB = db
	fmt.Println("MySQL connected successfully.")
	return nil
}
