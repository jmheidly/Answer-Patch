package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type PostgresDSN struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisDSN struct {
	Addr     string
	Password string
}

func GetPostgresDSN() string {

	dsn := &PostgresDSN{}

	getDSN("postgres", dsn)

	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dsn.User, dsn.Password, dsn.DBName, dsn.SSLMode)

}

func GetRedisDSN() *RedisDSN {
	dsn := &RedisDSN{}

	getDSN("redis", dsn)

	return dsn
}

func getDSN(dbName string, dsnStruct interface{}) {

	content, err := ioutil.ReadFile("/home/dipen/go/src/github.com/patelndipen/AP1/settings/" + os.Getenv("GO_ENV") + "/" + dbName + ".json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, dsnStruct)
	if err != nil {
		log.Fatal(err)
	}

}
