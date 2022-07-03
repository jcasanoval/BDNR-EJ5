Primero que nada actualizar el archivo .env incluido.

Luego se debe agregar el siguiente indice a la base de datos con el fin de optimizar y evitar que se repitan los emails:
`db.users.createIndex( { "email": 1 }, { unique: true } )`

Para correr: `go run main.go`

Para probar los endpoints con postman utilizar la colección incluida en el archivo BDNR.postman_collection.json
