package repositories

import (
	"api/src/models"
	"database/sql"
)

type VotesRepo struct {
	db *sql.DB
}

func NewVotesRepo(db *sql.DB) *VotesRepo {
	return &VotesRepo{db}
}

func (repo VotesRepo) Create(vote models.Vote) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO cryptoUpvoting.votes (userId, cryptoId) VALUES (?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(vote.UserId, vote.CryptoId)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertId), nil
}

func (repo VotesRepo) FindAll() ([]models.Ranking, error) {
	rows, err := repo.db.Query(
		"SELECT cr.name AS name, COUNT(vo.cryptoId) AS votes" +
			"FROM cryptoUpvoting.votes as vo" +
			"JOIN cryptoUpvoting.cryptos AS cr" +
			"ON vo.cryptoId = cr.id" +
			"GROUP BY vo.cryptoId" +
			"ORDER BY votes DESC, name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []models.Ranking
	for rows.Next() {
		var vote models.Ranking

		if err = rows.Scan(
			&vote.CryptoName, &vote.CryptoVotes,
		); err != nil {
			return nil, err
		}

		votes = append(votes, vote) // votes.push(vote)
	}

	return votes, nil
}

func (repo VotesRepo) FindVotesByUser(id uint64) ([]models.Ranking, error) {
	rows, err := repo.db.Query(
		"SELECT cr.name AS name, COUNT(vo.cryptoId) AS votes"+
			"FROM cryptoUpvoting.votes as vo"+
			"JOIN cryptoUpvoting.cryptos AS cr"+
			"ON vo.cryptoId = cr.id"+
			"JOIN cryptoUpvoting.users AS us"+
			"ON vo.userId = us.id"+
			"WHERE us.id = ?"+
			"GROUP BY vo.cryptoId"+
			"ORDER BY votes DESC, name ASC",
		id,
	)
	if err != nil {
		return []models.Ranking{}, err
	}
	defer rows.Close()

	var votes []models.Ranking
	for rows.Next() {
		var vote models.Ranking

		if err = rows.Scan(
			&vote.CryptoName, &vote.CryptoVotes,
		); err != nil {
			return nil, err
		}

		votes = append(votes, vote) // votes.push(vote)
	}

	return votes, nil
}
