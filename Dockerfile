# Descarga Go desde la pagina oficial
FROM golang:1.18-alpine

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia de los archivos de modulos
COPY go.mod ./
COPY go.sum ./

# Descarga las dependencias de los modulos
RUN go mod download

# Copia el resto de los archivos del directorio del proyecto
COPY . .

# Construye la aplicacion Go
RUN go build -o main .

# Expone el puerto a utilizar
EXPOSE 8080

# Comando para ejecutar el archivo principal de Go
CMD ["./main"]
