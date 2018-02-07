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
	rows, err := db.Query("SELECT id, name, description, logo_link, rating FROM teams ORDER BY id LIMIT $1 OFFSET $2;", limit, offset)
	if err != nil {
		logErr(err)
		return nil, errors.New("DB error")
	}

	i := 0
	for rows.Next() {
		log.Printf("here\n")
		var p apiobjects.Team
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.LogoLink, &p.Rating)
		if err != nil {
			log.Printf("error scanning db result: %v\n", err)
		}
		result[i] = p
		i++
		// result = append(result, p)
	}

	return result[0:i], nil
}

// SelectTeamByID gets one team from db by id
func SelectTeamByID(teamID uint64) (apiobjects.Team, error) {
	var result apiobjects.Team
	row := db.QueryRow("SELECT id, name, description, logo_link, rating FROM teams WHERE id = $1;", teamID)

	switch err := row.Scan(&result.ID, &result.Name, &result.Description, &result.LogoLink, &result.Rating); err {
	case sql.ErrNoRows:
		log.Printf("error scanning db result: %v\n", err)
		return result, nil
	case nil:
		return result, err
	default:
		{
			log.Printf("error while getting row: %v\n", err)
			return result, err
		}
	}
}

// InsertTeams inserts teams into db
func InsertTeams(teams []apiobjects.Team) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		tx.Rollback()
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenTeams := 0
	var newID uint64
	for i, team := range teams {
		lenTeams++
		row := tx.QueryRow("insert into teams (name, description, logo_link, rating) values ($1, $2, $3, $4) returning ID;",
			team.Name, team.Description, team.LogoLink, team.Rating)
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
		tx.Rollback()
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenTeams := 0
	var updatedID uint64
	for i, team := range teams {
		lenTeams++
		if team.ID == nil {
			log.Printf("no id for team %d\n", lenTeams-1)
			tx.Rollback()
			return nil, errors.New("no id for team")
		}
		row := tx.QueryRow("update teams set (name, description, logo_link, rating) = ($1, $2, $3, $4) where id=$5 returning id;",
			team.Name, team.Description, team.LogoLink, team.Rating, team.ID)
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
