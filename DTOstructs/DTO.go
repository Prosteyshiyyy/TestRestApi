package DTOstructs

import (
	"RestApi/IDstructs"
	"errors"
	"fmt"
	"math/rand"
)

type DTOTarget struct {
	Name string `json:"name"`
}

func (d DTOTarget) ValidateOnCreate() (IDstructs.Target, error) {
	if d.Name == "" {
		return IDstructs.Target{}, errors.New("Name is empty")
	}
	id := rand.Intn(10000)
	url := fmt.Sprintf("http://%v/%v/example", id, d.Name)

	return IDstructs.Target{
		Name:   d.Name,
		URL:    url,
		Active: true,
		ID:     id,
	}, nil
}
