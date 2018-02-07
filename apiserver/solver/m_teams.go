package solver

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	"github.com/Vasilesk/go-backend-cybersport/db"
)

func teamsGetByID(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.ID == nil {
		errorDesc := errNoID
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	teamID := *pData.ID
	var err error
	resData["team"], err = db.SelectTeamByID(teamID)
	if err != nil {
		log.Printf("error getting team by id: %v", err)
		errorDesc := "error getting team by id"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}
	return res
}

func teamsGet(pData *apiobjects.BaseRequest) apiobjects.IResponse {
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
	items, err := db.SelectTeams(offset, limit)
	if err != nil {
		log.Printf("error getting teams: %v", err)
	}
	resData["items"] = items

	return res
}

func teamsAdd(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	// resData["count"] = 100501
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Teams == nil {
		errorDesc := "no team sent"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if len(*pData.Teams) > db.MaxItems {
		errorDesc := "too many teams to add"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Teams != nil {
		// start := time.Now()
		items, err := db.InsertTeams(*pData.Teams)
		// t := time.Now()
		// elapsed := t.Sub(start)
		// log.Println("time for db:", elapsed)
		if err != nil {
			log.Printf("error adding teams: %v", err)
		}
		resData["items"] = items
	}

	return res
}

func teamsUpdate(pData *apiobjects.BaseRequest) apiobjects.IResponse {
	resData := make(map[string]interface{})
	res := apiobjects.BaseResponse{Data: &resData}

	if pData.Teams == nil {
		errorDesc := "no team sent"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if len(*pData.Teams) > db.MaxItems {
		errorDesc := "too many teams to update"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	if pData.Teams != nil {
		items, err := db.UpdateTeams(*pData.Teams)
		if err != nil {
			log.Printf("error updating teams: %v", err)
		}
		resData["updated_ids"] = items
	}

	return res
}
