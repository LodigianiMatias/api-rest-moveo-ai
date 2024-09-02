# Prueba técnica moveo.ai

Este proyecto es una API RESTful en Go que permite gestionar tareas utilizando una base de datos MySQL. La API ofrece operaciones CRUD (Crear, Leer, Actualizar, Eliminar). Incluye documentación generada automáticamente con Swagger.
Se encuentra también un archivo para ejecutar los endpoints con la extensión Thunder Client.

# Instalación
- [Go](https://golang.org/doc/install)
- [MySQL](https://dev.mysql.com/downloads/mysql/) para la base de datos

# Base de datos 
Se debe crear un archivo .env que reciba los siguientes parametros:

DB_USER= *Su usuario de MySQL Server*

DB_PASSWORD= *password*

DB_HOST=localhost

DB_PORT=3306

DB_NAME=testing-moveo-ai

## Instalar dependecias
### go mod dowload

## Ejecución de la app
### go run main.go

# Rutas-de-la-api

## Crear una nueva tarea

Tener en cuenta que el status solo puede recibir 3 parametros: completed, in-progress, pending.

Ruta: /tasks

Método: POST


Copiar código:

{

  "title": "Nueva Tarea",
  
  "description": "Descripción de la tarea",
  
  "status": "pending"
  
}

## Obtener todas las Tareas
Ruta: /tasks

Método: GET

## Obtener una Tarea por ID

Ruta: /tasks/{id}

Método: GET

## Actualizar una Tarea
Ruta: /tasks/{id}

Método: PUT

Copiar código:

{

  "title": "Título Actualizado",
  
  "description": "Descripción Actualizada",
  
  "status": "in-progress"
  
}

## Eliminar una Tarea
Ruta: /tasks/{id}

Método: DELETE
![image](https://github.com/user-attachments/assets/1872b819-0826-4bae-8d80-218976916fe7)

### http://localhost:8080/swagger/index.html
![image](https://github.com/user-attachments/assets/b7a0a9fc-8238-4e43-be63-afef12a6eeb4)
