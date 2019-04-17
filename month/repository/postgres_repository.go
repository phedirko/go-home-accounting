package repository

import (
	"database/sql"

	"github.com/phedirko/go-home-accounting/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func setup(db *sql.DB) {
	db.Exec(`CREATE TABLE months(
			id serial PRIMARY KEY, 
			started_on TIMESTAMP NOT NULL,
			balance integer);`)
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) close() { // read about "extension/method?"
	r.db.Close()
}

func (r *PostgresRepository) Insert(month models.Month) error {
	_, err := r.db.Exec("INSERT INTO months(started_on, balance) VALUES($1, $2)", month.StartedAt, month.Balance) // read about underscore
	return err
}

func (r *PostgresRepository) List() ([]models.Month, error) {
	rows, err := r.db.Query("SELECT * FROM months ORDER BY id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	months := []models.Month{}

	for rows.Next() {
		month := models.Month{}
		if err = rows.Scan(&month.ID, &month.StartedAt, &month.Balance); err == nil {
			months = append(months, month)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return months, nil
}
