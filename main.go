package main

import (
	"Martini/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/mux"
	"github.com/go-martini/martini"
)

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/users", controllers.GetAllUser).Methods("GET")
	// router.HandleFunc("/products", controllers.GetAllProduct).Methods("GET")
	// router.HandleFunc("/transactions", controllers.GetAllTransaction).Methods("GET")
	m := martini.Classic()

	m.Group("/users", func(r martini.Router) {
		r.Get("/print", controllers.GetAllUser)
		r.Post("/insert", controllers.InsertUser)
		r.Put("/update/:idUser", controllers.UpdateUser)
		r.Delete("/delete/:idUser", controllers.DeleteUser)
	})
	m.RunOnAddr(":8080")

	http.Handle("/", m)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
