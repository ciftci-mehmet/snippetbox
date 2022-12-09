package mysql

import (
	"database/sql"
	"errors"

	"github.com/ciftci-mehmet/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct.
	s := &models.Snippet{}

	// copy values from each field  in sql to corresponding field in snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// no rows returned
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// if ok return snippet obj
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// !!!ALWAYS defer close sql.Rows resultset
	defer rows.Close()

	// init empty slice to hold models.Snippets objects
	snippets := []*models.Snippet{}

	// loop each row
	for rows.Next() {
		s := &models.Snippet{}

		// copy values from fields to snippet object
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// append row to snippets slice
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
