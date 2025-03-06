package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
	"github.com/vicpoo/ApiPublish/src/core"
)

type MySQLPlatilloRepository struct {
	conn *sql.DB
}

func NewMySQLPlatilloRepository() domain.IPlatillo {
	conn := core.GetDB()
	return &MySQLPlatilloRepository{conn: conn}
}

func (mysql *MySQLPlatilloRepository) Save(platillo entities.Platillo) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO Platillos (nombre, descripcion, precio) VALUES (?, ?, ?)",
		platillo.Nombre,
		platillo.Descripcion,
		platillo.Precio,
	)
	if err != nil {
		log.Println("Error al guardar el platillo:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	platillo.SetID(int(idInserted))
	return nil
}

func (mysql *MySQLPlatilloRepository) Update(id int, platillo entities.Platillo) error {
	result, err := mysql.conn.Exec(
		"UPDATE Platillos SET nombre = ?, descripcion = ?, precio = ? WHERE id = ?",
		platillo.Nombre,
		platillo.Descripcion,
		platillo.Precio,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el platillo:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el platillo con ID:", id)
		return fmt.Errorf("platillo con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLPlatilloRepository) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM Platillos WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el platillo:", err)
		return err
	}
	return nil
}

func (mysql *MySQLPlatilloRepository) FindByID(id int) (entities.Platillo, error) {
	var platillo entities.Platillo
	row := mysql.conn.QueryRow("SELECT id, nombre, descripcion, precio FROM Platillos WHERE id = ?", id)

	err := row.Scan(
		&platillo.ID,
		&platillo.Nombre,
		&platillo.Descripcion,
		&platillo.Precio,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Platillo no encontrado:", err)
			return entities.Platillo{}, fmt.Errorf("platillo con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el platillo por ID:", err)
		return entities.Platillo{}, err
	}

	return platillo, nil
}

func (mysql *MySQLPlatilloRepository) GetAll() ([]entities.Platillo, error) {
	var platillos []entities.Platillo

	rows, err := mysql.conn.Query("SELECT id, nombre, descripcion, precio FROM Platillos")
	if err != nil {
		log.Println("Error al obtener todos los platillos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var platillo entities.Platillo
		err := rows.Scan(
			&platillo.ID,
			&platillo.Nombre,
			&platillo.Descripcion,
			&platillo.Precio,
		)
		if err != nil {
			log.Println("Error al escanear platillo:", err)
			return nil, err
		}
		platillos = append(platillos, platillo)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return platillos, nil
}

func (mysql *MySQLPlatilloRepository) GetByNombre(nombre string) (entities.Platillo, error) {
	var platillo entities.Platillo
	row := mysql.conn.QueryRow("SELECT id, nombre, descripcion, precio FROM Platillos WHERE nombre = ?", nombre)

	err := row.Scan(
		&platillo.ID,
		&platillo.Nombre,
		&platillo.Descripcion,
		&platillo.Precio,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Platillo no encontrado:", err)
			return entities.Platillo{}, fmt.Errorf("platillo con nombre %s no encontrado", nombre)
		}
		log.Println("Error al buscar el platillo por nombre:", err)
		return entities.Platillo{}, err
	}

	return platillo, nil
}
