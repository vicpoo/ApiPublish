// mysql_orden_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	detallesOrdenEntities "github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities" // Importación corregida
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
	tx, err := mysql.conn.Begin()
	if err != nil {
		log.Println("Error al iniciar la transacción:", err)
		return err
	}

	// Insertar la orden
	result, err := tx.Exec(
		"INSERT INTO Ordenes (mesa_id, estado) VALUES (?, ?)",
		orden.MesaID,
		orden.Estado,
	)
	if err != nil {
		tx.Rollback()
		log.Println("Error al insertar la orden:", err)
		return err
	}

	ordenID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println("Error al obtener el ID de la orden insertada:", err)
		return err
	}

	// Insertar los detalles de la orden
	for _, detalle := range orden.Detalles {
		_, err := tx.Exec(
			"INSERT INTO DetallesOrden (orden_id, platillo_id, cantidad) VALUES (?, ?, ?)",
			ordenID,
			detalle.PlatilloID,
			detalle.Cantidad,
		)
		if err != nil {
			tx.Rollback()
			log.Println("Error al insertar el detalle de la orden:", err)
			return err
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		log.Println("Error al confirmar la transacción:", err)
		return err
	}

	return nil
}

func (mysql *MySQLOrdenRepository) Update(id int, orden entities.Orden) error {
	tx, err := mysql.conn.Begin()
	if err != nil {
		log.Println("Error al iniciar la transacción:", err)
		return err
	}

	// Actualizar la orden
	result, err := tx.Exec(
		"UPDATE Ordenes SET mesa_id = ?, estado = ? WHERE id = ?",
		orden.MesaID,
		orden.Estado,
		id,
	)
	if err != nil {
		tx.Rollback()
		log.Println("Error al actualizar la orden:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		log.Println("No se encontró la orden con ID:", id)
		return fmt.Errorf("orden con ID %d no encontrada", id)
	}

	// Eliminar los detalles antiguos de la orden
	_, err = tx.Exec("DELETE FROM DetallesOrden WHERE orden_id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Println("Error al eliminar los detalles antiguos de la orden:", err)
		return err
	}

	// Insertar los nuevos detalles de la orden
	for _, detalle := range orden.Detalles {
		_, err := tx.Exec(
			"INSERT INTO DetallesOrden (orden_id, platillo_id, cantidad) VALUES (?, ?, ?)",
			id,
			detalle.PlatilloID,
			detalle.Cantidad,
		)
		if err != nil {
			tx.Rollback()
			log.Println("Error al insertar el nuevo detalle de la orden:", err)
			return err
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		log.Println("Error al confirmar la transacción:", err)
		return err
	}

	return nil
}

func (mysql *MySQLOrdenRepository) Delete(id int) error {
	tx, err := mysql.conn.Begin()
	if err != nil {
		log.Println("Error al iniciar la transacción:", err)
		return err
	}

	// Eliminar los detalles de la orden
	_, err = tx.Exec("DELETE FROM DetallesOrden WHERE orden_id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Println("Error al eliminar los detalles de la orden:", err)
		return err
	}

	// Eliminar la orden
	_, err = tx.Exec("DELETE FROM Ordenes WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Println("Error al eliminar la orden:", err)
		return err
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		log.Println("Error al confirmar la transacción:", err)
		return err
	}

	return nil
}

func (mysql *MySQLOrdenRepository) FindByID(id int) (entities.Orden, error) {
	var orden entities.Orden
	var fechaCreacion string

	// Obtener la orden
	row := mysql.conn.QueryRow("SELECT id, mesa_id, estado, fecha_creacion FROM Ordenes WHERE id = ?", id)
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

	// Parsear la fecha de creación
	orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)

	// Obtener los detalles de la orden
	rows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", id)
	if err != nil {
		log.Println("Error al obtener los detalles de la orden:", err)
		return entities.Orden{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var detalle detallesOrdenEntities.DetallesOrden // Usar el paquete correcto
		err := rows.Scan(
			&detalle.ID,
			&detalle.OrdenID,
			&detalle.PlatilloID,
			&detalle.Cantidad,
		)
		if err != nil {
			log.Println("Error al escanear el detalle de la orden:", err)
			return entities.Orden{}, err
		}
		orden.Detalles = append(orden.Detalles, detalle)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return entities.Orden{}, err
	}

	return orden, nil
}

func (mysql *MySQLOrdenRepository) GetAll() ([]entities.Orden, error) {
	var ordenes []entities.Orden

	// Obtener todas las órdenes
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
			log.Println("Error al escanear la orden:", err)
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)

		// Obtener los detalles de la orden
		detallesRows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", orden.ID)
		if err != nil {
			log.Println("Error al obtener los detalles de la orden:", err)
			return nil, err
		}
		defer detallesRows.Close()

		for detallesRows.Next() {
			var detalle detallesOrdenEntities.DetallesOrden // Usar el paquete correcto
			err := detallesRows.Scan(
				&detalle.ID,
				&detalle.OrdenID,
				&detalle.PlatilloID,
				&detalle.Cantidad,
			)
			if err != nil {
				log.Println("Error al escanear el detalle de la orden:", err)
				return nil, err
			}
			orden.Detalles = append(orden.Detalles, detalle)
		}

		if err := detallesRows.Err(); err != nil {
			log.Println("Error después de iterar filas de detalles:", err)
			return nil, err
		}

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

	// Obtener las órdenes por mesa ID
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
			log.Println("Error al escanear la orden:", err)
			return nil, err
		}
		orden.FechaCreacion, _ = time.Parse("2006-01-02 15:04:05", fechaCreacion)

		// Obtener los detalles de la orden
		detallesRows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", orden.ID)
		if err != nil {
			log.Println("Error al obtener los detalles de la orden:", err)
			return nil, err
		}
		defer detallesRows.Close()

		for detallesRows.Next() {
			var detalle detallesOrdenEntities.DetallesOrden // Usar el paquete correcto
			err := detallesRows.Scan(
				&detalle.ID,
				&detalle.OrdenID,
				&detalle.PlatilloID,
				&detalle.Cantidad,
			)
			if err != nil {
				log.Println("Error al escanear el detalle de la orden:", err)
				return nil, err
			}
			orden.Detalles = append(orden.Detalles, detalle)
		}

		if err := detallesRows.Err(); err != nil {
			log.Println("Error después de iterar filas de detalles:", err)
			return nil, err
		}

		ordenes = append(ordenes, orden)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return ordenes, nil
}
