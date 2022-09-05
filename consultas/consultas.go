package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// STRUCTS PARA JSONS DE CONSULTAS
type Consulta1 struct {
	Id_cliente int    `json:"ID"`
	Nombre     string `json:"NOMBRE"`
	Apellido   string `json:"APELLIDO"`
	Pais       string `json:"PAIS"`
	Monto      string `json:"MONTO_TOTAL"`
	No_Compras string `json:"veces_compra"`
}
type Consulta2 struct {
	Id        int     `json:"ID"`
	Nombre    string  `json:"NOMBRE"`
	Categoria string  `json:"CATEGORIA"`
	Unidades  int     `json:"UNIDADES"`
	Monto     float32 `json:"MONTO"`
}

type Consulta3 struct {
	Id_vendedor int     `json:"ID"`
	Nombre      string  `json:"NOMBRE"`
	Monto       float32 `json:"MONTO VENDIDO"`
}

type Consulta4 struct {
	Nombre string  `json:"NOMBRE"`
	Monto  float32 `json:"MONTO VENDIDO"`
}

type Consulta5 struct {
	Id_pais int     `json:"ID"`
	Nombre  string  `json:"NOMBRE"`
	Monto   float32 `json:"MONTO"`
}

type Consulta6 struct {
	Nombre   string `json:"NOMBRE"`
	Unidades int    `json:"CANTIDAD UNIDADES"`
}

type Consulta7 struct {
	Pais     string `json:"PAIS"`
	Nombre   string `json:"NOMBRE"`
	Cantidad int    `json:"CANTIDAD"`
}

type Consulta8 struct {
	Pais  string  `json:"PAIS"`
	Mes   int     `json:"MES"`
	Monto float32 `json:"MONTO"`
}

type Consulta9 struct {
	Mes   int     `json:"MES"`
	Monto float32 `json:"MONTO"`
}

type Consulta10 struct {
	Id_producto int     `json:"ID"`
	Nombre      string  `json:"NOMBRE"`
	Categoria   string  `json:"CATEGORIA"`
	Monto       float32 `json:"MONTO"`
}

var conn = MySQLConnection()

func MySQLConnection() *sql.DB {
	usuario := "root"
	pass := 503802
	host := "tcp(localhost:3306)"
	db := "proyecto1"
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%d@%s/%s", usuario, pass, host, db))
	if err != nil {
		fmt.Println("HAY ERROR: \n", err)
	} else {
		fmt.Println("se ha conectado a mysql!")
	}
	return conn
}

/*
Mostrar el cliente que más ha comprado. Se debe de mostrar el id del cliente,
nombre, apellido, país y monto total
*/
func Reporte1(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var listUsr []Consulta1
	query := `SELECT c.id_clientes AS ID,c.Nombre AS NOMBRE,c.Apellido AS APELLIDO,p.nombre AS PAIS,SUM(productos.Precio*d.cantidad) AS MONTO_TOTAL, COUNT(c.id_clientes) AS veces_compra
	FROM clientes c
	INNER JOIN ordenes o
	ON c.id_clientes = o.id_clientes
	INNER JOIN detalle_ordenes d
	ON o.id_orden = d.id_orden
	INNER JOIN productos
	ON d.id_producto = productos.id_producto
	INNER JOIN países p
	ON c.id_pais = p.id_pais
	GROUP BY c.id_clientes
	ORDER BY veces_compra DESC
	LIMIT 1;`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var usr Consulta1
		err := result.Scan(&usr.Id_cliente, &usr.Nombre, &usr.Apellido, &usr.Pais, &usr.Monto, &usr.No_Compras)
		if err != nil {
			fmt.Println(err)
		}
		listUsr = append(listUsr, usr)
	}
	json.NewEncoder(response).Encode(listUsr)
}

/*
Mostrar el producto más y menos comprado. Se debe mostrar el id del
producto, nombre del producto, categoría, cantidad de unidades y monto
vendido
*/
func Reporte2(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta2
	query := `(SELECT * 
		FROM unidades_producto 
		ORDER BY unidades ASC 
		LIMIT 1 )
	UNION
	(SELECT * 
		FROM unidades_producto may
		ORDER BY unidades DESC
		LIMIT 1);`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta2
		err := result.Scan(&c.Id, &c.Nombre, &c.Categoria, &c.Unidades, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar a la persona que más ha vendido. Se debe mostrar el id del
vendedor, nombre del vendedor, monto total vendido
*/
func Reporte3(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta3
	query := `SELECT v.id_vendedor AS Id_vendedor, v.nombre as Nombre,SUM(d.cantidad * p.Precio) AS Monto
	FROM vendedores v
	INNER JOIN detalle_ordenes d
	ON v.id_vendedor = d.id_orden
	INNER JOIN productos p
	ON d.id_producto = p.id_producto
	GROUP BY v.id_vendedor
	ORDER BY Monto DESC
	LIMIT 1;`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta3
		err := result.Scan(&c.Id_vendedor, &c.Nombre, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar el país que más y menos ha vendido. Debe mostrar el nombre del
país y el monto. (Una sola consulta).
*/
func Reporte4(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta4
	query := `(SELECT * 
		FROM ventas_por_pais p
		ORDER BY p.monto ASC
		LIMIT 1
	)
	UNION
	-- MAYOR 
	(SELECT *
		FROM ventas_por_pais p
		ORDER BY p.monto DESC
		LIMIT 1
	);`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta4
		err := result.Scan(&c.Nombre, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Top 5 de países que más han comprado en orden ascendente. Se le solicita
mostrar el id del país, nombre y monto total
*/
func Reporte5(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta5
	query := `SELECT p.id_pais, p.nombre,  SUM(d.cantidad*a.Precio) AS monto
	FROM países p
	INNER JOIN clientes v
	ON p.id_pais = v.id_pais
	INNER JOIN ordenes o
	ON v.id_clientes = o.id_clientes
	INNER JOIN detalle_ordenes d
	ON o.id_orden = d.id_orden 
	INNER	JOIN productos a
	ON a.id_producto = d.id_producto
	GROUP BY p.nombre
	ORDER BY monto ASC
	LIMIT 5;`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta5
		err := result.Scan(&c.Id_pais, &c.Nombre, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar la categoría que más y menos se ha comprado. Debe de mostrar el
nombre de la categoría y cantidad de unidades. (Una sola consulta).
*/
func Reporte6(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta6
	query := `(SELECT c.nombre,  SUM(d.cantidad) AS unidades
	FROM categorias c
	INNER JOIN productos p
	ON c.id_categorias = p.id_categorias
	INNER JOIN detalle_ordenes d
	ON p.id_producto = d.id_producto
	GROUP BY c.nombre 
	ORDER BY unidades DESC
	LIMIT 1)
	UNION 
	-- MIN
	(SELECT c.nombre,  SUM(d.cantidad) AS unidades
	FROM categorias c
	INNER JOIN productos p
	ON c.id_categorias = p.id_categorias
	INNER JOIN detalle_ordenes d
	ON p.id_producto = d.id_producto
	GROUP BY c.nombre 
	ORDER BY unidades ASC
	LIMIT 1);`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta6
		err := result.Scan(&c.Nombre, &c.Unidades)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar la categoría más comprada por cada país. Se debe de mostrar el
nombre del país, nombre de la categoría y cantidad de unidades.
*/
func Reporte7(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta7
	query := `SELECT a.*
	FROM compras_pais a
	INNER JOIN 
	(
		SELECT c.pais, c.nombre, MAX(c.cantidad) AS max_cantidad 
		FROM compras_pais c 
		GROUP BY c.pais
	) b
	ON a.pais = b.pais  AND a.cantidad = b.max_cantidad`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta7
		err := result.Scan(&c.Pais, &c.Nombre, &c.Cantidad)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar las ventas por mes de Inglaterra. Debe de mostrar el número del mes
y el monto.
*/
func Reporte8(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta8
	query := `SELECT p.nombre AS pais, MONTH(o.fecha_orden) AS mes, SUM(d.cantidad * pr.Precio) AS monto
	FROM vendedores v
	INNER JOIN países p
	ON v.id_pais = p.id_pais 
	INNER JOIN detalle_ordenes d
	ON v.id_vendedor = d.id_vendedor
	INNER JOIN ordenes o
	ON d.id_orden = o.id_orden
	INNER JOIN productos pr
	ON d.id_producto = pr.id_producto
	WHERE p.nombre = 'Inglaterra'
	GROUP BY mes,pais;`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta8
		err := result.Scan(&c.Pais, &c.Mes, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar el mes con más y menos ventas. Se debe de mostrar el número de
mes y monto. (Una sola consulta).
*/
func Reporte9(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta9
	query := `(SELECT MONTH(o.fecha_orden) AS mes, SUM(d.cantidad * p.Precio) AS monto
	FROM ordenes o
	INNER JOIN detalle_ordenes d
	ON o.id_orden = d.id_orden
	INNER JOIN productos p
	ON d.id_producto = p.id_producto
	GROUP BY mes
	ORDER BY monto DESC
	LIMIT 1)
	UNION 
	(SELECT MONTH(o.fecha_orden) AS mes, SUM(d.cantidad * p.Precio) AS monto
	FROM ordenes o
	INNER JOIN detalle_ordenes d
	ON o.id_orden = d.id_orden
	INNER JOIN productos p
	ON d.id_producto = p.id_producto
	GROUP BY mes
	ORDER BY monto ASC 
	LIMIT 1)`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta9
		err := result.Scan(&c.Mes, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}

/*
Mostrar las ventas de cada producto de la categoría deportes. Se debe de
mostrar el id del producto, nombre y monto
*/
func Reporte10(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var list []Consulta10
	query := `SELECT p.id_producto, p.Nombre, c.nombre AS categoria, SUM(p.Precio * d.cantidad) AS monto
	FROM productos p
	INNER JOIN categorias c
	ON p.id_categorias = c.id_categorias
	INNER JOIN detalle_ordenes d
	ON p.id_producto = d.id_producto
	WHERE c.nombre = 'Deportes'
	GROUP BY p.Nombre ,categoria;`
	result, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		var c Consulta10
		err := result.Scan(&c.Id_producto, &c.Nombre, &c.Categoria, &c.Monto)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	json.NewEncoder(response).Encode(list)
}
