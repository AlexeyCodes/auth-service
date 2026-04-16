# Auth Service

Auth Service is an authentication microservice implemented in Go. Its main purpose is to handle user registration, authentication, and access token issuance.

## Features

- User registration
- Authentication using username and password
- Issuing and validating JWT tokens
- Cookie-based session support
- Authorization checks for protected resources

## Technologies Used

- Go
- Gin (web framework)
- PostgreSQL
- pgx (PostgreSQL driver)
- godotenv (environment variable loader)

## Project Architecture

The project follows microservice architecture principles and is designed to be part of a larger system. In its current implementation, the service is divided into several core components: request handlers, database layer, data models, and configuration files.

## Configuration

All configuration parameters, including the database connection string, JWT secret, and service port, are defined via a `.env` file.

## Running the Service

To run the service, use the command:

```bash
go run cmd/login/main.go
