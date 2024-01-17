package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Completed   bool   `json:"completed"`
	CompletedAt string `json:"completed_at"`
}

func GetTodos() ([]Todo, error) {

	rows, err := DB.Query("SELECT * FROM todo_list")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := make([]Todo, 0)

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.CompletedAt)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodo(todo Todo) error {
	_, err := DB.Exec("INSERT INTO todo_list (name, completed, completed_at) VALUES (?, ?, ?)", todo.Name, todo.Completed, todo.CompletedAt)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoById(id string) error {
	_, err := DB.Exec("SELECT * FROM todo_list WHERE id=?", id)
	if err != nil {
		return err
	}
	_, err = DB.Exec("DELETE FROM todo_list WHERE id=?", id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateTodoById(id string, todo Todo) error {

	_, err := DB.Exec("SELECT * FROM todo_list WHERE id=?", id)
	if err != nil {
		return err
	}

	_, err = DB.Exec("UPDATE todo_list SET name=?, completed=?, completed_at=? WHERE id=?", todo.Name, todo.Completed, todo.CompletedAt, id)

	if err != nil {
		return err
	}

	return nil
}
