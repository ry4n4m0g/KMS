package main

import (
	"log"

	"github.com/epointpayment/key_management_system/app/controllers"
	"github.com/epointpayment/key_management_system/app/database"
	"github.com/epointpayment/key_management_system/app/router"
	"github.com/epointpayment/key_management_system/app/services"

	"github.com/namsral/flag"
)

var dsn string

func init() {
	flag.StringVar(&dsn, "dsn", "", "path to sample data i.e. user:pass@/example")

	flag.Parse()
}

func main() {
	var err error

	db := database.NewDatabase("mysql", dsn)
	err = db.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = services.New(db)
	if err != nil {
		log.Fatalln(err)
	}

	c := controllers.NewControllers()
	r := router.NewRouter(c)
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
