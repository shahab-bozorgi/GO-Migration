PostgreSQL Migration API with Golang

This project is a simple API built using Golang (Gin framework) that allows users to perform migrations and query a PostgreSQL database via an API. It provides a RESTful interface for executing database operations, including running migrations, performing basic CRUD operations, and executing custom SQL queries.
Features

    PostgreSQL Database Integration: Connect to a PostgreSQL database and perform various operations.
    Migrations: API endpoints to run database migrations for schema updates.
    Query Execution: Users can send custom SQL queries through the API to interact with the database.
    API Endpoints:
        Run Migrations: Trigger database schema updates.
        Execute Queries: Perform custom SQL queries like SELECT, INSERT, UPDATE, DELETE.

Prerequisites

    Go 1.16+ installed
    PostgreSQL database
    PostgreSQL driver for Golang (github.com/lib/pq)
