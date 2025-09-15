# How to Run the Project

This guide explains how to run the Meli Interview Marketplace project, which includes a Go backend (REST API) and a Next.js frontend. You can run both services using Docker Compose or run each one locally for development.

---

## Prerequisites

- [Go](https://golang.org/dl/) (for backend development)
- [Node.js](https://nodejs.org/) and npm (for frontend development)
- [Docker](https://www.docker.com/get-started) and Docker Compose

---

## 1. Clone the Repository

```
git clone https://github.com/LucasTI79/meli-interview.git
cd meli-interview
```

---

## 2. Run with Docker Compose (Recommended)

This will start both backend and frontend containers:

```
docker-compose up -d -V
```

- Backend API: http://api.localhost
- Frontend: http://app.localhost

To stop the services:

```
docker-compose down
```

---

## 3. Run Backend Locally (Go)

1. Open a terminal and navigate to the backend folder:

```
cd app
```

2. Create environment file

```
cp .env.example .env
```

3. Run the backend server:

```
go run cmd/http/main.go
```

- The API will be available at http://localhost:8080

---

## 4. Run Frontend Locally (Next.js)

1. Open a new terminal and navigate to the frontend folder:

```
cd marketplace
```

2. Create environment file

```
cp .env.example .env
```

3. Install dependencies:

```
npm install
```

4. Start the development server:

```
npm run dev
```

- The frontend will be available at http://localhost:3000

---

## 5. API Documentation

After starting the backend, access the Swagger UI for API documentation and testing:

- http://localhost:8080/swagger/index.html (local)
- http://api.localhost/swagger/index.html (with Traefik)

Or access scalar docs:

- http://localhost:8080/docs (local)
- http://api.localhost/docs (with Traefik)

---

## 6. Running Tests (Backend)

To run backend tests:

```
cd app
make test
```

---

## Troubleshooting

- Make sure required ports (8080 for backend, 3000 for frontend) are free.
- If you change backend code, restart the backend server.
- For Docker issues, try `docker-compose down` and then `docker-compose up` again.

---

If you have any questions or issues, please open an issue in the repository.
