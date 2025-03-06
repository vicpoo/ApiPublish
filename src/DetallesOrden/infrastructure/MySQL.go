package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
	"github.com/vicpoo/ApiPublish/src/core"
)

type MysqlDetallesOrden struct {
	conn *sql.DB
}

func NewMysqlDetallesOrdenRepository() domain.IDetallesOrden {
	conn := core.GetDB()
	return &MysqlDetallesOrden{conn: conn}
}

func (mysql *MysqlDetallesOrden) Save(detalle entities.DetallesOrden) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO DetallesOrden (orden_id, platillo_id, cantidad) VALUES (?, ?, ?)",
		detalle.OrdenID,
		detalle.PlatilloID,
		detalle.Cantidad,
	)
	if err != nil {
		log.Println("Error al guardar el detalle de orden:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	detalle.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlDetallesOrden) Update(id int, detalle entities.DetallesOrden) error {
	result, err := mysql.conn.Exec(
		"UPDATE DetallesOrden SET orden_id = ?, platillo_id = ?, cantidad = ? WHERE id = ?",
		detalle.OrdenID,
		detalle.PlatilloID,
		detalle.Cantidad,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el detalle de orden:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el detalle de orden con ID:", id)
		return fmt.Errorf("detalle de orden con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlDetallesOrden) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM DetallesOrden WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el detalle de orden:", err)
		return err
	}
	return nil
}

func (mysql *MysqlDetallesOrden) FindByID(id int) (entities.DetallesOrden, error) {
	var detalle entities.DetallesOrden
	row := mysql.conn.QueryRow("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE id = ?", id)

	err := row.Scan(
		&detalle.ID,
		&detalle.OrdenID,
		&detalle.PlatilloID,
		&detalle.Cantidad,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Detalle de orden no encontrado:", err)
			return entities.DetallesOrden{}, fmt.Errorf("detalle de orden con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el detalle de orden por ID:", err)
		return entities.DetallesOrden{}, err
	}

	return detalle, nil
}

func (mysql *MysqlDetallesOrden) GetAll() ([]entities.DetallesOrden, error) {
	var detalles []entities.DetallesOrden

	rows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden")
	if err != nil {
		log.Println("Error al obtener todos los detalles de orden:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var detalle entities.DetallesOrden
		err := rows.Scan(
			&detalle.ID,
			&detalle.OrdenID,
			&detalle.PlatilloID,
			&detalle.Cantidad,
		)
		if err != nil {
			log.Println("Error al escanear detalle de orden:", err)
			return nil, err
		}
		detalles = append(detalles, detalle)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return detalles, nil
}

func (mysql *MysqlDetallesOrden) GetByOrdenID(ordenID int) ([]entities.DetallesOrden, error) {
	var detalles []entities.DetallesOrden

	rows, err := mysql.conn.Query("SELECT id, orden_id, platillo_id, cantidad FROM DetallesOrden WHERE orden_id = ?", ordenID)
	if err != nil {
		log.Println("Error al obtener detalles de orden por ID de orden:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var detalle entities.DetallesOrden
		err := rows.Scan(
			&detalle.ID,
			&detalle.OrdenID,
			&detalle.PlatilloID,
			&detalle.Cantidad,
		)
		if err != nil {
			log.Println("Error al escanear detalle de orden:", err)
			return nil, err
		}
		detalles = append(detalles, detalle)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return detalles, nil
}
