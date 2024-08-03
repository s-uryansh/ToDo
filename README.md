# ToDo Task API

A simple ToDo task API built with Go and MariaDB

## Table of Contents

* [Features](#features)
* [Getting Started](#getting-started)
* [API Endpoints](#api-endpoints)
* [Database Schema](#database-schema)
* [Contributing](#contributing)
* [License](#license)

## Features

* User registration and login functionality
* Task creation endpoint

### Prerequisites

* Go
* MariaDB 

## API Endpoints

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
