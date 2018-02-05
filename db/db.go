package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	// don`t understand why we need this import
	// but nothing works without it
	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	_ "github.com/lib/pq"
)

const maxItems = 100

var db *sql.DB

type conf struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
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

// GetPlayers gets players from db
func GetPlayers() ([]apiobjects.Player, error) {
	var result [maxItems]apiobjects.Player
	// db := getDbConn()
	rows, err := db.Query("SELECT id, name, description, logo_link, rating FROM players;")
	if err != nil {
		logErr(err)
		return nil, errors.New("DB error")
	}

	// got := []row{}
	i := 0
	for rows.Next() {
		log.Printf("here\n")
		var p apiobjects.Player
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.LogoLink, &p.Rating)
		if err != nil {
			log.Printf("error scanning db result: %v\n", err)
		}
		result[i] = p
		i++
		// result = append(result, p)
	}

	// var player *apiobjects.Player
	// for rows.Next() {
	// 	err = rows.Scan(&username, &message)
	// 	if err != nil {
	// 		logErr(err)
	// 		return result, errors.New("Backend error")
	// 	}
	// 	result[i] = username + ": " + message
	// 	i++
	// }

	return result[0:i], nil
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
func GetHistory() ([maxItems]string, error) {
	var result [maxItems]string
	// db := getDbConn()
	rows, err := db.Query("SELECT login, message FROM history inner join users on users.id = user_id ORDER BY history.id LIMIT $1;",
		maxItems)
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

// func initDb() error {
// 	c := getConf()
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=5432 host=vasilesk.ru",
// 		c.User, c.Password, c.Database)
// 	var err error
// 	db, err = sql.Open("postgres", dbinfo)
//
// 	return err
// }

func getDbConn() (*sql.DB, error) {
	// c := getConf()
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=5432 host=vasilesk.ru",
	// 	c.User, c.Password, c.Database)
	dbinfo := fmt.Sprintf("sslmode=disable user=%s port=%s host=%s",
		"postgres", "5432", "localhost")
	conn, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, errors.New("db was not connected")
	}

	return conn, nil
}

func init() {
	var err error
	db, err = getDbConn()
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
