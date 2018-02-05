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

	if method == "players.get" {
		return playersGet(&data)
	}

	eText := "unknown method"
	return apiobjects.ErrorResponse{Error: &eText}
}

func playersGet(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	resData["name"] = "vasya"
	resData["count"] = 100500
	resData["items"] = make([]string, 0, 10)
	res := apiobjects.BaseResponse{Data: &resData}

	items, err := db.GetPlayers()
	if err != nil {
		log.Printf("error getting players: %v", err)
	}
	resData["items"] = items

	return res
}

func init() {

}
