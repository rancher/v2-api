package model

import (
	"github.com/rancher/go-rancher/client"
)

func NewSchema() *client.Schemas {
	schemas := &client.Schemas{}

	schemas.AddType("apiVersion", client.Resource{})
	schemas.AddType("schema", client.Schema{})
	schemas.AddType("service", Service{})

	return schemas
}
