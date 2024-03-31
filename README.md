# Go General E-Commerce API

This is a simple e-commerce API with JWT authentication implemented in Go.

## Features

- User register and login
- Products CRUD
- JWT authentication

## Prerequisites

- Go programming language installed on your system
- MySQL or another compatible database installed and running
- `.env` file containing environment variables (e.g., database connection details, server port)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/Kei-K23/go-ecom.git
```

2. Navigate to the project directory:

```bash
cd go-ecom

```

3. Create a .env file in the project root and add the following environment variables:

```bash
SECRET_KEY=<SECRET_KEY>
```

## Usage

1. Install dependencies:

```bash
go mod tidy
```

1. Run migration

```bash
make migration
```

2. Push database table

```bash
make migrate-up
```

2. Run server

```bash
make run
```

3. Access the API endpoints using tools like cURL or Postman.

## API endpoints

- `POST /register`: Register a new user
- `POST /login`: Login user
- `GET /products`: Retrieve all products
- `GET /products/{id}`: Retrieve specific product by ID
- `POST /products`: Create new product
- `PUT /products/{id}`: Update product by ID
- `DELETE /products/{id}`: Delete product by ID
