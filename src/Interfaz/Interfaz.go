package Interfaz

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "log"
)

func CrearVentana(onClickFunc func(ruta string), onListFunc func() ([]string, error), onSearchFunc func(string) ([]string, error)) {
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
    scrollContainer.Hide()

    // Declarar los botones
    var volver *widget.Button
    var ms *widget.Button

    ms = widget.NewButton("Canciones almacenadas", func() {
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
        scrollContainer.Show() 
        volver.Show() 
        ms.Hide() 
    })

    volver = widget.NewButton("Volver", func() {
        scrollContainer.Hide() 
        volver.Hide()
        ms.Show()
    })
    volver.Hide() 

    // Crear botón de salida
    exit := widget.NewButton("Salir", func() {
        a.Quit() 
    })

	var searchEntry *widget.Entry
	
    // Crear entrada de búsqueda y botón de búsqueda
    searchEntry = widget.NewEntry()
    searchEntry.SetPlaceHolder("Buscar por álbum, título o género")
	searchEntry.Hide()
	
	search := widget.NewButton("Buscar", func() {
		searchEntry.Show()
        criterio := searchEntry.Text
        canciones, err := onSearchFunc(criterio)
        if err != nil {
            log.Printf("Error al buscar las canciones: %v", err)
            resultLabel.SetText("Error al buscar las canciones")
            return
        }
        listLabel.SetText("")
        for _, cancion := range canciones {
            listLabel.SetText(listLabel.Text + cancion + "\n")
        }
        scrollContainer.Show()
        volver.Show()
        ms.Hide()
    })

    w.SetContent(container.NewVBox(
        rutaEntry,
        btn,
        searchEntry,
        search,
        ms,
        scrollContainer,
        volver,
        exit,
        resultLabel, 
    ))

    w.ShowAndRun()
}
