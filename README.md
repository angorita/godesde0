# godesde0
Repositorio de curso de Golang desde 0
go mod init github.com/angorita/godesde0
import(
	"github.com/angorita/godesde0/variables"
)
func main() {
	variables.MostrarEnteros()
}
05:04$git status
En la rama main
Tu rama está actualizada con 'origin/main'.

Cambios no rastreados para el commit:
  (usa "git add <archivo>..." para actualizar lo que será confirmado)
  (usa "git restore <archivo>..." para descartar los cambios en el directorio de trabajo)
        modificados:     README.md

Archivos sin seguimiento:
  (usa "git add <archivo>..." para incluirlo a lo que será confirmado)
        .gitignore
        go.mod
        main.go
        tmp/
        variables/

sin cambios agregados al commit (usa "git add" y/o "git commit -a")
[~/godesde0]
05:04$git add .
[~/godesde0]
05:04$git commit -m "primer commit"
[main 983cdc3] primer commit
 7 files changed, 33 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 go.mod
 create mode 100644 main.go
 create mode 100644 tmp/build-errors.log
 create mode 100755 tmp/main
 create mode 100644 variables/enteras.go
[~/godesde0]
05:04$git push
Enumerando objetos: 13, listo.
Contando objetos: 100% (13/13), listo.
Compresión delta usando hasta 12 hilos
Comprimiendo objetos: 100% (8/8), listo.
Escribiendo objetos: 100% (11/11), 1.27 MiB | 1.57 MiB/s, listo.
Total 11 (delta 0), reusados 0 (delta 0), pack-reusados 0
To https://github.com/angorita/godesde0.git
   bc7039f..983cdc3  main -> main
[~/godesde0]