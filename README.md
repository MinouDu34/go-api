Overview
This project is a RESTful API built using Go. The API provides endpoints for managing film databases. It follows standard REST conventions and is designed to be lightweight and efficient.

Features
- RESTful endpoints for CRUD operations
- Logging

Prerequisites
Before running this project, ensure that you have the following installed:

- Go 1.XX or higher

Getting Started
Clone the repository

```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```

Install dependencies
Use go mod to install the project dependencies:

```bash
go mod tidy
```

Running the API
To start the API, use the following command:

```bash
go run main.go
```

The API will be available at http://localhost:8080.


API Endpoints
Here is a brief overview of the available endpoints:

GET /api/v1/resource - Retrieves a list of resources
GET /api/v1/resource/{id} - Retrieves a specific resource by ID
POST /api/v1/resource - Creates a new resource
PUT /api/v1/resource/{id} - Updates a specific resource by ID
DELETE /api/v1/resource/{id} - Deletes a specific resource by ID
