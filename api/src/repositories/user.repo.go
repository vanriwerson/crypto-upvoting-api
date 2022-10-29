package repos

import (
	"api/src/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (repo UsersRepo) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO cryptoUpvoting.users (name, email, password) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertId), nil
}

func (repo UsersRepo) FindAll() ([]models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, name, email, createdAt FROM cryptoUpvoting.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user) // users.push(user)
	}

	return users, nil
}

func (repo UsersRepo) FindOne(id uint64) (models.User, error) {
	row, err := repo.db.Query(
		"SELECT id, name, email, createdAt FROM cryptoUpvoting.users WHERE id = ?",
		id,
	)
	if err != nil {
		return models.User{}, err // passando um usuário vazio para o retorno em caso de erro
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID, &user.Name, &user.Email, &user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo UsersRepo) Update(id uint64, user models.User) error {
	statement, erro := repo.db.Prepare(
		"UPDATE cryptoUpvoting.users SET name = ?, email = ?, password = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Email, user.Password, id); erro != nil {
		return erro
	}

	return nil
}

func (repo UsersRepo) Delete(id uint64) error {
	statement, err := repo.db.Prepare(
		"DELETE FROM cryptoUpvoting.usuarios WHERE id = ?",
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
