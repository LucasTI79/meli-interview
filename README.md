# Meli Interview Marketplace & Backend

This repository contains a fullstack project with a Go backend (REST API) and a Next.js frontend, simulating a marketplace with product management and shopping cart features. The backend is containerized with Docker and includes Swagger documentation for easy API exploration.

## Information

- **Title**: Meli Interview Marketplace
- **Version**: 1.0
- **Backend Host**: http://localhost:8080
- **Frontend Host**: http://localhost:3000

## Prerequisites

Before getting started, make sure you have the following installed:

- [Go](https://golang.org/dl/): Go programming language (for backend development)
- [Node.js](https://nodejs.org/): Node.js and npm (for frontend development)
- [Docker](https://www.docker.com/get-started): For running services in containers

## Installation

Clone the repository:

```
git clone https://github.com/LucasTI79/meli-interview.git
cd meli-interview
```

Build and start all services with Docker Compose:

```
docker-compose up backend frontend -d
```

Or, to run only the backend or frontend locally:

### Backend (Go)

```
cd app
cp .env.example .env
go run cmd/http/main.go
```

### Frontend (Next.js)

```
cd marketplace
cp .env.example .env
npm install
npm run dev
```

## Running the Application

- Backend API: [http://localhost:8080](http://localhost:8080)
- Frontend: [http://localhost:3000](http://localhost:3000)

## API Documentation

After starting the backend, access the Swagger UI for API documentation and testing:

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Or scalar docs

- [http://localhost:8080/docs](http://localhost:8080/docs)

## Features

- Product listing and details
- Shopping cart management
- RESTful API for products
- Dockerized backend and frontend
- Swagger/OpenAPI documentation

## Example API Usage

You can use [curl](https://curl.se/) or [Postman](https://www.postman.com/) to test the API endpoints. Example:

```
curl -X GET http://localhost:8080/products
```

## Data Models (Backend)

### Product Entity

- `id` (string): Unique product identifier
- `name` (string): Product name
- `description` (string): Product description
- `price` (float): Product price
- `image` (string): Image URL or path

### API Error Response

- `code` (integer): Error code
- `error` (string): Error description
- `message` (string): Error message

## Main Endpoints

### Products

- `GET /products` — List all products
- `GET /products/{productId}` — Get product details by ID

### Category

- `GET /categories` — List all categories
- `GET /categories/{categoryName}` - Get category by category name

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

If you have any questions or need support, feel free to open an issue. Enjoy exploring the Meli Interview Marketplace!