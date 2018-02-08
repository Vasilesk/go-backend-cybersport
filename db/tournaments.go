package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	// making postgres work
	_ "github.com/lib/pq"
)

// SelectTournaments gets tournaments from db
func SelectTournaments(offset uint64, limit uint64) ([]apiobjects.Tournament, error) {
	var result [MaxItems]apiobjects.Tournament
	rows, err := db.Query("SELECT id, name, description, logo_link, is_active, game_id FROM tournaments ORDER BY id LIMIT $1 OFFSET $2;", limit, offset)
	if err != nil {
		logErr(err)
		return nil, errors.New("DB error")
	}

	i := 0
	for rows.Next() {
		var p apiobjects.Tournament
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.LogoLink, &p.IsActive, &p.GameID)
		if err != nil {
			log.Printf("error scanning db result: %v\n", err)
		}
		result[i] = p
		i++
	}

	return result[0:i], nil
}

// SelectTournamentByID gets one tournament from db by id
func SelectTournamentByID(tournamentID uint64) (apiobjects.Tournament, error) {
	var result apiobjects.Tournament
	row := db.QueryRow("SELECT id, name, description, logo_link, is_active, game_id FROM tournaments WHERE id = $1;", tournamentID)

	switch err := row.Scan(&result.ID, &result.Name, &result.Description, &result.LogoLink, &result.IsActive, &result.GameID); err {
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

// InsertTournaments inserts tournaments into db
func InsertTournaments(tournaments []apiobjects.Tournament) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenTournaments := 0
	var newID uint64
	for i, tournament := range tournaments {
		lenTournaments++
		row := tx.QueryRow("insert into tournaments (name, description, logo_link, is_active, game_id) values ($1, $2, $3, $4, $5) returning ID;",
			tournament.Name, tournament.Description, tournament.LogoLink, tournament.IsActive, tournament.GameID)
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
	return result[0:lenTournaments], nil
}

// UpdateTournaments inserts tournaments into db
func UpdateTournaments(tournaments []apiobjects.Tournament) ([]uint64, error) {
	var result [MaxItems]uint64
	tx, err := db.Begin()
	if err != nil {
		log.Printf("error starting tx: %v\n", err)
		return nil, err
	}

	// stmt, err := tx.Prepare(query)
	lenTournaments := 0
	var updatedID uint64
	for i, tournament := range tournaments {
		lenTournaments++
		if tournament.ID == nil {
			log.Printf("no id for tournament %d\n", lenTournaments-1)
			tx.Rollback()
			return nil, errors.New("no id for tournament")
		}
		row := tx.QueryRow("update tournaments set (name, description, logo_link, is_active, game_id) = ($1, $2, $3, $4, $5) where id=$6 returning id;",
			tournament.Name, tournament.Description, tournament.LogoLink, tournament.IsActive, tournament.GameID, tournament.ID)
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
	return result[0:lenTournaments], nil
}
