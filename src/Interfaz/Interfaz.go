package Interfaz

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
	"log"
)

func CrearVentana(onClickFunc func(ruta string), onListFunc func() ([]string, error)) {
    a := app.New()
    w := a.NewWindow("Music Database Miner")

    rutaEntry := widget.NewEntry()
    rutaEntry.SetPlaceHolder("Introduce la ruta del directorio que contiene los archivos MP3")
    resultLabel := widget.NewLabel("")

    btn := widget.NewButton("Minar Directorio", func() {
        ruta := rutaEntry.Text
        onClickFunc(ruta)
        resultLabel.SetText("Metadatos guardados exitosamente")
    })

	listLabel := widget.NewLabel("")
	scrollContainer := container.NewScroll(listLabel)
	scrollContainer.SetMinSize(fyne.NewSize(400, 300))
	
    ms := widget.NewButton("Canciones almacenadas", func() {
        canciones, err := onListFunc()
        if err != nil {
            log.Printf("Error al listar las canciones: %v", err)
            resultLabel.SetText("Error al listar las canciones")
            return
        }
        listLabel.SetText("")
        for _, cancion := range canciones {
            listLabel.SetText(listLabel.Text + cancion + "\n")
        }
    })

    w.SetContent(container.NewVBox(
        rutaEntry,
        btn,
        ms,
		scrollContainer,
        resultLabel,
    ))

    w.ShowAndRun()
}
