TRUNCATE TABLE productos;CREATE TABLE `categorias` (
  `id_categorias` int NOT NULL,
  `nombre` varchar(45) NOT NULL,
  PRIMARY KEY (`id_categorias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `clientes` (
  `id_clientes` int NOT NULL,
  `Nombre` varchar(45) NOT NULL,
  `Apellido` varchar(45) NOT NULL,
  `Direccion` varchar(150) NOT NULL,
  `Telefono` varchar(15) NOT NULL,
  `Tarjeta` varchar(20) NOT NULL,
  `Edad` int NOT NULL,
  `Salario` int NOT NULL,
  `Genero` char(1) NOT NULL,
  `id_pais` int NOT NULL,
  PRIMARY KEY (`id_clientes`),
  KEY `fk_Clientes_Países1_idx` (`id_pais`) USING BTREE,
  CONSTRAINT `fk_Clientes_Países1` FOREIGN KEY (`id_pais`) REFERENCES `países` (`id_pais`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `detalle_ordenes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `linea_orden` int NOT NULL,
  `cantidad` int NOT NULL,
  `id_orden` int NOT NULL,
  `id_vendedor` int NOT NULL,
  `id_producto` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Detalle_Ordenes_Ordenes_idx` (`id_orden`),
  KEY `fk_Detalle_Ordenes_Vendedores1_idx` (`id_vendedor`),
  KEY `fk_Detalle_Ordenes_Productos1_idx` (`id_producto`) USING BTREE,
  CONSTRAINT `fk_Detalle_Ordenes_Ordenes` FOREIGN KEY (`id_orden`) REFERENCES `ordenes` (`id_orden`),
  CONSTRAINT `fk_Detalle_Ordenes_Productos1` FOREIGN KEY (`id_producto`) REFERENCES `productos` (`id_producto`),
  CONSTRAINT `fk_Detalle_Ordenes_Vendedores1` FOREIGN KEY (`id_vendedor`) REFERENCES `vendedores` (`id_vendedor`)
) ENGINE=InnoDB AUTO_INCREMENT=60353 DEFAULT CHARSET=utf8mb3;

CREATE TABLE `ordenes` (
  `id_orden` int NOT NULL,
  `fecha_orden` date NOT NULL,
  `id_clientes` int NOT NULL,
  PRIMARY KEY (`id_orden`),
  KEY `fk_Ordenes_Clientes1_idx` (`id_clientes`),
  CONSTRAINT `fk_Ordenes_Clientes1` FOREIGN KEY (`id_clientes`) REFERENCES `clientes` (`id_clientes`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `países` (
  `id_pais` int NOT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_pais`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `productos` (
  `id_producto` int NOT NULL,
  `Nombre` varchar(100) NOT NULL,
  `Precio` double NOT NULL DEFAULT '0',
  `id_categorias` int NOT NULL,
  PRIMARY KEY (`id_producto`),
  KEY `fk_Productos_Categorias1_idx` (`id_categorias`),
  CONSTRAINT `fk_Productos_Categorias1` FOREIGN KEY (`id_categorias`) REFERENCES `categorias` (`id_categorias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `vendedores` (
  `id_vendedor` int NOT NULL,
  `nombre` varchar(80) NOT NULL,
  `id_pais` int NOT NULL,
  PRIMARY KEY (`id_vendedor`),
  KEY `fk_Vendedores_Países1_idx` (`id_pais`),
  CONSTRAINT `fk_Vendedores_Países1` FOREIGN KEY (`id_pais`) REFERENCES `países` (`id_pais`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;