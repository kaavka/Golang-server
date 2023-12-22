# Go Users Management API

This project is a simple API for managing user data. It provides basic operations such as creating, retrieving, updating, and deleting user records.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)

## Introduction

The Go Users Management API is designed to handle CRUD operations for user entities. It uses the Gin web framework to create a simple RESTful API.

## Features

- List all users
- Create a new user
- Update an existing user
- Delete a user by ID

## Getting Started

### Prerequisites

- Go installed on your machine

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/kaavka/Golang-server
    cd ./Golang-server
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

## Usage

The API provides endpoints for managing user data.

1. **Get users list:** `GET /users`

  Example JSON response:

  ```json
  {
    "id": "b96bcaad-9a4e-4fc0-bdc0-ebfc51de52a6",
    "name": "John",
    "surname": "Doe",
    "age": 30,
    "email": "john.doe@example.com",
    "about": "Software Engineer",
    "iconColor": "blue"
  }...
  ```

2. **Post a new user:** `POST /users`

  Example JSON request:

  ```json
  {
    "name": "John",
    "surname": "Doe",
    "age": 30,
    "email": "john.doe@example.com",
    "about": "Software Engineer",
    "iconColor": "blue"
  }
  ```

  Example JSON response:

  ```json
  {
    "id": "b96bcaad-9a4e-4fc0-bdc0-ebfc51de52a6",
    "name": "John",
    "surname": "Doe",
    "age": 30,
    "email": "john.doe@example.com",
    "about": "Software Engineer",
    "iconColor": "blue"
  }
  ```
  Please note that the `id` field is generated automatically on server, so you don't have to do it on your own.

3. **Update an existing user:** `PUT /users`

  Example JSON request:

  ```json
  {
    "id": "user-id-to-edit",
    "name": "UpdatedName",
    "surname": "UpdatedSurname",
    "age": 25,
    "email": "updated.email@example.com",
    "about": "Updated Job Role",
    "iconColor": "green"
  }
  ```

  Example JSON response: 

  ```json
  {
    "id": "user-id-to-edit",
    "name": "UpdatedName",
    "surname": "UpdatedSurname",
    "age": 25,
    "email": "updated.email@example.com",
    "about": "Updated Job Role",
    "iconColor": "green"
  }
  ```

4. **Delete a user:** `DELETE /users?id=user-id-to-delete`

## Structure

The project follows a standard Go project structure:

- **/constants**
  - `constants.go`
  
- **/controller**
  - `user_controller.go`
  
- **/entity**
  - `user.go`
  
- **/service**
  - `user_service.go`
  
- **/utils**
  - `util_functions.go`
  
- `main.go`

## Configuration

The default server port is set to `:3000`, and user-related operations are available at the `/users` route. You can customize these configurations in the `constants/constants.go` file.

## Please enjoy &hearts;