package Minero

import (
	
	"log"
	"os"
	"path/filepath"
	"time"
	"fmt"
	"github.com/dhowden/tag"
)

type SongMetaData struct {
	Artist string
	Title  string
	Album  string
	Year   int
	Genre  string
	Track       string
    TotalTracks int
}

func MinarDirectorio(ruta string) ([]SongMetaData, error) {
	var songs []SongMetaData

	err := filepath.Walk(ruta, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".mp3" {
			file, err := os.Open(path)
			if err != nil {
				log.Printf("Error al abrir el archivo %s: %v", path, err)
				return nil // Continuar con otros archivos
			}
			defer file.Close()

			metadata, err := tag.ReadFrom(file)
			if err != nil {
				log.Printf("Error al leer los metadatos de %s: %v", path, err)
				return nil // Continuar con otros archivos
			}
			track, totalTracks := metadata.Track()
			// Crear instancia de SongMetaData con valores predeterminados si faltan etiquetas

			if totalTracks > 0 && track > totalTracks {
                log.Printf("Número de pista inválido en %s: pista %d, total de pistas %d. Asignando valor predeterminado.", path, track, totalTracks)
                track = 1 // Asignar un valor predeterminado
            }
			
			song := SongMetaData{
				Artist: metadata.Artist(),
				Title:  metadata.Title(),
				Album:  metadata.Album(),
				Year:   metadata.Year(),
				Genre:  metadata.Genre(),
				Track:       fmt.Sprintf("%d/%d", track, totalTracks),
                TotalTracks: totalTracks,
			}

			// Asignar valores por defecto si los campos están vacíos
			if song.Artist == "" {
				song.Artist = "Unknown"
			}
			if song.Title == "" {
				song.Title = "Unknown"
			}
			if song.Album == "" {
				song.Album = filepath.Base(filepath.Dir(path)) // Usar nombre del directorio si no hay álbum
			}
			if song.Year == 0 {
				song.Year = time.Now().Year()
			}
			if song.Genre == "" {
				song.Genre = "Unknown"
			}
			if song.Track == "0/0" {
				song.Track = "1/1" // Valor por defecto para el número de pista
			}

			songs = append(songs, song)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return songs, nil
}
