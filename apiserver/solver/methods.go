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
	}

	eText := "unknown method"
	return apiobjects.ErrorResponse{Error: &eText}
}

func playersGet(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	resData["name"] = "vasya"
	resData["count"] = 100500
	res := apiobjects.BaseResponse{Data: &resData}

	// offset, limit
	items, err := db.SelectPlayers(100, 50)
	if err != nil {
		log.Printf("error getting players: %v", err)
	}
	resData["items"] = items

	return res
}

func playersAdd(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	resData["name"] = "kolya"
	resData["count"] = 100501
	// resData["items"] = make([]string, 0, 10)
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Players != nil {
		items, err := db.InsertPlayers(*pData.Players)
		if err != nil {
			log.Printf("error getting players: %v", err)
		}
		resData["items"] = items
	}

	return res
}

func init() {

}
