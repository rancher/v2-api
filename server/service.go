package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rancher/go-rancher/client"
	"github.com/rancher/v2-api/model"
)

func (s *Server) ServiceByID(rw http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	return s.getService(rw, r, vars["id"])
}

func (s *Server) ServiceList(rw http.ResponseWriter, r *http.Request) error {
	return s.getService(rw, r, "")
}

func (s *Server) getService(rw http.ResponseWriter, r *http.Request, id string) error {
	id = s.deobfuscate(r, "service", id)

	rows, err := s.namedQuery(s.getServicesSQL(r, id), map[string]interface{}{
		"account_id": s.getAccountID(r),
		"id":         id,
	})
	if err != nil {
		return err
	}
	defer rows.Close()

	response := &client.GenericCollection{
		Collection: client.Collection{
			Type:         "collection",
			ResourceType: "service",
		},
	}

	for rows.Next() {
		obj := &model.Service{}
		obj.Type = "service"

		var data string

		// The scan must match the query exactly
		if err := rows.Scan(&obj.Name, &obj.Id, &obj.UUID, &data); err != nil {
			return err
		}

		// Obfuscate Ids
		obj.Id = s.obfuscate(r, "service", obj.Id)

		// Probably do something more with data

		response.Data = append(response.Data, obj)
	}

	return s.writeResponse(rows.Err(), r, response)
}

func (s *Server) getServicesSQL(r *http.Request, id string) string {
	q := `
		SELECT
			name, id, uuid, data 
		FROM service
		WHERE
			account_id = :account_id
			AND removed IS NULL`

	if id != "" {
		q += " AND id = :id"
	}

	return q
}

func (s *Server) ServiceCreate(rw http.ResponseWriter, r *http.Request) error {
	rancherClient, err := s.getClient(r)
	if err != nil {
		return err
	}

	data := s.parseInputParameters(r)

	service, err := rancherClient.Service.Create(&client.Service{
		Name: data.String("name"),
	})

	if err != nil {
		return err
	}

	return s.getService(rw, r, service.Id)
}
