package main

import (
	"github.com/pscn/go-vue-starter/api"
	"github.com/pscn/go-vue-starter/models"
	"github.com/pscn/go-vue-starter/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
