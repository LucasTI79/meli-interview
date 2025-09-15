# Meli Interview Marketplace & Backend

This repository contains a fullstack project with a Go backend (REST API) and a Next.js frontend, simulating a marketplace with product management and shopping cart features. The backend is containerized with Docker and includes Swagger documentation for easy API exploration.

## Information

- **Title**: Meli Interview Marketplace
- **Version**: 1.0
- **Backend Host**:
	- http://localhost:8080 (default Docker Compose or local)
	- http://api.localhost (with Traefik)
- **Frontend Host**:
	- http://localhost:3000 (default Docker Compose or local)
	- http://app.localhost (with Traefik)

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


### Option 1: Default Docker Compose (localhost)

```
docker-compose up backend frontend -d
```

Access:
- Backend: http://localhost:8080
- Frontend: http://localhost:3000

### Option 2: With Traefik Reverse Proxy (recommended for custom domains)

```
docker-compose -f docker-compose-traefik.yml up -d
```

Access:
- Backend: http://api.localhost
- Frontend: http://app.localhost

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

- If using default compose: 
	- Backend API: [http://localhost:8080](http://localhost:8080)
	- Frontend: [http://localhost:3000](http://localhost:3000)
- If using Traefik:
	- Backend API: [http://api.localhost](http://api.localhost)
	- Frontend: [http://app.localhost](http://app.localhost)

## API Documentation


After starting the backend, access the Swagger UI for API documentation and testing:

- Default compose: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Traefik: [http://api.localhost/swagger/index.html](http://api.localhost/swagger/index.html)

Or scalar docs:
- Default compose: [http://localhost:8080/docs](http://localhost:8080/docs)
- Traefik: [http://api.localhost/docs](http://api.localhost/docs)

## Features

- Product listing and details
- Shopping cart management
- RESTful API for products
- Dockerized backend and frontend
- Swagger/OpenAPI documentation

## Example API Usage


You can use [curl](https://curl.se/) or [Postman](https://www.postman.com/) to test the API endpoints. Example:

Default compose:
```
curl -X GET http://localhost:8080/products
```
Traefik:
```
curl -X GET http://api.localhost/products
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