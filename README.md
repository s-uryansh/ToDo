# ToDo Task API

A simple ToDo task API built with Go and MariaDB

## Table of Contents

* [Features](#features)
* [Getting Started](#getting-started)
* [API Endpoints](#api-endpoints)
* [Database Schema](#database-schema)
* [License](#license)

## Features

* User registration and login functionality
* Storing user data with Redis
* Task creation endpoint

### Prerequisites

* Go
* MariaDB 

## API Endpoints
* Registering a new user
* Logging in an existing user
* Creating a new task
* Getting an existing task
### User Endpoints

* `POST /register`: Register a new user
	+ Request Body: `{"Name": "string", "Email": "string", "Password": "string"}`
	+ Response: `{"New User": "Name" + " has succesfully created account"}`
* `POST /login`: Login an existing user
	+ Request Body: `{"Email": "string", "Password": "string"}`
	+ Response: `{"username": " logged in successfully"}`

    ### Task Endpoints

* `POST /todo`: Create a new task
	+ Request Body: `{"task": "string", "description": "string" , "status" : "bool" , "UserID" : "int"}`
	+ Response: `{"message": "todo created successfully"}`
* `GET /todo` : Get exisiting task
    + Params : `"ID" : "int"`
    + Response: `{"description" : "string" , "task" : "string"}`

## Database Schema

### users table

| Field    | Type         | Null | Key | Default | Extra          |
+----------+--------------+------+-----+---------+----------------+
| id       | int(11)      | NO   | PRI | NULL    | auto_increment |
| name     | varchar(50)  | NO   |     | NULL    |                |
| email    | varchar(100) | NO   | UNI | NULL    |                |
| password | varchar(255) | NO   |     | NULL    |                |
| role     | varchar(50)  | NO   |     | NULL    |                |
+----------+--------------+------+-----+---------+----------------+

### tasks table

| Field       | Type         | Null | Key | Default | Extra          |
+-------------+--------------+------+-----+---------+----------------+
| id          | int(11)      | NO   | PRI | NULL    | auto_increment |
| user_id     | int(11)      | YES  |     | NULL    |                |
| task        | varchar(255) | YES  |     | NULL    |                |
| description | text         | YES  |     | NULL    |                |
| status      | tinyint(1)   | YES  |     | NULL    |                |
+-------------+--------------+------+-----+---------+----------------+