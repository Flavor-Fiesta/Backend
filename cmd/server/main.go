package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfcheca/FlavorFiesta/cmd/server/handler"
	"github.com/jfcheca/FlavorFiesta/internal/carritos"
	"github.com/jfcheca/FlavorFiesta/internal/imagenes"
	"github.com/jfcheca/FlavorFiesta/internal/productos"
	"github.com/jfcheca/FlavorFiesta/internal/usuarios"
	"github.com/jfcheca/FlavorFiesta/internal/ordenes"
	"github.com/jfcheca/FlavorFiesta/pkg/store"

	//	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/domain"
	//	"github.com/joho/godotenv"
	"io/ioutil"
	"strings"
)

func main() {
	/*	// Cargar variables de entorno desde el archivo .env
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error al cargar el archivo .env:", err)
		}*/

	// Abrir una conexión temporal a MySQL para ejecutar comandos administrativos
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/FlavorFiesta")
	if err != nil {
		log.Fatal("Error al conectar con MySQL:", err)
	}
	defer db.Close()

	// Eliminar la base de datos 'FlavorFiesta' si ya existe
	_, err = db.Exec("DROP DATABASE IF EXISTS FlavorFiesta")
	if err != nil {
		log.Fatal("Error al eliminar la base de datos 'FlavorFiesta':", err)
	}

	// Crear la base de datos 'FlavorFiesta'
	_, err = db.Exec("CREATE DATABASE FlavorFiesta")
	if err != nil {
		log.Fatal("Error al crear la base de datos 'FlavorFiesta':", err)
	}

	// Conectar a la base de datos 'FlavorFiesta'
	bd, err := sql.Open("mysql", "root:root@(localhost:3306)/FlavorFiesta")
	if err != nil {
		log.Fatal("Error al conectar con la base de datos 'FlavorFiesta':", err)
	}
	defer bd.Close()

	// Cargar contenido del archivo schema.sql
	sqlFile, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error al leer el archivo schema.sql:", err)
	}

	// Dividir el contenido en sentencias SQL individuales
	sqlStatements := strings.Split(string(sqlFile), ";")

	// Ejecutar cada sentencia SQL en el archivo schema.sql
	for _, statement := range sqlStatements {
		// Limpiar la sentencia SQL
		cleanedStatement := strings.TrimSpace(statement)
		if cleanedStatement == "" {
			continue
		}

		_, err := bd.Exec(cleanedStatement)
		if err != nil {
			log.Fatal("Error al ejecutar la sentencia SQL:", err)
		}
	}

	// Configurar el enrutador Gin
	r := gin.Default()

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>PRODUCTOS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storage := store.NewSqlStoreProductos(bd)
	repo := productos.NewRepository(storage)
	service := productos.NewService(repo)
	productoHandler := handler.NewProductHandler(service)

	// Rutas para el manejo de productos
	productos := r.Group("/productos")
	{
		productos.GET("/:id", productoHandler.BuscarProducto())
		productos.GET("/", productoHandler.GetAll())
		productos.POST("/crear", productoHandler.Post())
		productos.DELETE("/:id", productoHandler.Delete())
		productos.PATCH("/:id", productoHandler.Patch())
		productos.PUT("/:id", productoHandler.Put())
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>IMAGENES>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageImagen := store.NewSqlStoreImagen(bd)
	repoImagen := imagenes.NewRepository(storageImagen)
	serviceImagen := imagenes.NewService(repoImagen)
	imagenHandler := handler.NewImagenHandler(serviceImagen)

	// Rutas para el manejo de imagenes
	imagenes := r.Group("/imagenes")
	{
		imagenes.GET("/:id", imagenHandler.GetByID())
		imagenes.POST("/crear", imagenHandler.Post())
		imagenes.DELETE("/:id", imagenHandler.Delete())
		imagenes.PATCH("/:id", imagenHandler.Patch())
		imagenes.PUT("/:id", imagenHandler.Put())
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>USUARIOS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageUsuario := store.NewSqlStoreUsuarios(bd)
	repoUsuario := usuarios.NewRepository(storageUsuario)
	serviceUsuario := usuarios.NewService(repoUsuario)
	usuariosHandler := handler.NewUsuarioHandler(serviceUsuario)

	// Rutas para el manejo de imagenes
	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("/:id", usuariosHandler.GetByID())
		usuarios.GET("/", usuariosHandler.GetAll())
		usuarios.POST("/crear", usuariosHandler.Post())
		usuarios.DELETE("/:id", usuariosHandler.DeleteUsuario())
		usuarios.PATCH("/:id", usuariosHandler.Patch())
		usuarios.PUT("/:id", usuariosHandler.Put())
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>CARRITOS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageCarrito := store.NewSqlStoreCarrito(bd)
	repoCarrito := carritos.NewRepository(storageCarrito)
	serviceCarrito := carritos.NewService(repoCarrito)
	carritoHandler := handler.NewCarritoHandler(serviceCarrito)

	// Rutas para el manejo de carritos
	carritos := r.Group("/carritos")
	{
		carritos.GET("/:id", carritoHandler.GetCarritoByID())
		carritos.POST("/crear", carritoHandler.Post())
		carritos.PUT("/:id", carritoHandler.Put())
		carritos.DELETE("/:id", carritoHandler.Delete())
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ORDENES>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageOrden := store.NewSqlStoreOrden(bd)
	repoOrden := orden.NewRepository(storageOrden)
	serviceOrden := orden.NewService(repoOrden)
	ordenHandler := handler.NewOrdenHandler(serviceOrden)

	// Rutas para el manejo de órdenes
	ordenes := r.Group("/ordenes")
	{
		ordenes.GET("/:id", ordenHandler.GetOrdenByID())
		ordenes.POST("/crear", ordenHandler.Post())
		ordenes.PUT("/:id", ordenHandler.Put())
		ordenes.DELETE("/:id", ordenHandler.Delete())
	}

	/*
	   	storageProductoImagen := store.NewSqlStoreProductoImagen(bd)
	   	repoProductoImagen := productoImagen.NewRepositoryProductoImagen(storageProductoImagen)
	   	ServiceProductoImagen:= productoImagen.NewServiceProductoImagen(repoProductoImagen)
	   	productoImagenHandler := handler.NewProductoImagenHandler(ServiceProductoImagen)

	   	productoImagen := r.Group("/productoImagen")
	   	{
	   		productoImagen.GET("/:id", productoImagenHandler.BuscarProductoImagen())
	   		productoImagen.POST("/crear", productoImagenHandler.CrearProductoImagen())
	   		productoImagen.DELETE("/:id", productoImagenHandler.DeleteProductoImagen())
	       //	pacientes.PATCH("/:id", pacienHandler.PatchPaciente())
	   	productoImagen.PUT("/:id", productoImagenHandler.UpdateProductoImagen())
	   	}
	*/
	// Ejecutar el servidor en el puerto 8080
	r.Run(":8080")
}
