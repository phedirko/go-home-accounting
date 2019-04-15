package repository

import (
	"database/sql"
	"home-accounting/models"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func init()  {
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Printf("init error:", err)
	}

}

func NewPostgres(url string) (*PostgresRepository, error)  {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close()  { // read about "extension"
	r.db.Close()
}

func (r *PostgresRepository) Insert(month models.Month) error {
	_, err := r.db.Exec("INSERT INTO months(startedAt, balance) VALUES($1, $2)", month.StartedAt, month.Balance) // read about underscore
	return err
}

func (r *PostgresRepository) List()  ([]models.Month, error) {
	rows, err := r.db.Query("SELECT * FROM months ORDER BY id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	months := []models.Month{}

	for rows.Next() {
		month := models.Month{}
		if err = rows.Scan(&month.ID, &month.StartedAt, &month.Balance, &month.FinishedAt); err == nil {
			months = append(months, month)
		}
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return  months, nil
}

