package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	// making postgres work
	_ "github.com/lib/pq"
)

// SelectTeams gets teams from db
func SelectTeams(offset uint64, limit uint64) ([]apiobjects.Team, error) {
	var result [MaxItems]apiobjects.Team
	queryStr := "SELECT id, name, description, logo_link, rating, game_id FROM teams ORDER BY id LIMIT $1 OFFSET $2;"
	rows, err := db.Query(queryStr, limit, offset)
	if err != nil {
		logErr(err)
		err = rescueDb()
		if err != nil {
			logErr(err)
			return nil, errors.New("DB error")
		}

		rows, err = db.Query(queryStr, limit, offset)
		if err != nil {
			return nil, errors.New("DB error")
		}
	}

	i := 0
	for rows.Next() {
		var p apiobjects.Team
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.LogoLink, &p.Rating, &p.GameID)
		if err != nil {
			log.Printf("error scanning db result: %v\n", err)
		}
		result[i] = p
		i++
		// result = append(result, p)
	}

	// var team *apiobjects.Team
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

// SelectTeamByID gets one team from db by id
func SelectTeamByID(teamID uint64) (apiobjects.Team, error) {
	var result apiobjects.Team
	queryStr := "SELECT id, name, description, logo_link, rating, game_id FROM teams WHERE id = $1;"
	row := db.QueryRow(queryStr, teamID)

	switch err := row.Scan(&result.ID, &result.Name, &result.Description, &result.LogoLink, &result.Rating, &result.GameID); err {
	case sql.ErrNoRows:
		log.Printf("error scanning db result: %v\n", err)
		return result, err
	case nil:
		return result, nil
	default:
		{
			log.Printf("error while getting row: %v\n", err)
			err2 := rescueDb()
			if err2 != nil {
				logErr(err2)
				return result, err
			}
			row := db.QueryRow(queryStr, teamID)
			switch err := row.Scan(&result.ID, &result.Name, &result.Description, &result.LogoLink, &result.Rating, &result.GameID); err {
			case sql.ErrNoRows:
				log.Printf("error scanning db result: %v\n", err)
				return result, err
			case nil:
				return result, nil
			default:
				return result, err
			}
		}
	}
}

// InsertTeams inserts teams into db
func InsertTeams(teams []apiobjects.Team) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		err2 := rescueDb()
		if err2 != nil {
			logErr(err2)
			return nil, err
		}
		// time.Sleep(3 * time.Second)
		tx, err2 = db.Begin()
		if err2 != nil {
			logErr(err2)
			return nil, err
		}
	}

	// stmt, err := tx.Prepare(query)
	lenTeams := 0
	var newID uint64
	for i, team := range teams {
		lenTeams++
		row := tx.QueryRow("INSERT INTO teams (name, description, logo_link, rating, game_id) VALUES ($1, $2, $3, $4, $5) RETURNING id;",
			team.Name, team.Description, team.LogoLink, team.Rating, team.GameID)
		switch err := row.Scan(&newID); err {
		case sql.ErrNoRows:
			log.Printf("error scanning db result: %v\n", err)
		case nil:
			result[i] = newID
		default:
			log.Printf("error while inserting row: %v\n", err)
			tx.Rollback()

			return nil, err
		}

	}
	tx.Commit()
	return result[0:lenTeams], nil
}

// UpdateTeams inserts teams into db
func UpdateTeams(teams []apiobjects.Team) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		err2 := rescueDb()
		if err2 != nil {
			logErr(err2)
			return nil, err
		}
		// time.Sleep(3 * time.Second)
		tx, err2 = db.Begin()
		if err2 != nil {
			logErr(err2)
			return nil, err
		}
	}

	lenTeams := 0
	var updatedID uint64
	for i, team := range teams {
		lenTeams++
		if team.ID == nil {
			log.Printf("no id for team %d\n", lenTeams-1)
			tx.Rollback()
			return nil, errors.New("no id for team")
		}
		row := tx.QueryRow("UPDATE teams SET (name, description, logo_link, rating, game_id) = ($1, $2, $3, $4, $5) WHERE id=$6 RETURNING id;",
			team.Name, team.Description, team.LogoLink, team.Rating, team.GameID, team.ID)
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
	return result[0:lenTeams], nil
}
