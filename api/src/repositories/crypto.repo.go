package repositories

import (
	"api/src/models"
	"database/sql"
)

type CryptosRepo struct {
	db *sql.DB
}

func NewCryptosRepo(db *sql.DB) *CryptosRepo {
	return &CryptosRepo{db}
}

func (repo CryptosRepo) Create(crypto models.Crypto) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO cryptoUpvoting.cryptos (name) VALUES (?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(crypto.Name)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertId), nil
}

func (repo CryptosRepo) FindAll() ([]models.Crypto, error) {
	rows, err := repo.db.Query(
		"SELECT * FROM cryptoUpvoting.cryptos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []models.Crypto
	for rows.Next() {
		var crypto models.Crypto

		if err = rows.Scan(
			&crypto.ID, &crypto.Name,
		); err != nil {
			return nil, err
		}

		cryptos = append(cryptos, crypto) // cryptos.push(crypto)
	}

	return cryptos, nil
}

func (repo CryptosRepo) FindOne(id uint64) (models.Crypto, error) {
	row, err := repo.db.Query(
		"SELECT id, name, email, createdAt FROM cryptoUpvoting.users WHERE id = ?",
		id,
	)
	if err != nil {
		return models.Crypto{}, err
	}
	defer row.Close()

	var crypto models.Crypto

	if row.Next() {
		if err = row.Scan(
			&crypto.ID, &crypto.Name,
		); err != nil {
			return models.Crypto{}, err
		}
	}

	return crypto, nil
}

func (repo CryptosRepo) Update(id uint64, crypto models.Crypto) error {
	statement, err := repo.db.Prepare(
		"UPDATE cryptoUpvoting.cryptos SET name = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(crypto.Name, id); err != nil {
		return err
	}

	return nil
}

func (repo CryptosRepo) Delete(id uint64) error {
	statement, err := repo.db.Prepare(
		"DELETE FROM cryptoUpvoting.cryptos WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}
