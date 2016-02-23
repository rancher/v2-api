package model

import "time"

type ID string

type Common struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	UUID        string    `json:"uuid"`
	Created     time.Time `json:"created"`
	Removed     time.Time `json:"removed"`

	Transitioning        string `json:"transitioning"`
	TransitioningMessage string `json:"transitioningMessage"`
}
