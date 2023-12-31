package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

type postgresql struct {
	host   string
	port   uint
	dbname string
	usr    string
	pass   string
}

func NewPostgresqlDataSource(host string, port uint, dbname string, usr string, pass string) DataSource {
	return &postgresql{host, port, dbname, usr, pass}
}

type Config struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

func (m *postgresql) dsn() string {
	config := Config{
		Host:     m.host,
		Port:     m.port,
		User:     m.usr,
		Password: m.pass,
		DbName:   m.dbname,
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)
}

func (m *postgresql) Open(ops ...Option) *sql.DB {
	db, err := sql.Open("pgx", m.dsn())
	if err != nil {
		log.Printf("datasource Open Error: %v\n", err)
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	for _, op := range ops {
		op.apply(db)
	}
	return db
}
