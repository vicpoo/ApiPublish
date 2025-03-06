package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
	"github.com/vicpoo/ApiPublish/src/core"
)

type MySQLOrdenRepository struct {
	conn *sql.DB
}

func NewMySQLOrdenRepository() domain.IOrden {
	conn := core.GetDB()
	return &MySQLOrdenRepository{conn: conn}
}

func (mysql *MySQLOrdenRepository) Save(orden entities.Orden) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO Ordenes (mesa_id, estado) VALUES (?, ?)",
		orden.MesaID,
		orden.Estado,
	)
	if err != nil {
		log.Println("Error al guardar la orden:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	orden.SetID(int(idInserted))
	return nil
}

func (mysql *MySQLOrdenRepository) Update(id int, orden entities.Orden) error {
	result, err := mysql.conn.Exec(
		"UPDATE Ordenes SET mesa_id = ?, estado = ? WHERE id = ?",
		orden.MesaID,
		orden.Estado,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar la orden:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró la orden con ID:", id)
		return fmt.Errorf("orden con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLOrdenRepository) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM Ordenes WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar la orden:", err)
		return err
	}
	return nil
}

func (mysql *MySQLOrdenRepository) FindByID(id int) (entities.Orden, error) {
	var orden entities.Orden
	row := mysql.conn.QueryRow("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes WHERE id = ?", id)

	var fechaCreacion string
	err := row.Scan(
		&orden.ID,
		&orden.MesaID,
		&orden.Estado,
		&fechaCreacion,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Orden no encontrada:", err)
			return entities.Orden{}, fmt.Errorf("orden con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la orden por ID:", err)
		return entities.Orden{}, err
	}

	orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)
	return orden, nil
}

func (mysql *MySQLOrdenRepository) GetAll() ([]entities.Orden, error) {
	var ordenes []entities.Orden

	rows, err := mysql.conn.Query("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes")
	if err != nil {
		log.Println("Error al obtener todas las órdenes:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orden entities.Orden
		var fechaCreacion string
		err := rows.Scan(
			&orden.ID,
			&orden.MesaID,
			&orden.Estado,
			&fechaCreacion,
		)
		if err != nil {
			log.Println("Error al escanear orden:", err)
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)
		ordenes = append(ordenes, orden)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return ordenes, nil
}

func (mysql *MySQLOrdenRepository) GetByMesaID(mesaID int) ([]entities.Orden, error) {
	var ordenes []entities.Orden

	rows, err := mysql.conn.Query("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes WHERE mesa_id = ?", mesaID)
	if err != nil {
		log.Println("Error al obtener las órdenes por mesa ID:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orden entities.Orden
		var fechaCreacion string
		err := rows.Scan(
			&orden.ID,
			&orden.MesaID,
			&orden.Estado,
			&fechaCreacion,
		)
		if err != nil {
			log.Println("Error al escanear orden:", err)
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)
		ordenes = append(ordenes, orden)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return ordenes, nil
}
