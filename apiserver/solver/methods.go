package solver

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
)

// SolveMethod processes data that server gets
func SolveMethod(method string, data apiobjects.BaseRequest) apiobjects.IResponse {
	if data.V == nil {
		log.Printf("error getting `v`: nil val\n")
		errorDesc := "no api version"
		return apiobjects.ErrorResponse{Error: &errorDesc}
	}

	resData := make(map[string]interface{})
	resData["name"] = "vasya"
	return apiobjects.BaseResponse{Data: &resData}
}

func init() {

}
