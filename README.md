# Music-database

## Gael Arturo Leon Garcia

Este proyecto es un reproductor de archivos mp3 haciendo uso de una interfaz grafica y manejo de una base de datos. Esta hecho en Golang, como intefaz grafica se uso fyne y la base de datos es manejada por SQLite.

## Requisitos

Antes de utilizar este programa, asegúrate de tener instalados los siguientes componentes en tu sistema:

- **Go (Golang)**: Lenguaje de programación utilizado para desarrollar la aplicación.
- **Fyne**: Biblioteca para la creación de interfaces gráficas.
- **SQLite**: Sistema de gestión de bases de datos ligero.
- **Herramientas de Compilación**: Dependencias necesarias para compilar ciertas librerías de Go.

Para usar este programa es necesario tener instalado go en tu equipo:

### En Arch Linux
```
sudo pacman -S go
```

### En Fedora
```
sudo dnf install go
```

### En Ubuntu/Debian 
```
sudo apt install go
```


Se recomienda actualizar el equipo primero, para revisar la version solo tiene que escribir:
 go version

## Uso del programa

Dentro de la carpeta del programa, puede ejectutarlo con el siguiente comando en la terminal:
```
./Music-database
```
### Minado

Para minar tienes que escribir la ruta con el siguiente formato:
``` 
/tu/ruta/del/directorio/
```
Por ejemplo: 
```/home/user/Música/```

Se debe instalar la version de go 1.23.2
```
https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
```
 Puedes revisar la base de datos con sqlite3:

### En Arch Linux
```
sudo pacman -S sqlite
```

### En Fedora
```
sudo dnf install sqlite
```
Para instalar las bibliotecas de desarrollo de SQLite:

```
sudo dnf install sqlite-devel
```

### En Ubuntu/Debian 
```
sudo apt install sqlite3
```
Para las bibliotecas de desarrollo de SQLite, si es necesario:
```
sudo apt install libsqlite3-dev
```
Puedes verificar que SQLite está instalado correctamente ejecutando:
```
sqlite3 --version
```
