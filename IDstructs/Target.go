package IDstructs

import (
	"errors"
)

var status bool

type Target struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Active bool   `json:"active"`
}

type ListTargets struct {
	Targets []Target
}

func (t *ListTargets) AddTarget(target Target) {
	t.Targets = append(t.Targets, target)
}

func FindById(id int, targets ListTargets) (Target, error) {
	status = false
	for _, target := range targets.Targets {
		if target.ID == id {
			status = true
			return target, nil
		}
	}
	return Target{}, errors.New("Target not found")

}

func DeleteById(id int, targets *ListTargets) error {
	status = false
	for i, target := range targets.Targets {
		if target.ID == id {
			status = true
			targets.Targets = append(targets.Targets[:i], targets.Targets[i+1:]...)
			return nil
		}
	}
	return errors.New("Target not found")
}
