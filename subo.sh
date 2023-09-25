#!/bin/bash

# Agregar y confirmar los cambios en Git
git add .
git commit -m "Ultimo Commit"

# Empaquetar el ejecutable en un archivo ZIP
rm -f main.zip  # Eliminar el archivo ZIP anterior si existe
zip main.zip main
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

# Subir los cambios a GitHub
git push




