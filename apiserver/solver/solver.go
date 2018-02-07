package solver

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
)

const errNoID = "id is not provided"

// SolveMethod processes data that server gets
func SolveMethod(method string, data apiobjects.BaseRequest) apiobjects.IResponse {
	if data.V == nil {
		log.Printf("error getting `v`: nil val\n")
		errorDesc := "no api version"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	switch method {
	case "players.get":
		return playersGet(&data)
	case "players.add":
		return playersAdd(&data)
	case "players.update":
		return playersUpdate(&data)
	case "players.getById":
		return playersGetByID(&data)
	case "teams.get":
		return teamsGet(&data)
	case "teams.add":
		return teamsAdd(&data)
	case "teams.update":
		return teamsUpdate(&data)
	case "teams.getById":
		return teamsGetByID(&data)
	case "tournaments.get":
		return tournamentsGet(&data)
	case "tournaments.add":
		return tournamentsAdd(&data)
	case "tournaments.update":
		return tournamentsUpdate(&data)
	case "tournaments.getById":
		return tournamentsGetByID(&data)
	}

	eText := "unknown method"
	return apiobjects.ErrorResponse{Error: &eText}
}

func init() {

}
