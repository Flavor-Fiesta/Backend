package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"log"
)

type sqlStoreUsuarios struct {
	db *sql.DB
}

func NewSqlStoreUsuarios(db *sql.DB) StoreInterfaceUsuarios {
	return &sqlStoreUsuarios{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UNA NUEVA USUARIO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreUsuarios) CrearUsuario(usuario domain.Usuarios) error {
    query := "INSERT INTO usuarios (nombre, email, telefono, password, id_rol) VALUES (?, ?, ?, ?, ?);"
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error preparing query: %w", err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(usuario.Nombre, usuario.Email, usuario.Telefono, usuario.Password, usuario.Id_rol)
    if err != nil {
        return fmt.Errorf("error executing query: %w", err)
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return fmt.Errorf("error fetching rows affected: %w", err)
    }

    if rowsAffected != 1 {
        return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
    }

    return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR USUARIO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreUsuarios) BuscarUsuario(id int) (domain.Usuarios, error) {
	var usuario domain.Usuarios
	query := "SELECT id, nombre, email, telefono, password FROM usuarios WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Telefono, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Usuarios{}, errors.New("usuario not found")
		}
		return domain.Usuarios{}, err
	}

	return usuario, nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR USUARIO POR EMAIL Y CLAVE <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreUsuarios) BuscarUsuarioPorEmailYPassword(email, password string) (bool, error) {
	var usuario domain.Usuarios
	query := "SELECT id FROM usuarios WHERE email = ? AND password = ?"

	err := s.db.QueryRow(query, email, password).Scan(&usuario.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
//ESTE TRAE TODOS LOS DATOS

func (s *sqlStoreUsuarios) BuscarUsuarioPorEmailYPassword2(email, password string) (domain.Usuarios, error) {
	var usuario domain.Usuarios
	query := "SELECT id, nombre, email, telefono, password FROM usuarios WHERE email = ? AND password = ?"

	err := s.db.QueryRow(query, email, password).Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Telefono, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Usuarios{}, errors.New("usuario not found")
		}
		return domain.Usuarios{}, err
	}

	return usuario, nil
}

func (s *sqlStoreUsuarios) BuscarUsuarioPorEmailYPassword3(email, password string) (bool, error, domain.Usuarios){
	var usuario domain.Usuarios
	query := "SELECT * FROM usuarios WHERE email = ? AND password = ?"

	err := s.db.QueryRow(query, email, password).Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Telefono, &usuario.Password)
	if err != nil {
        if err == sql.ErrNoRows {
            return false, err, domain.Usuarios{}
        }
        return false, err, domain.Usuarios{}
    }
    return true, nil, usuario
}

func (s *sqlStoreUsuarios) BuscarTodosLosUsuarios() ([]domain.Usuarios, error) {
	var usuarios []domain.Usuarios
	query := "SELECT id, nombre, email, telefono, password FROM usuarios"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying usuarios: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usuario domain.Usuarios
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Telefono, &usuario.Password); err != nil {
			return nil, fmt.Errorf("error scanning usuario: %w", err)
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return usuarios, nil
}



/*func (s *sqlStoreUsuarios) BuscarProductoPorID(id int) (domain.Producto, error) {
    // Preparar la consulta SQL para buscar un producto por su ID
    query := "SELECT id, nombre, codigo, categoria, fecha_alta, fecha_vencimiento FROM productos WHERE id = ?"
    
    // Ejecutar la consulta SQL y obtener el resultado
    var producto domain.Producto
    err := s.db.QueryRow(query, id).Scan(&producto.ID, &producto.Nombre, &producto.Codigo, &producto.Categoria, &producto.FechaDeAlta, &producto.FechaDeVencimiento)
    if err != nil {
        // Manejar el error, por ejemplo, devolver un error específico si no se encuentra el producto
        if err == sql.ErrNoRows {
            return domain.Producto{}, fmt.Errorf("producto con ID %d no encontrado", id)
        }
        return domain.Producto{}, err
    }

    // Devolver el producto encontrado
    return producto, nil
}
*/



// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ELIMINAR UNA IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreUsuarios) DeleteUsuario(id int) error {
	query := "DELETE FROM usuarios WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func (s *sqlStoreUsuarios) ExistsByIDUsuario(id int) (bool, error) {
    // Consulta SQL para buscar un producto por su ID
    query := "SELECT id FROM usuarios WHERE id = ?"

    // Ejecutar la consulta SQL y escanear el resultado en una variable id
    var count int
    err := s.db.QueryRow(query, id).Scan(&count)
    if err != nil {
        // Si se produce un error, verificamos si se trata de un error de "ninguna fila encontrada"
        if err == sql.ErrNoRows {
            // No se encontró ninguna fila, por lo que el producto no existe
            return false, nil
			fmt.Print(" REVISANDO EN SQL STORE:  ", query, count )
        }
        // Otro tipo de error, devolver el error
        return false, err
    }

    // Si se encontró un producto con el ID dado, devolver true
    return true, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> VALIDACIONES >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *sqlStoreUsuarios) ExisteEmail(email string) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM usuarios WHERE email = ?)"
    err := s.db.QueryRow(query, email).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func (s *sqlStoreUsuarios) ExisteCelular(celular string) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM usuarios WHERE telefono = ?)"
    err := s.db.QueryRow(query, celular).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}






// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UNA IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreUsuarios) Update(p domain.Usuarios) error {

    // Preparar la consulta SQL para actualizar el odontólogo
    query := "UPDATE usuarios SET nombre = ?, email = ?, telefono = ?, password = ? WHERE id = ?;"

    // Ejecutar la consulta SQL
    result, err := s.db.Exec(query, p.Nombre, p.Email, p.Telefono, p.Password, p.ID)
    if err != nil {
        return err // Devolver el error si ocurre alguno al ejecutar la consulta
    }

    // Verificar si se actualizó algún registro
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    // Si no se actualizó ningún registro, significa que el odontólogo con el ID dado no existe
    if rowsAffected == 0 {
        return fmt.Errorf("Usuario con ID %d no encontrado", p.ID)
    }

    // Si todo fue exitoso, retornar nil
    return nil
}




//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *sqlStoreUsuarios) PatchUsuario(id int, updatedFields map[string]interface{}) error {
    // Comprobar si se proporcionan campos para actualizar
    if len(updatedFields) == 0 {
        return errors.New("no fields provided for patching")
    }

    // Construir la consulta SQL para actualizar los campos
    query := "UPDATE usuarios SET"
    values := make([]interface{}, 0)
    index := 0
    for field, value := range updatedFields {
        query += fmt.Sprintf(" %s = ?", field)
        values = append(values, value)
        index++
        if index < len(updatedFields) {
            query += ","
        }
    }
    query += " WHERE id = ?"
    values = append(values, id)

    // Preparar y ejecutar la consulta SQL
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
}