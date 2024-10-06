package Base

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./music.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func CrearTablas(db *sql.DB) error {
    esquemaSQL := `
    CREATE TABLE IF NOT EXISTS types (
        id_type       INTEGER PRIMARY KEY,
        description   TEXT
    );

    CREATE TABLE IF NOT EXISTS performers (
        id_performer  INTEGER PRIMARY KEY,
        id_type       INTEGER,
        name          TEXT,
        FOREIGN KEY   (id_type) REFERENCES types(id_type)
    );

    CREATE TABLE IF NOT EXISTS persons (
        id_person     INTEGER PRIMARY KEY,
        stage_name    TEXT,
        real_name     TEXT,
        birth_date    TEXT,
        death_date    TEXT
    );

    CREATE TABLE IF NOT EXISTS groups (
        id_group      INTEGER PRIMARY KEY,
        name          TEXT,
        start_date    TEXT,
        end_date      TEXT
    );

    CREATE TABLE IF NOT EXISTS in_group (
        id_person     INTEGER,
        id_group      INTEGER,
        PRIMARY KEY   (id_person, id_group),
        FOREIGN KEY   (id_person) REFERENCES persons(id_person),
        FOREIGN KEY   (id_group) REFERENCES groups(id_group)
    );

    CREATE TABLE IF NOT EXISTS albums (
        id_album      INTEGER PRIMARY KEY,
        path          TEXT,
        name          TEXT,
        year          INTEGER
    );

    CREATE TABLE IF NOT EXISTS rolas (
        id_rola       INTEGER PRIMARY KEY,
        id_performer  INTEGER,
        id_album      INTEGER,
        path          TEXT,
        title         TEXT,
        track         INTEGER,
        year          INTEGER,
        genre         TEXT,
        FOREIGN KEY   (id_performer) REFERENCES performers(id_performer),
        FOREIGN KEY   (id_album) REFERENCES albums(id_album)
    );
    `

    _, err := db.Exec(esquemaSQL)
    return err
}


