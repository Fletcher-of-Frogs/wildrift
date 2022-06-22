package database

import (
	"database/sql"
	"fmt"
)

type Website struct {
	*sql.DB
}

var DB *sql.DB

func CheckErr(e error) {
	fmt.Println(e)
}

func Hi() {
	fmt.Println("hi")
}

func championData(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS champion (
		name TEXT,
		role TEXT,
		strong TEXT,
		weak TEXT);
	`)
	if err != nil {
		(CheckErr(err))
	}
	stmt.Exec()
}

func Connect(db *sql.DB) *Website {
	championData(db)

	return &Website{
		DB: db,
	}
}

func (website *Website) CreateChampion(champion Champion) {
	stmt, err := website.DB.Prepare("INSERT INTO champion (name, role, strong, weak) VALUES (?,?,?,?);")
	if err != nil {
		CheckErr(err)
	}

	stmt.Exec(champion.Name, champion.Role, champion.StrongVS, champion.WeakVS)
	defer stmt.Close()
}

func (data *Website) GetChamp() []Champion {
	champions := []Champion{}

	rows, _ := data.DB.Query(`SELECT * FROM champion`)

	var name string
	var role string
	var strong string
	var weak string

	for rows.Next() {
		rows.Scan(&name, &role, &strong, &weak)
	}
	champions = append(champions, Champion{
		Name:     name,
		Role:     role,
		StrongVS: strong,
		WeakVS:   weak,
	})
	return champions
}