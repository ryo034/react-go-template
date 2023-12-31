package datasource

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	md "github.com/go-sql-driver/mysql"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/datetime"
)

type DataSource interface {
	Open(ops ...Option) *sql.DB
}

type mysql struct {
	host       string
	port       uint
	dbname     string
	usr        string
	pass       string
	dbRootCert string
	dbCert     string
	dbKey      string
}

func NewMysqlDataSource(host string, port uint, dbname string, usr string, pass string, dbRootCert string, dbCert string, dbKey string) DataSource {
	return &mysql{host, port, dbname, usr, pass, dbRootCert, dbCert, dbKey}
}

func (m *mysql) dsn() string {
	tls, err := m.tls()
	if err != nil {
		log.Println("database tls settings error")
		log.Println(err)
		return ""
	}
	tz := datetime.CountryTz["Tokyo"]
	loc, _ := time.LoadLocation(tz)
	c := md.Config{
		DBName:               m.dbname,
		User:                 m.usr,
		Passwd:               m.pass,
		Addr:                 fmt.Sprintf("%s:%d", m.host, m.port),
		Net:                  "tcp",
		TLSConfig:            tls,
		Loc:                  loc,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return c.FormatDSN()
}

// https://github.com/GoogleCloudPlatform/golang-samples/blob/b4e5aa87903747985c8c096f29598d0de5b6cbbc/cloudsql/mysql/database-sql/cloudsql.go
func (m *mysql) tls() (string, error) {
	dbURI := ""
	// [START_EXCLUDE]
	// [START cloud_sql_mysql_databasesql_sslcerts]
	// (OPTIONAL) Configure SSL certificates
	// For deployments that connect directly to a Cloud SQL instance without
	// using the Cloud SQL Proxy, configuring SSL certificates will ensure the
	// connection is encrypted. This step is entirely OPTIONAL.
	if m.dbRootCert != "" {
		pool := x509.NewCertPool()
		pem, err := os.ReadFile(m.dbRootCert)
		if err != nil {
			return "", err
		}
		if ok := pool.AppendCertsFromPEM(pem); !ok {
			return "", err
		}
		cert, err := tls.LoadX509KeyPair(m.dbCert, m.dbKey)
		if err != nil {
			return "", err
		}
		md.RegisterTLSConfig("cloudsql", &tls.Config{
			RootCAs:               pool,
			Certificates:          []tls.Certificate{cert},
			InsecureSkipVerify:    true,
			VerifyPeerCertificate: verifyPeerCertFunc(pool),
		})
		dbURI += "cloudsql"
	}
	return dbURI, nil
}

func (m *mysql) Open(ops ...Option) *sql.DB {
	db, err := sql.Open("mysql", m.dsn())
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

type Option interface {
	apply(db *sql.DB)
}

func MaxOpenConns(val int) Option {
	return maxOpenConns{val}
}

type maxOpenConns struct {
	val int
}

func (o maxOpenConns) apply(db *sql.DB) {
	db.SetMaxOpenConns(o.val)
}

type maxIdleConns struct {
	val int
}

func (o maxIdleConns) apply(db *sql.DB) {
	db.SetMaxIdleConns(o.val)
}

func ConnMaxLifetime(val time.Duration) Option {
	return connMaxLifetime{val}
}

type connMaxLifetime struct {
	val time.Duration
}

func (o connMaxLifetime) apply(db *sql.DB) {
	db.SetConnMaxLifetime(o.val)
}

// verifyPeerCertFunc returns a function that verifies the peer certificate is
// in the cert pool.
func verifyPeerCertFunc(pool *x509.CertPool) func([][]byte, [][]*x509.Certificate) error {
	return func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
		if len(rawCerts) == 0 {
			return fmt.Errorf("no certificates available to verify")
		}

		cert, err := x509.ParseCertificate(rawCerts[0])
		if err != nil {
			return err
		}

		opts := x509.VerifyOptions{Roots: pool}
		if _, err = cert.Verify(opts); err != nil {
			return err
		}
		return nil
	}
}
