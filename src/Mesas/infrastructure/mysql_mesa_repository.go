package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPublish/src/Mesas/domain"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
	"github.com/vicpoo/ApiPublish/src/core"
)

type MySQLMesaRepository struct {
	conn *sql.DB
}

func NewMySQLMesaRepository() domain.IMesa {
	conn := core.GetDB()
	return &MySQLMesaRepository{conn: conn}
}

func (mysql *MySQLMesaRepository) Save(mesa entities.Mesa) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO Mesas (numero_mesa) VALUES (?)",
		mesa.NumeroMesa,
	)
	if err != nil {
		log.Println("Error al guardar la mesa:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	mesa.SetID(int(idInserted))
	return nil
}

func (mysql *MySQLMesaRepository) Update(id int, mesa entities.Mesa) error {
	result, err := mysql.conn.Exec(
		"UPDATE Mesas SET numero_mesa = ? WHERE id = ?",
		mesa.NumeroMesa,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar la mesa:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró la mesa con ID:", id)
		return fmt.Errorf("mesa con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLMesaRepository) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM Mesas WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar la mesa:", err)
		return err
	}
	return nil
}

func (mysql *MySQLMesaRepository) FindByID(id int) (entities.Mesa, error) {
	var mesa entities.Mesa
	row := mysql.conn.QueryRow("SELECT id, numero_mesa FROM Mesas WHERE id = ?", id)

	err := row.Scan(
		&mesa.ID,
		&mesa.NumeroMesa,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Mesa no encontrada:", err)
			return entities.Mesa{}, fmt.Errorf("mesa con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la mesa por ID:", err)
		return entities.Mesa{}, err
	}

	return mesa, nil
}

func (mysql *MySQLMesaRepository) GetAll() ([]entities.Mesa, error) {
	var mesas []entities.Mesa

	rows, err := mysql.conn.Query("SELECT id, numero_mesa FROM Mesas")
	if err != nil {
		log.Println("Error al obtener todas las mesas:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mesa entities.Mesa
		err := rows.Scan(
			&mesa.ID,
			&mesa.NumeroMesa,
		)
		if err != nil {
			log.Println("Error al escanear mesa:", err)
			return nil, err
		}
		mesas = append(mesas, mesa)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mesas, nil
}

// Método GetByNumeroMesa para cumplir con la interfaz IMesa
func (mysql *MySQLMesaRepository) GetByNumeroMesa(numeroMesa int) (entities.Mesa, error) {
	var mesa entities.Mesa
	row := mysql.conn.QueryRow("SELECT id, numero_mesa FROM Mesas WHERE numero_mesa = ?", numeroMesa)

	err := row.Scan(
		&mesa.ID,
		&mesa.NumeroMesa,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Mesa no encontrada:", err)
			return entities.Mesa{}, fmt.Errorf("mesa con número %d no encontrada", numeroMesa)
		}
		log.Println("Error al buscar la mesa por número:", err)
		return entities.Mesa{}, err
	}

	return mesa, nil
}
