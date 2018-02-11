package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	// making postgres work
	_ "github.com/lib/pq"
)

// SelectPlayers gets players from db
func SelectPlayers(offset uint64, limit uint64) ([]apiobjects.Player, error) {
	var result [MaxItems]apiobjects.Player
	rows, err := db.Query("SELECT id, name, description, logo_link, rating FROM players ORDER BY id LIMIT $1 OFFSET $2;", limit, offset)
	if err != nil {
		logErr(err)
		err = rescueDb()
		if err != nil {
			logErr(err)
			return nil, errors.New("DB error")
		}

		rows, err = db.Query("SELECT id, name, description, logo_link, rating FROM players ORDER BY id LIMIT $1 OFFSET $2;", limit, offset)
		if err != nil {
			return nil, errors.New("DB error")
		}
	}

	i := 0
	for rows.Next() {
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

// SelectPlayerByID gets one player from db by id
func SelectPlayerByID(playerID uint64) (apiobjects.Player, error) {
	var result apiobjects.Player
	row := db.QueryRow("SELECT id, name, description, logo_link, rating FROM players WHERE id = $1;", playerID)

	switch err := row.Scan(&result.ID, &result.Name, &result.Description, &result.LogoLink, &result.Rating); err {
	case sql.ErrNoRows:
		log.Printf("error scanning db result: %v\n", err)
		return result, nil
	case nil:
		return result, err
	default:
		{
			log.Printf("error while getting row: %v\n", err)
			err2 := rescueDb()
			if err2 != nil {
				logErr(err2)
			}
			return result, err
		}
	}
}

// InsertPlayers inserts players into db
func InsertPlayers(players []apiobjects.Player) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		err2 := rescueDb()
		if err2 != nil {
			logErr(err2)
		}
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenPlayers := 0
	var newID uint64
	for i, player := range players {
		lenPlayers++
		row := tx.QueryRow("insert into players (name, description, logo_link, rating) values ($1, $2, $3, $4) returning ID;",
			player.Name, player.Description, player.LogoLink, player.Rating)
		switch err := row.Scan(&newID); err {
		case sql.ErrNoRows:
			log.Printf("error scanning db result: %v\n", err)
		case nil:
			result[i] = newID
		default:
			log.Printf("error while inserting row: %v\n", err)
			tx.Rollback()
			err2 := rescueDb()
			if err2 != nil {
				logErr(err2)
			}
			return nil, err
		}

	}
	tx.Commit()
	return result[0:lenPlayers], nil
}

// UpdatePlayers inserts players into db
func UpdatePlayers(players []apiobjects.Player) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		err2 := rescueDb()
		if err2 != nil {
			logErr(err2)
		}
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenPlayers := 0
	var updatedID uint64
	for i, player := range players {
		lenPlayers++
		if player.ID == nil {
			log.Printf("no id for player %d\n", lenPlayers-1)
			tx.Rollback()
			return nil, errors.New("no id for player")
		}
		row := tx.QueryRow("update players set (name, description, logo_link, rating) = ($1, $2, $3, $4) where id=$5 returning id;",
			player.Name, player.Description, player.LogoLink, player.Rating, player.ID)
		switch err := row.Scan(&updatedID); err {
		case sql.ErrNoRows:
			log.Printf("error scanning db result: %v\n", err)
		case nil:
			result[i] = updatedID
		default:
			log.Printf("error while inserting row: %v\n", err)
			tx.Rollback()
			return nil, err
		}

	}
	tx.Commit()
	return result[0:lenPlayers], nil
}
