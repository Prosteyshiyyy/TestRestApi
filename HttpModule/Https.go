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

func PostTargetsHand(w http.ResponseWriter, r *http.Request) {

	var Target DTOstructs.DTOTarget
	if err := json.NewDecoder(r.Body).Decode(&Target); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	NewTarget, err := Target.ValidateOnCreate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	AllTargets.AddTarget(NewTarget)

	responce, err := json.MarshalIndent(NewTarget, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(responce)
		pp.Println(string(responce))
	}

}

func DeleteTargetsHand(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	IdTarget, err := IDstructs.FindById(id, AllTargets)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responce, err := json.MarshalIndent(IdTarget, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(responce)
	}
}

func DelTarget(w http.ResponseWriter, r *http.Request) {

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
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

func GetTargets(w http.ResponseWriter, r *http.Request) {

	b, err := json.MarshalIndent(AllTargets.Targets, "", "    ")
	if err != nil {
		return
	}
	if _, err := w.Write(b); err != nil {
		return
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
