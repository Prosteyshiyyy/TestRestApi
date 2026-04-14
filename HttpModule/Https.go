package HttpModule

import (
	"RestApi/DTOstructs"
	"RestApi/IDstructs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

var AllTargets = IDstructs.ListTargets{}

//var IdTarget IDstructs.Target

func PostTargetsHand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var Target DTOstructs.DTOTarget
	if err := json.NewDecoder(r.Body).Decode(&Target); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	NewTarget, err := Target.ValidateOnCreate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	AllTargets.AddTarget(NewTarget)

	responce, err := json.MarshalIndent(NewTarget, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Success"))
		w.Write(responce)
		pp.Println(string(responce))
	}

}

func DeleteTargetsHand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	IdTarget, err := IDstructs.FindById(id, AllTargets)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	responce, err := json.MarshalIndent(IdTarget, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(responce)
	}
}

func DelTarget(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := IDstructs.DeleteById(id, &AllTargets); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

func GetTargets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, err := json.MarshalIndent(AllTargets.Targets, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		panic(err)
	}

}

func StartServer() {
	router := mux.NewRouter()
	router.Path("/targets").Methods("POST").HandlerFunc(PostTargetsHand)
	router.Path("/targets/{id}").Methods("GET").HandlerFunc(DeleteTargetsHand)
	router.Path("/targets").Methods("GET").HandlerFunc(GetTargets)
	router.Path("/targets/{id}").Methods("DELETE").HandlerFunc(DelTarget)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
