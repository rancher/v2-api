package server

import (
	"net/http"

	"github.com/rancher/go-rancher/api"
	"github.com/rancher/go-rancher/client"
)

const (
	stackTable = "environment"
)

func (s *Server) StackList(rw http.ResponseWriter, r *http.Request) error {
	context := api.GetApiContext(r)
	context.Write(&client.GenericCollection{
		Data: []interface{}{
			client.Resource{
				Id:   "foo",
				Type: "stack",
			},
			client.Resource{
				Id:   "bar",
				Type: "stack",
			},
		},
	})
	return nil
}
