package main

import (
	"path/filepath"
	"database/sql"
	"fmt"
	"log"
	"Music-database/src/Minero"
		_ "github.com/mattn/go-sqlite3"
)

func GuardarMetadatosEnBD(database *sql.DB, songs []Minero.SongMetaData, rutaBase string) error {
    stmt, err := database.Prepare(`INSERT INTO rolas (id_performer, id_album, path, title, track, year, genre) VALUES (?, ?, ?, ?, ?, ?, ?)`)
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, song := range songs {
        rutaCompleta := filepath.Join(rutaBase, song.Title + ".mp3") // Crear ruta completa del archivo
        _, err = stmt.Exec(1, 1, rutaCompleta, song.Title, song.Track, song.Year, song.Genre) // Suponiendo id_performer e id_album = 1
        if err != nil {
            log.Printf("Error al insertar la canción %s: %v", song.Title, err)
            continue
        }
    }

    return nil
}

func main() {
    // Ruta del directorio que contiene los archivos MP3
    var ruta string

    // Solicitar al usuario que ingrese la ruta del directorio
    fmt.Print("Introduce la ruta del directorio que contiene los archivos MP3: ")
    _, err := fmt.Scanf("%s", &ruta)
    if err != nil {
        log.Fatalf("Error al leer la ruta: %v", err)
    }

    // Abrir la base de datos
    database, err := sql.Open("sqlite3", "/home/grauler/Vídeos/Base.db")
    if err != nil {
        panic(err)
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

    // Mostrar la metadata extraída
    for _, song := range songs {
        fmt.Printf("Título: %s\nArtista: %s\nÁlbum: %s\nAño: %d\nGénero: %s\nTrack: %d\n\n", 
            song.Title, song.Artist, song.Album, song.Year, song.Genre, song.Track)
    }

    log.Println("Metadatos guardados exitosamente")
}
