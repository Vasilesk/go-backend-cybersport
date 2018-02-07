package solver

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	"github.com/Vasilesk/go-backend-cybersport/db"
)

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
	}

	eText := "unknown method"
	return apiobjects.ErrorResponse{Error: &eText}
}

func playersGetByID(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.ID == nil {
		errorDesc := "id is not set"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	playerID := *pData.ID
	var err error
	resData["player"], err = db.SelectPlayerByID(playerID)
	if err != nil {
		log.Printf("error getting player by id: %v", err)
		errorDesc := "error getting player by id"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}
	return res
}

func playersGet(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	// resData["count"] = 100500
	res := apiobjects.BaseResponse{Data: &resData}

	var offset uint64
	var limit uint64
	if pData.Offset == nil {
		offset = 0
	} else {
		offset = *pData.Offset
	}
	if pData.Limit == nil {
		limit = db.MaxItems
	} else {
		limit = *pData.Limit
	}
	items, err := db.SelectPlayers(offset, limit)
	if err != nil {
		log.Printf("error getting players: %v", err)
	}
	resData["items"] = items

	return res
}

func playersAdd(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	// resData["count"] = 100501
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Players == nil {
		errorDesc := "no player sent"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if len(*pData.Players) > db.MaxItems {
		errorDesc := "too many players to add"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Players != nil {
		// start := time.Now()
		items, err := db.InsertPlayers(*pData.Players)
		// t := time.Now()
		// elapsed := t.Sub(start)
		// log.Println("time for db:", elapsed)
		if err != nil {
			log.Printf("error adding players: %v", err)
		}
		resData["items"] = items
	}

	return res
}

func playersUpdate(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	// resData["count"] = 100502
	res := apiobjects.BaseResponse{Data: &resData}

	if len(*pData.Players) > db.MaxItems {
		errorDesc := "too many players to update"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Players != nil {
		items, err := db.UpdatePlayers(*pData.Players)
		if err != nil {
			log.Printf("error updating players: %v", err)
		}
		resData["updated_ids"] = items
	}

	return res
}

func init() {

}
