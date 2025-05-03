# Microservices-Common

Esta es una biblioteca de utilidades para microservicios en Go diseñada para el proyecto Class Connect del GRUPO 5.

## Tabla de Contenidos

- [Instalación](#instalación)
- [Módulos](#módulos)
  - [Logger](#logger)
  - [Database](#database)
  - [Middleware](#middleware)
  - [Models](#models)
  - [Repository](#repository)
  - [Utils](#utils)

## Instalación

Usar el siguiente comando:
```bash
go get github.com/Class-Connect-GRUPO-5/microservices-common
```

## Módulos

### Logger

Proporciona funciones para inicializar y configurar logs estructurados utilizando logrus.

**Componentes principales:**
- `Logger`: Variable global que contiene la instancia del logger.
- `InitLogger`: Función para inicializar el logger con un nivel específico y, opcionalmente, un archivo de salida.

### Database

Ofrece funcionalidades para conectarse a una base de datos PostgreSQL y ejecutar migraciones.

**Componentes principales:**
- `DB`: Variable global que contiene el pool de conexiones a la base de datos.
- `Connect`: Función para establecer la conexión con la base de datos.
- `RunMigrations`: Función para ejecutar migraciones de la base de datos.

### Middleware

Proporciona interceptores para las rutas de Gin, principalmente para autenticación y autorización.

**Componentes principales:**
- `RequireRole`: Middleware para verificar si un usuario tiene el rol requerido.
- `ExtractUserJWT`: Función para extraer y verificar un JWT del contexto de la petición.

### Models

Contiene interfaces y estructuras para representar distintos tipos de respuestas API.

**Componentes principales:**
- `APIResponse`: Interfaz para respuestas API genéricas.
- `Model`: Interfaz para modelos serializables.
- `ProblemDetails`: Estructura para errores según RFC 7807.
- `SuccessDetails`: Estructura para respuestas exitosas.

### Repository

Ofrece una implementación genérica del patrón repositorio para interactuar con la base de datos.

**Componentes principales:**
- `Repository`: Clase genérica para operaciones CRUD.
- `QueryParser`: Interfaz para generar consultas SQL.

### Utils

Contiene funciones de utilidad diversas para tareas comunes en microservicios.

**Componentes principales:**
- `JWT`: Funciones para generar y verificar tokens JWT.
- `Password`: Funciones para hashear y verificar contraseñas.
- `PinGenerator`: Interfaz y implementación para generar PINs de autenticación.
- `MailSender`: Interfaz y implementación para enviar correos de verificación.
- `Error/Success Handlers`: Funciones para manejar errores y éxitos de manera consistente.