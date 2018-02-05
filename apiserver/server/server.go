package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/apiobjects"
	"github.com/Vasilesk/go-backend-cybersport/apiserver/solver"
	"github.com/Vasilesk/go-backend-cybersport/apiserver/workers"
	"github.com/gorilla/mux"
)

// IPool is an interface for the pool of workers
type IPool interface {
	Size() int
	Run()
	AddTaskSyncTimed(f workers.Func, timeout time.Duration) (interface{}, error)
}

var wp IPool = workers.NewPool(10)

const requestWaitInQueueTimeout = time.Millisecond * 500

func methodsHandler(w http.ResponseWriter, r *http.Request) {
	// accepts json
	defer w.Write([]byte("\n"))
	start := time.Now()

	_, err := wp.AddTaskSyncTimed(func() interface{} {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		rawData := buf.Bytes()
		println("Server got:", string(rawData))

		var data apiobjects.BaseRequest

		err := json.Unmarshal(rawData, &data)
		if err != nil {
			log.Printf("error parsing json: %v", err)
			return nil
		}

		vars := mux.Vars(r)
		method := vars["method"]

		result := solver.SolveMethod(method, data)

		resp := make(map[string]interface{})

		statusP := result.GetStatus()
		if statusP != nil {
			resp["status"] = *statusP
		}
		errorP := result.GetError()
		if errorP != nil {
			resp["error"] = *errorP
		}
		dataP := result.GetData()
		if dataP != nil {
			resp["data"] = *dataP
		}

		bytesResp, err := json.Marshal(resp)
		log.Printf("response: %s", bytesResp)
		if err != nil {
			log.Printf("error creating json: %v", err)
			return nil
		}
		w.Write(bytesResp)

		return nil
	}, requestWaitInQueueTimeout)

	if err != nil {
		log.Printf("error handling request: %v ", err)
		// http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("total time for request:", elapsed)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "apiserver/html/home.html")
}

// RunHTTPServer runs backend api server
func RunHTTPServer(addr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/method/{method}", methodsHandler)
	r.HandleFunc("/", staticHandler)
	return http.ListenAndServe(addr, r)
}

func init() {
	wp.Run()
}
