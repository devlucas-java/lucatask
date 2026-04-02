# LucaTask API

**LucaTask** is a task management backend built with **Go**, following **Clean Architecture** principles for scalability, maintainability, and testability.

The project uses **Chi** for routing, **GORM** for database access, **JWT authentication middleware**, **DTO validation**, and **automated tests**.

---

# 🇬🇧 English

## Overview

LucaTask is a modern REST API for task management designed with a professional backend structure.

### ✨ Features

* ✅ Clean Architecture
* ✅ RESTful API with **Chi Router**
* ✅ Database integration with **GORM**
* ✅ JWT authentication middleware
* ✅ DTO validation
* ✅ Unit and integration tests
* ✅ Middleware-based request authentication
* ✅ Secure password hashing with bcrypt
* ✅ Context-based authenticated user access
* ✅ Modular repository and use case layers

---

## 🏗️ Architecture

The project follows **Clean Architecture**:

```text
cmd/
internal/
 ├── delivery/      # handlers, routes, DTOs
 ├── domain/        # entities
 ├── usecase/       # business rules
 ├── infra/         # database, jwt, repositories
 └── shared/        # request context, helpers
pkg/
```

### Layers

* **Delivery** → HTTP handlers, routes, request validation
* **UseCase** → Business logic
* **Domain** → Core entities
* **Infrastructure** → Database, JWT, repositories

---

## 🔐 JWT Authentication Middleware

A JWT validation middleware was implemented to protect private routes.

### What it does

* Reads token from Authorization header
* Validates JWT signature
* Extracts claims:

    * `user_id`
    * `email`
    * `role`
* Stores auth data in request context
* Makes authenticated user data available inside handlers

### Example

```go
r.Route("/tasks", func(r chi.Router) {
    r.Use(middleware.AuthMiddleware(jwtService))
    r.Get("/", taskHandler.List)
})
```

---

## ✅ DTO Validation

DTO validation was added to improve request safety and avoid invalid payloads.

### Examples

* Required fields
* Email validation
* Password rules
* Update payload validation

This helps keep the **use case layer clean** and prevents invalid data from reaching the business logic.

---

## 🧪 Tests

The project includes tests for:

* Use cases
* Repositories
* Update flows
* JWT middleware behavior
* DTO validation

### Example tested flow

* Create user
* Update user
* Validate hashed password
* Assert persisted data

---

## ⚙️ Tech Stack

* **Go**
* **Chi Router**
* **GORM**
* **SQLite / PostgreSQL**
* **JWT**
* **bcrypt**
* **Testify**

---

## 🚀 Future Improvements

* Refresh token support
* Task ownership authorization
* Rate limiting middleware
* Redis cache
* Background jobs
* Docker deployment
* CI/CD pipeline

---

# 🇪🇸 Español

## Descripción

**LucaTask** es una API REST para gestión de tareas desarrollada en **Go**, siguiendo los principios de **Arquitectura Limpia**.

Está diseñada para ser **escalable, mantenible y segura**, utilizando **Chi**, **GORM**, **middleware JWT**, validación de DTOs y pruebas automatizadas.

---

## ✨ Características

* ✅ Arquitectura Limpia
* ✅ API REST con **Chi Router**
* ✅ Persistencia con **GORM**
* ✅ Middleware de autenticación JWT
* ✅ Validación de DTOs
* ✅ Pruebas unitarias e integración
* ✅ Hash seguro de contraseñas con bcrypt
* ✅ Acceso al usuario autenticado por contexto
* ✅ Separación por capas

---

## 🏗️ Arquitectura

El proyecto sigue **Clean Architecture**:

```text
cmd/
internal/
 ├── delivery/      # handlers, rutas, DTOs
 ├── domain/        # entidades
 ├── usecase/       # reglas de negocio
 ├── infra/         # base de datos, jwt, repositorios
 └── shared/        # contexto y utilidades
pkg/
```

### Capas

* **Delivery** → handlers HTTP, rutas, validación
* **UseCase** → lógica de negocio
* **Domain** → entidades principales
* **Infrastructure** → base de datos, JWT, repositorios

---

## 🔐 Middleware JWT

Se implementó un middleware para validar el token JWT y proteger rutas privadas.

### Funcionalidad

* Lee el token desde el header Authorization
* Valida la firma JWT
* Extrae claims:

    * `user_id`
    * `email`
    * `role`
* Guarda la información en el contexto
* Permite acceder al usuario autenticado en los handlers

### Ejemplo

```go
r.Route("/tasks", func(r chi.Router) {
    r.Use(middleware.AuthMiddleware(jwtService))
    r.Get("/", taskHandler.List)
})
```

---

## ✅ Validación de DTOs

La validación de DTOs fue añadida para evitar payloads inválidos.

### Ejemplos

* Campos obligatorios
* Validación de email
* Reglas de contraseña
* Validación de actualizaciones

Esto mantiene la capa de casos de uso más limpia y segura.

---

## 🧪 Tests

El proyecto incluye pruebas para:

* Casos de uso
* Repositorios
* Flujo de actualización
* Middleware JWT
* Validación de DTOs

### Flujo probado

* Crear usuario
* Actualizar usuario
* Verificar password hasheada
* Confirmar persistencia

---

## ⚙️ Stack Tecnológico

* **Go**
* **Chi Router**
* **GORM**
* **SQLite / PostgreSQL**
* **JWT**
* **bcrypt**
* **Testify**

---

## 🚀 Mejoras Futuras

* Refresh tokens
* Autorización por propietario de tarea
* Rate limiting
* Redis
* Jobs en background
* Docker
* CI/CD

---

## 👨‍💻 Author

Developed by **devLucas-Java** with focus on **backend scalability, security, and clean architecture \
github: https://github.com/devlucas-java \
linkedin: https://www.linkedin.com/in/devlucas-java/ \