package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	// making postgres work
	_ "github.com/lib/pq"
)

// MaxItems is a const restricting count of items to get, add, update, delete
const MaxItems = 1000

var db *sql.DB

type conf struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	RHost    string `yaml:"r_host"`
	RPort    string `yaml:"r_port"`
}

// old
func logErr(err error) {
	log.Printf("Error %v ", err)
}

func getConf() conf {

	yamlFile, err := ioutil.ReadFile("config/db.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var c conf
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// old below

// CreateKey creates key for user by his id
func CreateKey(userID int) string {
	// db := getDbConn()
	db.QueryRow("INSERT INTO access_keys(user_id,key) VALUES($1,$2);", userID, "s")

	return "s"
}

// GetUserByKey returns user login by key or "" if no user with the key exists
func GetUserByKey(key string) (int, string) {
	// db := getDbConn()
	rows, err := db.Query("SELECT user_id, login FROM access_keys JOIN users ON user_id=id WHERE key=$1 LIMIT 1;", key)
	if err != nil {
		logErr(err)
		return 0, ""
	}

	var userID int
	var username string
	for rows.Next() {
		err = rows.Scan(&userID, &username)
		if err != nil {
			logErr(err)
			return 0, ""
		}
		return userID, username
	}
	return 0, ""
}

// GetHistory returns message history
func GetHistory() ([MaxItems]string, error) {
	var result [MaxItems]string
	// db := getDbConn()
	rows, err := db.Query("SELECT login, message FROM history inner join users on users.id = user_id ORDER BY history.id LIMIT $1;",
		MaxItems)
	if err != nil {
		logErr(err)
		return result, errors.New("DB error")
	}

	i := 0
	var username string
	var message string
	for rows.Next() {
		err = rows.Scan(&username, &message)
		if err != nil {
			logErr(err)
			return result, errors.New("Backend error")
		}
		result[i] = username + ": " + message
		i++
	}

	return result, nil
}

// AddMessage stores message into DB
func AddMessage(userID int, message []byte) error {
	// db := getDbConn()
	db.QueryRow("INSERT INTO history(user_id, message) VALUES($1, $2);", userID, message)

	return nil
}

func getDbConn(c conf) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s host=%s",
		c.User, c.Password, c.Database, c.Port, c.Host)
	// dbinfo := fmt.Sprintf("sslmode=disable user=%s port=%s host=%s",
	// 	"postgres", "5432", "localhost")
	conn, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, errors.New("db was not connected")
	}

	return conn, nil
}

func init() {
	var err error
	c := getConf()
	db, err = getDbConn(c)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM players LIMIT 1;")
	if err != nil {
		log.Printf("oops: %v\n", err)
	} else {
		log.Printf("result: %v\n", rows)
		for rows.Next() {
			log.Printf("result row: %v\n", rows)
		}
	}
}
