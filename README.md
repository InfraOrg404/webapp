# webapp
GoLang API Application

# Health Check API

This is a simple Health Check API implemented in Go. The API provides an endpoint `/healthz` that checks the connectivity to the database and returns appropriate HTTP status codes.

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [MySQL](https://dev.mysql.com/downloads/mysql/) or any other database of your choice
- A code editor or IDE of your choice (e.g., [Visual Studio Code](https://code.visualstudio.com/))

## Build and Deploy Instructions

### Clone the Repository

First, clone the repository to your local machine using Git:

```bash
git clone <repository-url>
cd health-check-api
```

To build the application, run the following command in the project root directory:

```bash
go build -o health-check-api
```

Test the Endpoint : 

```bash
curl -X GET http://localhost:8080/healthz
```