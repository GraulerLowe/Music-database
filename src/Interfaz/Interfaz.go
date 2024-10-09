package Interfaz

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func CrearVentana(onClickFunc func(ruta string)) {
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
	
	baseEntry := widget.NewEntry()

	ms := widget.NewButton("Canciones almacenadas", func() {
		base := baseEntry.Text
		onClickFunc(base)
		resultLabel.SetText("Canciones en la base de datos")
	})
		
		
    w.SetContent(container.NewVBox(
        rutaEntry,
        btn,
		ms,
        resultLabel,
    ))

    w.ShowAndRun()
}
