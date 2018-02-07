package solver

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	"github.com/Vasilesk/go-backend-cybersport/db"
)

func tournamentsGetByID(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.ID == nil {
		errorDesc := errNoID
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	tournamentID := *pData.ID
	var err error
	resData["tournament"], err = db.SelectTournamentByID(tournamentID)
	if err != nil {
		log.Printf("error getting tournament by id: %v", err)
		errorDesc := "error getting tournament by id"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}
	return res
}

func tournamentsGet(pData *apiobjects.BaseRequest) apiobjects.IResponse {
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
	items, err := db.SelectTournaments(offset, limit)
	if err != nil {
		log.Printf("error getting tournaments: %v", err)
	}
	resData["items"] = items

	return res
}

func tournamentsAdd(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	// resData["count"] = 100501
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Tournaments == nil {
		errorDesc := "no tournament sent"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if len(*pData.Tournaments) > db.MaxItems {
		errorDesc := "too many tournaments to add"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Tournaments != nil {
		// start := time.Now()
		items, err := db.InsertTournaments(*pData.Tournaments)
		// t := time.Now()
		// elapsed := t.Sub(start)
		// log.Println("time for db:", elapsed)
		if err != nil {
			log.Printf("error adding tournaments: %v", err)
		}
		resData["items"] = items
	}

	return res
}

func tournamentsUpdate(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Tournaments == nil {
		errorDesc := "no tournament sent"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if len(*pData.Tournaments) > db.MaxItems {
		errorDesc := "too many tournaments to update"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Tournaments != nil {
		items, err := db.UpdateTournaments(*pData.Tournaments)
		if err != nil {
			log.Printf("error updating tournaments: %v", err)
		}
		resData["updated_ids"] = items
	}

	return res
}
