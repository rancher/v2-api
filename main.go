package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-sql-driver/mysql"
	"github.com/rancher/v2-api/router"
	"github.com/rancher/v2-api/server"
)

func main() {
	listen := ":8899"
	logrus.Info("Starting Rancher V2 API on ", listen)
	config := mysql.Config{
		User:      "cattle",
		Passwd:    "cattle",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "cattle",
		Collation: "utf8_general_ci",
	}
	server, err := server.New("mysql", config.FormatDSN())
	if err != nil {
		logrus.Fatal(err)
	}
	r := router.New(server)
	http.ListenAndServe(listen, r)
}
