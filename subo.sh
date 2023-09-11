#!/bin/bash

# Agregar y confirmar los cambios en Git
git add .
git commit -m "Ultimo Commit"

# Empaquetar el ejecutable en un archivo ZIP
rm -f main.zip  # Eliminar el archivo ZIP anterior si existe
zip main.zip main
go build main.go 

# Subir los cambios a GitHub
git push




