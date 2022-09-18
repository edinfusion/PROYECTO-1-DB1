-- REPORTE 1
	SELECT c.id_clientes AS ID,c.Nombre AS NOMBRE,c.Apellido AS APELLIDO,p.nombre AS PAIS,SUM(productos.Precio*d.cantidad) AS MONTO_TOTAL, 
	COUNT(c.id_clientes) AS veces_compra
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
	LIMIT 1;

-- REPORTE 2
CREATE VIEW unidades_producto AS
SELECT p.id_producto, p.Nombre, c.nombre AS categoria, SUM(d.cantidad) AS unidades, SUM(d.cantidad * p.Precio) AS 'monto total'
FROM productos p
INNER JOIN detalle_ordenes d
ON p.id_producto = d.id_producto
INNER JOIN categorias c
ON p.id_categorias = c.id_categorias
GROUP BY d.id_producto;

(SELECT * 
	FROM unidades_producto 
	ORDER BY unidades ASC 
	LIMIT 1 )
UNION
(SELECT * 
	FROM unidades_producto 
	ORDER BY unidades DESC
	LIMIT 1); 

-- REPORTE 3
SELECT v.id_vendedor, v.nombre,SUM(d.cantidad * p.Precio) AS `MONTO VENDIDO`
FROM vendedores v
INNER JOIN detalle_ordenes d
ON v.id_vendedor = d.id_vendedor
INNER JOIN productos p
ON d.id_producto = p.id_producto
GROUP BY v.id_vendedor
ORDER BY `MONTO VENDIDO` DESC
LIMIT 1;

-- REPORTE 4
CREATE VIEW ventas_por_pais AS
SELECT p.nombre,  SUM(d.cantidad*a.Precio) AS monto
FROM países p
INNER JOIN vendedores v
ON p.id_pais = v.id_pais
INNER JOIN detalle_ordenes d
ON v.id_vendedor = d.id_vendedor 
INNER	JOIN productos a
ON a.id_producto = d.id_producto
GROUP BY p.nombre;

-- MENOR
(SELECT * 
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
);

-- REPORTE 5
SELECT p.id_pais, p.nombre,  SUM(d.cantidad*a.Precio) AS monto
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
LIMIT 5;

-- REPORTE 6
-- MAX
(SELECT c.nombre,  SUM(d.cantidad) AS `cantidad unidades`
FROM categorias c
INNER JOIN productos p
ON c.id_categorias = p.id_categorias
INNER JOIN detalle_ordenes d
ON p.id_producto = d.id_producto
GROUP BY c.nombre 
ORDER BY `cantidad unidades` DESC
LIMIT 1)
UNION 
-- MIN
(SELECT c.nombre,  SUM(d.cantidad) AS `cantidad unidades`
FROM categorias c
INNER JOIN productos p
ON c.id_categorias = p.id_categorias
INNER JOIN detalle_ordenes d
ON p.id_producto = d.id_producto
GROUP BY c.nombre 
ORDER BY `cantidad unidades` ASC
LIMIT 1);

-- REPORTE 7
-- COMPRAS POR CATEGORIA DE CADA PAIS
CREATE VIEW compras_pais AS
SELECT p.nombre AS pais,ca.nombre,SUM(d.cantidad) AS cantidad
FROM clientes c
INNER JOIN países p
ON c.id_pais = p.id_pais
INNER JOIN ordenes o
ON c.id_clientes = o.id_clientes
INNER JOIN detalle_ordenes d
ON o.id_orden = d.id_orden
INNER  JOIN productos pr
ON d.id_producto = pr.id_producto
INNER JOIN categorias ca
ON pr.id_categorias = ca.id_categorias
GROUP BY p.nombre, ca.nombre
ORDER BY p.nombre

SELECT a.*
FROM compras_pais a
INNER JOIN 
(
	SELECT c.pais, c.nombre, MAX(c.cantidad) AS max_cantidad 
	FROM compras_pais c 
	GROUP BY c.pais
) b
ON a.pais = b.pais  AND a.cantidad = b.max_cantidad
ORDER BY a.cantidad ASC


-- REPORTE 8
SELECT p.nombre AS pais, MONTH(o.fecha_orden) AS mes, SUM(d.cantidad * pr.Precio) AS monto
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
GROUP BY mes,pais;

-- REPORTE 9
(SELECT MONTH(o.fecha_orden) AS mes, SUM(d.cantidad * p.Precio) AS monto
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
LIMIT 1)

-- REPORTE 10
SELECT p.id_producto, p.Nombre, c.nombre AS categoria, SUM(p.Precio * d.cantidad) AS monto
FROM productos p
INNER JOIN categorias c
ON p.id_categorias = c.id_categorias
INNER JOIN detalle_ordenes d
ON p.id_producto = d.id_producto
WHERE c.nombre = 'Deportes'
GROUP BY p.Nombre ,categoria;







	