# CodingJunior_Assignment

## Overview

This project is a secure REST API built with Golang and the Fiber framework for user authentication and personal notes management. The API allows users to register, log in, and manage their own notes. Only authenticated users can access their notes.

## Tech Stack

- **Golang**: The programming language used for building the API.  
- **Fiber**: A fast HTTP web framework for Go.  
- **MySQL**: The database used for storing user and note data, accessed via GORM.  
- **JWT**: JSON Web Tokens for user authentication.  
- **Docker**: For containerization of the application and MySQL database.

## Folder Structure

```text
go-fiber-notes-api
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── wait-for-mysql.sh
├── handlers
│   ├── authHandler.go
│   ├── noteHandler.go
├── middleware
│   └── jwtMiddleware.go
├── models
│   ├── db.go
│   ├── note.go
│   └── user.go
├── routes
│   ├── auth.go
│   └── note.go
└── utils
    ├── hash.go
    └── jwt.go
