package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type sqlStoreOrdenProductos struct {
	db *sql.DB
}

func NewSqlStoreOrdenProducto(db *sql.DB) StoreInterfaceOrdenProducto {
	return &sqlStoreOrdenProductos{
		db: db,
	}
}

func (s *sqlStoreOrdenProductos) CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error) {
    query := "INSERT INTO OrdenProducto (id_orden, id_producto, cantidad, total) VALUES (?, ?, ?, ?);"
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return domain.OrdenProducto{}, fmt.Errorf("error preparing query: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(op.ID_Orden, op.ID_Producto, op.Cantidad, op.Total)
    if err != nil {
        return domain.OrdenProducto{}, fmt.Errorf("error executing query: %w", err)
    }

    // Retrieve the product details after creating the order-product
    productQuery := "SELECT id, nombre, descripcion, precio, stock, ranking, id_categoria FROM Productos WHERE id = ?"
    var product domain.Producto
    err = s.db.QueryRow(productQuery, op.ID_Producto).Scan(
        &product.ID, &product.Nombre, &product.Descripcion, &product.Precio, &product.Stock, &product.Ranking, &product.Id_categoria,
    )
    if err != nil {
        return domain.OrdenProducto{}, fmt.Errorf("error querying product details: %w", err)
    }

    op.Producto = product

    return op, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
/*func (s *sqlStoreOrdenProductos) CrearOrdenProducto(op domain.OrdenProducto) error {
	query := "INSERT INTO OrdenProducto (id_orden, id_producto, cantidad, total) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(op.ID_Orden, op.ID_Producto, op.Cantidad, op.Total)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}*/

// BuscarOrdenProductoPorID busca una relación entre orden y producto por su ID
func (s *sqlStoreOrdenProductos) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
	log.Printf("SQL Store: Buscando OrdenProducto con ID: %d", id)
	var op domain.OrdenProducto
	query := "SELECT id, id_orden, id_producto, cantidad, total FROM OrdenProducto WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&op.ID, &op.ID_Orden, &op.ID_Producto, &op.Cantidad, &op.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("SQL Store: OrdenProducto no encontrado")
			return domain.OrdenProducto{}, fmt.Errorf("order product not found")
		}
		log.Printf("SQL Store: Error ejecutando la consulta: %v", err)
		return domain.OrdenProducto{}, fmt.Errorf("error executing query: %w", err)
	}

	log.Printf("SQL Store: OrdenProducto encontrado: %+v", op)

	// Retrieve the product details
	productQuery := "SELECT id, nombre, descripcion, precio, stock, ranking, id_categoria FROM productos WHERE id = ?"
	var product domain.Producto
	err = s.db.QueryRow(productQuery, op.ID_Producto).Scan(
		&product.ID, &product.Nombre, &product.Descripcion, &product.Precio, &product.Stock, &product.Ranking, &product.Id_categoria,
	)
	if err != nil {
		log.Printf("SQL Store: Error al consultar los detalles del producto: %v", err)
		return domain.OrdenProducto{}, fmt.Errorf("error querying product details: %w", err)
	}

	op.Producto = product

	log.Printf("SQL Store: Detalles del producto encontrados: %+v", product)

	return op, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> FILTRAR ORDENPRODUCTO POR ID DE ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *sqlStoreOrdenProductos) BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error) {
    log.Printf("SQL Store: Buscando OrdenesProducto con ID de orden: %d", idOrden)
    var ordenesProducto []domain.OrdenProducto
    query := "SELECT id, id_orden, id_producto, cantidad, total FROM OrdenProducto WHERE id_orden = ?"

    rows, err := s.db.Query(query, idOrden)
    if err != nil {
        log.Printf("SQL Store: Error al ejecutar la consulta: %v", err)
        return nil, fmt.Errorf("error executing query: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var op domain.OrdenProducto
        if err := rows.Scan(&op.ID, &op.ID_Orden, &op.ID_Producto, &op.Cantidad, &op.Total); err != nil {
            log.Printf("SQL Store: Error al escanear fila: %v", err)
            return nil, fmt.Errorf("error scanning row: %w", err)
        }

        // Obtener detalles del producto
        productQuery := "SELECT id, nombre, descripcion, precio, stock, ranking, id_categoria FROM productos WHERE id = ?"
        var product domain.Producto
        err := s.db.QueryRow(productQuery, op.ID_Producto).Scan(
            &product.ID, &product.Nombre, &product.Descripcion, &product.Precio, &product.Stock, &product.Ranking, &product.Id_categoria,
        )
        if err != nil {
            log.Printf("SQL Store: Error al consultar detalles del producto: %v", err)
            return nil, fmt.Errorf("error querying product details: %w", err)
        }

        op.Producto = product
        ordenesProducto = append(ordenesProducto, op)
    }

    if err := rows.Err(); err != nil {
        log.Printf("SQL Store: Error al iterar filas: %v", err)
        return nil, fmt.Errorf("error iterating rows: %w", err)
    }

    return ordenesProducto, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCA TODAS LAS ORDENES PRODUCTOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (s *sqlStoreOrdenProductos) BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error) {
	var ordenesProductos []domain.OrdenProducto
	query := `
		SELECT 
			op.id, op.id_orden, op.id_producto, op.cantidad, op.total,
			p.id, p.nombre, p.descripcion, p.precio, p.stock, p.ranking, p.id_categoria, c.nombre AS categoria
		FROM 
			OrdenProducto op
		INNER JOIN 
			Productos p ON op.id_producto = p.id
		INNER JOIN 
			Categorias c ON p.id_categoria = c.id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying ordenes de productos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ordenProducto domain.OrdenProducto
		var producto domain.Producto
		if err := rows.Scan(
			&ordenProducto.ID, &ordenProducto.ID_Orden, &ordenProducto.ID_Producto, &ordenProducto.Cantidad, &ordenProducto.Total,
			&producto.ID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.Ranking, &producto.Id_categoria, &producto.Categoria,
		); err != nil {
			return nil, fmt.Errorf("error scanning ordenProducto: %w", err)
		}
		ordenProducto.Producto = producto // Asignar el producto a la orden
		ordenesProductos = append(ordenesProductos, ordenProducto)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return ordenesProductos, nil
}
/*func (s *sqlStoreOrdenProductos) BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error) {
	var usuarios []domain.OrdenProducto
	query := "SELECT id, id_orden, id_producto, cantidad, total FROM OrdenProducto"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying usuarios: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usuario domain.OrdenProducto
		if err := rows.Scan(&usuario.ID, &usuario.ID_Orden, &usuario.ID_Producto, &usuario.Cantidad, &usuario.Total); err != nil {
			return nil, fmt.Errorf("error scanning usuario: %w", err)
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return usuarios, nil
}*/


func (s *sqlStoreOrdenProductos) UpdateOrdenProducto(id int, op domain.OrdenProducto) error {
	query := "UPDATE OrdenProducto SET id_orden=?, id_producto=?, cantidad=?, total=? WHERE id=?;"


	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query, op.ID_Orden, op.ID_Producto, op.Cantidad, op.Total, id)
	if err != nil {
		return err 
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("OrdenProducto con ID %d no encontrado", id)
	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> VERIFICA SI EXISTE ORDENPRODUCTO CON ESE ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreOrdenProductos) ExistsByID(id int) bool {
	// Preparar la consulta SQL para verificar si un odontólogo con el ID dado existe
	query := "SELECT COUNT(*) FROM OrdenProducto WHERE id = ?"
	// Ejecutar la consulta SQL y obtener el número de filas devueltas
	var count int
	err := s.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		// Manejar el error, por ejemplo, loguearlo o devolver false si ocurre un error
		return false
	}
	// Si el número de filas devueltas es mayor que cero, significa que el odontólogo con el ID dado existe
	return count > 0
}
