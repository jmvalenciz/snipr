package repository

import (
	"database/sql"
	"errors"
	"snipr/src/model"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type ISnippetRepository interface {
	GetAll() ([]model.Snippet, error)
	// GetByNmae(name string) ([]model.Snippet, error)
	// GetByTags(tag string)
	// GetByNameAndTags(name string, tag string)
	// Create(model.CreateSnippet)
	Delete(id uint) (bool, error)
	Create(newSnippet model.CreateSnippet) (*model.Snippet, error)
}

type snippetRepository struct {
	conn *sql.DB
}

func NewSnippetRepository(path string) (*snippetRepository, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTablesQuery := `
    CREATE TABLE IF NOT EXISTS snippets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    note TEXT,
    body TEXT NOT NULL,
    tags TEXT,
    format TEXT NOT NULL
    );`
	_, err = conn.Exec(createTablesQuery)
	if err != nil {
		return nil, err
	}

	return &snippetRepository{conn}, nil
}

func sqlRowsToSnippet(rows *sql.Rows) ([]model.Snippet, error) {
	var err error
	var snippets []model.Snippet = make([]model.Snippet, 0)
	for rows.Next() {
		var snippet model.Snippet
		var tags string
		err = rows.Scan(&snippet.Id, &snippet.Name, &snippet.Note, &snippet.Body, &tags, &snippet.Format)
		snippet.Tags = strings.Split(tags, ",")
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}
	return snippets, nil
}

func (sr snippetRepository) GetAll() ([]model.Snippet, error) {
	rows, err := sr.conn.Query("SELECT * FROM snippets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets, err := sqlRowsToSnippet(rows)
	if err != nil {
		return nil, err
	}
	return snippets, nil
}

func (sr snippetRepository) GetByName(name string) ([]model.Snippet, error) {
	if name == "" {
		return nil, errors.New("Name is required an cannot be nil")
	}
	query := "SELECT * FROM snippets WHERE name == ?"

	rows, err := sr.conn.Query(query, name)
	if err != nil {
		return nil, err
	}

	snippets, err := sqlRowsToSnippet(rows)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (sr snippetRepository) GetByTag(tag string) ([]model.Snippet, error) {
	query := "SELECT * FROM snippets WHERE name = ?"

	rows, err := sr.conn.Query(query, tag)
	if err != nil {
		return nil, err
	}

	snippets, err := sqlRowsToSnippet(rows)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (sr snippetRepository) Delete(id uint) (bool, error) {
	rows, err := sr.conn.Exec("DELETE FROM snippets WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := rows.RowsAffected()
	return rowsAffected == 1, nil
}

func (sr snippetRepository) Create(newSnippet model.CreateSnippet) (*model.Snippet, error) {
	rows, err := sr.conn.Exec(
		"INSERT INTO snippets(name, note, body, tags, format) VALUES(?,?,?,?,?);",
		newSnippet.Name,
		newSnippet.Note,
		newSnippet.Body,
		strings.Join(newSnippet.Tags, ","),
		newSnippet.Format,
	)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	snippet := &model.Snippet{
		Id:     uint(id),
		Name:   newSnippet.Name,
		Body:   newSnippet.Body,
		Tags:   newSnippet.Tags,
		Format: newSnippet.Format,
		Note:   newSnippet.Note,
	}
	return snippet, nil
}
