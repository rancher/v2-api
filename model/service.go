package model

import "github.com/rancher/go-rancher/client"

type Service struct {
	client.Resource
	Common

	ContainerIds       []ID                `json:"containerIds"`
	ContainerSelector  string              `json:"containerSelector"`
	ContainerTemplates []ContainerTemplate `json:"containerTemplates"`
	CreateIndex        string              `json:"createIndex"`
	Fqdn               string              `json:"fqdn"`
	HealthState        string              `json:"healthState"`
	HostnameOverride   string              `json:"hostnameOverride"`
	LinkSelector       string              `json:"linkSelector"`
	LinkedServiceIds   []ID                `json:"linkedServiceIds"`
	Metadata           string              `json:"metadata"`
	PublicEndpoints    []PublicEndpoint    `json:"publicEndpoints"`
	PullImage          string              `json:"pullImage"`
	RequestedIPAddress string              `json:"requestedIpAddress"`
	RetainIPAddress    bool                `json:"retainIpAddress"`
	Scale              int                 `json:"scale"`
	ServiceIPAddress   string              `json:"serviceIpAddress"`
	StackID            ID                  `json:"stackId"`
	StartOnce          bool                `json:"startOnce"`
}

type ServiceList struct {
	client.Collection
	Data []Service `json:"data,omitempty"`
}
