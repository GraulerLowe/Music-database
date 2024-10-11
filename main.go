package main

import (
	"os"
    "database/sql"
    "log"
    "Music-database/src/Minero"
    "Music-database/src/Interfaz"
    "Music-database/src/Base"
    _ "github.com/mattn/go-sqlite3"
    "path/filepath"
    "fmt"
)

func dirHOME() string {
    dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(dirname)

    dbPath := filepath.Join(dirname, ".local", "share", "Music-database")

    err = os.MkdirAll(dbPath, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }

    return dbPath
}

func GuardarMetadatosEnBD(database *sql.DB, songs []Minero.SongMetaData, rutaBase string) error {
    stmtSelect, err := database.Prepare(`SELECT COUNT(*) FROM rolas WHERE path = ?`)
    if err != nil {
        return err
    }
    defer stmtSelect.Close()

    stmtInsert, err := database.Prepare(`INSERT INTO rolas (id_performer, id_album, path, title, track, year, genre) VALUES (?, ?, ?, ?, ?, ?, ?)`)
    if err != nil {
        return err
    }
    defer stmtInsert.Close()

    for _, song := range songs {
        rutaCompleta := filepath.Join(rutaBase, song.Title + ".mp3") 

        var count int
        err = stmtSelect.QueryRow(rutaCompleta).Scan(&count)
        if err != nil {
            return err
        }

        if count > 0 {
            log.Printf("Canción ya existente: %s, saltando inserción.", song.Title)
            continue
        }


        _, err = stmtInsert.Exec(1, 1, rutaCompleta, song.Title, song.Track, song.Year, song.Genre) 
        if err != nil {
            log.Printf("Error al insertar la canción %s: %v", song.Title, err)
            continue
        }
    }

    return nil
}

func minarYGuardar(ruta string) {

	dbPath := dirHOME()
    dbFilePath := filepath.Join(dbPath, "Base.db")
	
    // Abrir la base de datos
    database, err := sql.Open("sqlite3", dbFilePath)
    if err != nil {
        log.Printf("Error al abrir la base de datos: %v", err)
       
        return
    }
    defer database.Close()

    err = Base.CrearTablas(database)
    if err != nil {
        log.Printf("Error al crear la base de datos: %v", err)
        
        return
    }

    // Llamada a la función MinarDirectorio
    songs, err := Minero.MinarDirectorio(ruta)
    if err != nil {
        log.Printf("Error al minar el directorio: %v", err)
        
        return
    }

    // Guardar metadatos en la base de datos
    err = GuardarMetadatosEnBD(database, songs, ruta)
    if err != nil {
        log.Printf("Error al guardar metadatos en la base de datos: %v", err)
        
        return
    }
   
}

func listarCanciones() ([]string, error) {

	dbPath := dirHOME()
    dbFilePath := filepath.Join(dbPath, "Base.db")
	
    database, err := sql.Open("sqlite3", dbFilePath)
    if err != nil {
        return nil, err
    }
    defer database.Close()

    err = Base.CrearTablas(database)
    if err != nil {
        return nil, err
    }

    return ListarCanciones(database)
}

func ListarCanciones(database *sql.DB) ([]string, error) {
    rows, err := database.Query(`SELECT title, track, year, genre FROM rolas`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var canciones []string
    for rows.Next() {
        var title, genre, track string
        var year int
        err = rows.Scan(&title, &track, &year, &genre)
        if err != nil {
            return nil, err
        }
        canciones = append(canciones, fmt.Sprintf("Título: %s, Pista: %s, Año: %d, Género: %s", title, track, year, genre))
    }

    return canciones, nil
}

// Nueva función para buscar canciones
func BuscarCanciones(criterio string) ([]string, error) {
    // Abrir conexión a la base de datos
	dbPath := dirHOME()
    dbFilePath := filepath.Join(dbPath, "Base.db")
	
    database, err := sql.Open("sqlite3", dbFilePath)
    if err != nil {
        return nil, err
    }
    defer database.Close()

    // Modificar la consulta para que solo busque en el campo "title"
    rows, err := database.Query(`SELECT title, track, year, genre FROM rolas WHERE title LIKE ?`, "%"+criterio+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var canciones []string
    for rows.Next() {
        var title, genre, track string
        var year int
        
        err = rows.Scan(&title, &track, &year, &genre)
        if err != nil {
            return nil, err
        }
        // Formatear los datos y agregarlos al slice de canciones
        canciones = append(canciones, fmt.Sprintf("Título: %s, Pista: %s, Año: %d, Género: %s", title, track, year, genre))
    }

    return canciones, nil
}



func main() {
    Interfaz.CrearVentana(minarYGuardar, listarCanciones, BuscarCanciones)
}
