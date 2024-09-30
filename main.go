package main

import (
	"fmt"
	"log"
	"Music-database/src/Minero" // Asegúrate de que el paquete esté correctamente importado
)

func main() {
	// Ruta del directorio que contiene los archivos MP3
	ruta := "/home/grauler/Música/Escritorio/"

	// Llamada a la función MinarDirectorio
	songs, err := Minero.MinarDirectorio(ruta)
	if err != nil {
		log.Fatalf("Error al minar el directorio: %v", err)
	}

	// Mostrar la metadata extraída
	for _, song := range songs {
		fmt.Printf("Título: %s\nArtista: %s\nÁlbum: %s\nAño: %d\nGénero: %s\nTrack: %d\n\n", 
			song.Title, song.Artist, song.Album, song.Year, song.Genre, song.Track)
	}
}
