# tzAPI

tzAPI is a simple Go application for managing personal information using a RESTful API. It includes basic CRUD operations (Create, Read, Update, Delete) for a 'Person' entity.
## Installation
1. Create a PostgreSQL database and configure the connection in conn.env:

```bash
DATABASE_URL=your-database-url
```
2. Apply database migrations:
```bash
go run cmd/main.go migrate
```
3. Run the application:

```bash
go run cmd/main.go
```

## Usage
The API exposes the following endpoints:

*POST /person: Create a new person
*GET /people: Get a list of people with optional query parameters
*GET /person/{id}: Get a specific person by ID
*PUT /person/{id}: Update a person's information
*DELETE /person/{id}: Delete a person by ID