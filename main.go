package main

import (
    "database/sql"
    "log"
    "Music-database/src/Minero"
    "Music-database/src/Interfaz"
    _ "github.com/mattn/go-sqlite3"
    "path/filepath"
)

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
        rutaCompleta := filepath.Join(rutaBase, song.Title + ".mp3") // Crear ruta completa del archivo

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
    // Abrir la base de datos
    database, err := sql.Open("sqlite3", "/home/grauler/Documentos/Ciencias de la Computacion/3 semestre/Modelado Y Programacion/Music-database/src/Base/Base.db")
    if err != nil {
        log.Fatalf("Error al abrir la base de datos: %v", err)
    }
    defer database.Close()

    // Llamada a la función MinarDirectorio
    songs, err := Minero.MinarDirectorio(ruta)
    if err != nil {
        log.Fatalf("Error al minar el directorio: %v", err)
    }

    // Guardar metadatos en la base de datos
    err = GuardarMetadatosEnBD(database, songs, ruta)
    if err != nil {
        log.Fatalf("Error al guardar metadatos en la base de datos: %v", err)
    }
}

func main() {
    Interfaz.CrearVentana(minarYGuardar)
}
