//conexion hacia mysql desde go
//ojo colocar el password,ip para que funcione
/*----------------------------------
SE USARON LAS SIGUIENTES LIBRERIAS
go get github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
*/
package main

import (
	"fmt"
	"main/consultas"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	//Mostrar el cliente que más ha comprado. Se debe de mostrar el id del cliente,
	//nombre, apellido, país y monto total
	router.HandleFunc("/1", consultas.Reporte1).Methods("GET")
	router.HandleFunc("/2", consultas.Reporte2).Methods("GET")
	router.HandleFunc("/3", consultas.Reporte3).Methods("GET")
	router.HandleFunc("/4", consultas.Reporte4).Methods("GET")
	router.HandleFunc("/5", consultas.Reporte5).Methods("GET")
	router.HandleFunc("/6", consultas.Reporte6).Methods("GET")
	router.HandleFunc("/7", consultas.Reporte7).Methods("GET")
	router.HandleFunc("/8", consultas.Reporte8).Methods("GET")
	router.HandleFunc("/9", consultas.Reporte9).Methods("GET")
	router.HandleFunc("/10", consultas.Reporte10).Methods("GET")

	fmt.Println("Server on port", 8000)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}

}
