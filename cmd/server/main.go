package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfcheca/FlavorFiesta/cmd/server/handler"
	//"github.com/jfcheca/FlavorFiesta/cmd/server/middleware"
    "github.com/jfcheca/FlavorFiesta/internal/favoritos"
    "github.com/jfcheca/FlavorFiesta/internal/auth"
	"github.com/jfcheca/FlavorFiesta/internal/categorias"
	"github.com/jfcheca/FlavorFiesta/internal/estados"
	"github.com/jfcheca/FlavorFiesta/internal/imagenes"
	"github.com/jfcheca/FlavorFiesta/internal/ordenProducto"
	"github.com/jfcheca/FlavorFiesta/internal/ordenes"
	"github.com/jfcheca/FlavorFiesta/internal/productos"
	"github.com/jfcheca/FlavorFiesta/internal/roles"
	"github.com/jfcheca/FlavorFiesta/internal/usuarios"
 //   "github.com/jfcheca/FlavorFiesta/internal/favoritos"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
	"github.com/joho/godotenv"
	// "github.com/jfcheca/FlavorFiesta/internal/auth"
	//	"gopkg.in/mail.v2"

	"io/ioutil"
	"strings"
)

/*func enviarCorreoDePrueba() error {
    m := mail.NewMessage()
    m.SetHeader("From", "")
    m.SetHeader("To", "")
    m.SetHeader("Subject", "Correo de Prueba")
    m.SetBody("text/plain", "Este es un correo de prueba para verificar la configuración SMTP.")

    d := mail.NewDialer("smtp.gmail.com", 587, "", "")

    if err := d.DialAndSend(m); err != nil {
        return fmt.Errorf("error sending test email: %w", err)
    }
    return nil
}*/

func main() {

 /*   if err := enviarCorreoDePrueba(); err != nil {
        log.Fatalf("Error sending test email: %v", err)
    }
    log.Println("Correo de prueba enviado exitosamente")
*/
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CARGAMOS LAS VARIABLES DE ENTORNO DEL ARCHIVO .ENV >>>>>>>>>>>>>>>>>>>>
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error al cargar el archivo .env:", err)
    }

    log.Printf("SMTP_EMAIL: %s", os.Getenv("SMTP_EMAIL"))
    log.Printf("SMTP_PASSWORD: %s", os.Getenv("SMTP_PASSWORD"))
    
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Abrir una conexión temporal a MySQL para ejecutar comandos administrativos
    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error al conectar con MySQL:", err)
    }
    defer db.Close()

    // Eliminar la base de datos si ya existe
    _, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)
    if err != nil {
        log.Fatal("Error al eliminar la base de datos '" + dbName + "':", err)
    }

    // Crear la base de datos
    _, err = db.Exec("CREATE DATABASE " + dbName)
    if err != nil {
        log.Fatal("Error al crear la base de datos '" + dbName + "':", err)
    }

    // Conectar a la base de datos
    dsn = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
    bd, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error al conectar con la base de datos '" + dbName + "':", err)
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
    r.Static("/Probando", "./public")

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // URL del frontend
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Definir rutas
    r.GET("/api/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

////////////////////////////////////////// >>>>>>>>>>>>>> TODO LO REFERIDO A LA AUTENTICACION >>>>>>>>>>>>>>>>>>>>>>>>>>>>
 /*   r.POST("/api/login", func(c *gin.Context) {
        var credentials auth.Credentials
        if err := c.ShouldBindJSON(&credentials); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        token, err := auth.Authenticate(credentials)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    })
*/


 //Middleware de autenticación
  // r.Use(middleware.AuthMiddleware())

    // >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PRODUCTOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
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

    // >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> IMAGENES <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
    storageImagen := store.NewSqlStoreImagen(bd)
    repoImagen := imagenes.NewRepository(storageImagen)
    serviceImagen := imagenes.NewService(repoImagen)
    imagenHandler := handler.NewImagenHandler(serviceImagen)

    // Rutas para el manejo de imágenes
    imagenes := r.Group("/imagenes")
    {
        imagenes.GET("/:id", imagenHandler.GetByID())
        imagenes.POST("/crear", imagenHandler.Post())
        imagenes.DELETE("/:id", imagenHandler.Delete())
        imagenes.PATCH("/:id", imagenHandler.Patch())
        imagenes.PUT("/:id", imagenHandler.Put())
    }


      // Authent
      storageAuth := store.NewSqlStoreUsuarios(bd) // Reuse user store for authentication
      repoAuth := auth.NewRepository(storageAuth)
      serviceAuth := auth.NewService(repoAuth)
      storageUsuario := store.NewSqlStoreUsuarios(bd)
      repoUsuario := usuarios.NewRepository(storageUsuario)
      serviceUsuario := usuarios.NewService(repoUsuario)
      usuariosHandler := handler.NewUsuarioHandler(serviceUsuario, serviceAuth)
      authHandler := handler.NewAuthHandler(serviceAuth, serviceUsuario)
  
      authRoutes := r.Group("/auth")
      {
          authRoutes.POST("/login", authHandler.Login())
          authRoutes.POST("/forgotPassword", authHandler.ForgotPassword())
      }

    // >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> USUARIOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<


    // Rutas para el manejo de usuarios
    usuarios := r.Group("/usuarios")
    {
        usuarios.GET("/:id", usuariosHandler.GetByID())
        usuarios.GET("/email&pass", usuariosHandler.GetByEmailAndPassword())
        usuarios.GET("/email&passdatos", usuariosHandler.GetByEmailAndPasswordConDatos())
        usuarios.GET("/", usuariosHandler.GetAll())
        usuarios.POST("/crear", usuariosHandler.Post())
        usuarios.DELETE("/:id", usuariosHandler.DeleteUsuario())
        usuarios.PUT("/:id", usuariosHandler.Put())
        usuarios.PATCH("/:id", usuariosHandler.Patch())
        usuarios.PUT("/forgotPassword/:id", usuariosHandler.UpdatePassword())
    }

    // >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CATEGORIAS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
    storageCategoria := store.NewSqlStoreCategorias(bd)
    repoCategoria := categorias.NewRepository(storageCategoria)
    serviceCategoria := categorias.NewService(repoCategoria)
    categoriasHandler := handler.NewCategoriaHandler(serviceCategoria)

    // Rutas para el manejo de categorías
    categorias := r.Group("/categorias")
    {
        categorias.GET("/:id", categoriasHandler.GetByID())
        categorias.GET("/", categoriasHandler.GetAll())
        categorias.POST("/crear", categoriasHandler.Post())
        categorias.DELETE("/:id", categoriasHandler.DeleteCategoria())
        categorias.PATCH("/:id", categoriasHandler.Patch())
        categorias.PUT("/:id", categoriasHandler.Put())
    }

    //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ORDENES>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageOrden := store.NewSqlStoreOrden(bd)
	repoOrden := ordenes.NewRepository(storageOrden)
	serviceOrden := ordenes.NewService(repoOrden)
	ordenHandler := handler.NewOrdenHandler(serviceOrden)

	// Rutas para el manejo de órdenes
	ordenes := r.Group("/ordenes")
	{
		ordenes.GET("/:id", ordenHandler.GetOrdenByID())
        ordenes.GET("/usuario&estado", ordenHandler.GetOrdenByUserIDyOrden())
        ordenes.GET("/usuario&estadoConDatos", ordenHandler.GetOrdenByUsuarioYEstadoConDatos())
		ordenes.POST("/crear", ordenHandler.Post())
		ordenes.PUT("/:id", ordenHandler.Put())
		ordenes.DELETE("/:id", ordenHandler.Delete())
	}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ORDEN PRODUCTO>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
	storageOrdenProducto := store.NewSqlStoreOrdenProducto(bd)
	repoOrdenProducto := ordenProductos.NewRepository(storageOrdenProducto)
	serviceOrdenProducto := ordenProductos.NewService(repoOrdenProducto)
	ordenProductoHandler := handler.NewOrdenProductoHandler(serviceOrdenProducto)

	// Rutas para el manejo de órdenes
	ordenProductos := r.Group("/ordenProductos")
	{
		ordenProductos.GET("/:id", ordenProductoHandler.GetByID())
        ordenProductos.GET("/orden/:idOrden", ordenProductoHandler.BuscarPorIDOrden())
        ordenProductos.GET("/", ordenProductoHandler.GetAll())
		ordenProductos.POST("/crear", ordenProductoHandler.Post())
		ordenProductos.PUT("/:id", ordenProductoHandler.Put())
	}

    //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ESTADOS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
    storageEstado := store.NewSqlStoreEstados(bd)
	repoEstado := estados.NewRepository(storageEstado)
	serviceEstado := estados.NewService(repoEstado)
    estadoHandler := handler.NewEstadoHandler(serviceEstado)

    // Rutas para el manejo de órdenes
    estados := r.Group("/estados")
    {
        estados.GET("/:id", estadoHandler.BuscarEstado())
        estados.GET("/", estadoHandler.GetAll())
        estados.POST("/crear", estadoHandler.Post())
    }

   //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FAVORITOS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Crear el almacenamiento SQL con la base de datos 'FlavorFiesta'
    storageFavoritos := store.NewSqlStoreFavoritos(db)
    repoFavoritos := favoritos.NewRepository(storageFavoritos)
    serviceFavoritos := favoritos.NewServiceFavoritos(repoFavoritos)
    favoritoHandler := handler.NewFavoritosHandler(serviceFavoritos)

    // Rutas para el manejo de órdenes
    favoritos := r.Group("/favoritos")
    {
       
        favoritos.POST("/agregar", favoritoHandler.Post())
    }



    // >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ROLES <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
    storageRol := store.NewSqlStoreRoles(bd)
    repoRol := roles.NewRepository(storageRol)
    serviceRol := roles.NewService(repoRol)
    rolHandler := handler.NewRolHandler(serviceRol)

    // Rutas para el manejo de roles
    roles := r.Group("/roles")
    {
        roles.GET("/", rolHandler.GetAll())
        roles.POST("/crear", rolHandler.Post())
    }

  
    
 /*   // Endpoints protegidos con middleware de rol ADMIN
    adminRoutes := r.Group("/admin")
    adminRoutes.Use(middleware.AdminRoleMiddleware())
    {
        adminRoutes.PUT("/roles/cambiar", rolHandler.CambiarRol())
    }
*/

    // Ejecutar el servidor en el puerto 8080
    r.Run(":8080")

    
}