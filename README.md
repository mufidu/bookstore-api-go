# Online Bookstore Management System

Ported from [bookstore-api](https://github.com/mufidu/bookstore-api) to Go.

This project is a backend system for managing an online bookstore. It allows customers to browse, search, and purchase books while providing functionality for inventory and customers management for admins.

## Features

### CI/CD

-   **Testing:** Unit and integration tests are run on GitHub Actions every time a commit is pushed.

### User Authentication

-   **Roles:** Two user roles are supported: customer and admin.
-   **Customer Authentication:** Customers can register, log in, and update their profile information.
-   **Admin Management:** Admins can manage customers' accounts.

### Book Management

-   **Add Books:** Admins can add new books to the inventory via an API endpoint.
-   **Update Books:** An endpoint is available for admins to update existing book details.
-   **Retrieve Books:** Users can retrieve a list of books with filtering options; genre, author, and year.

### Shopping Cart

-   **Add to Cart:** Customers can add books to their shopping cart.
-   **Update Cart:** Customers can update the quantity or remove items from their cart.
-   **Total Price Calculation:** The total price of items in the cart is automatically calculated each time the cart is updated.

### Order Processing

-   **Place Orders:** Customers can place orders to checkout their carts.
-   **Inventory Deduction:** The ordered quantity is deducted from the inventory upon order placement.
-   **Payment Gateway:** Midtrans payment gateway is integrated to process order payments. Customers can pay using QRIS.
-   **Email Notification:** Customers receive an email confirmation after placing an order.

### Inventory Management

-   **View and Manage Inventory:** Admins can view and manage the current inventory.

### Security

-   **Validation and Error Handling:** All API endpoints have proper validation and error handling.
-   **Sensitive Information:** Sensitive information is securely encrypted.

### Logging and Monitoring

-   **Logging:** Important events and errors are logged with Morgan.
-   **Monitoring:** Basic monitoring for API performance is implemented using Prometheus.

### Testing

-   **Unit Tests:** Important API endpoints (book and order) have unit tests.
-   **Integration Testing:** Integration testing ensures components work together seamlessly.

## Documentation

The API is documented using Swagger and can be accessed at:

```
http://localhost:9000/swagger/index.html
```

## Database Schema

![Database Schema](https://raw.githubusercontent.com/mufidu/jobhun-devops-test/main/Screenshot%202024-06-28%20at%2021.57.30.jpg)

## Folder structure

```
golang-rest-api-template/
|-- bin/
|-- cmd/
|   |-- server/
|       |-- main.go
|-- pkg/
|   |-- api/
|       |-- handler.go
|       |-- router.go
|   |-- models/
|       |-- user.go
|   |-- database/
|       |-- db.go
|-- scripts/
|-- Dockerfile
|-- go.mod
|-- go.sum
|-- README.md
```

### Explanation of Directories and Files:

1. **`bin/`**: Contains the compiled binaries.

2. **`cmd/`**: Main applications for this project. The directory name for each application should match the name of the executable.

    - **`main.go`**: The entry point.

3. **`pkg/`**: Libraries and packages that are okay to be used by applications from other projects. 

    - **`api/`**: API logic.
        - **`handler.go`**: HTTP handlers.
        - **`router.go`**: Routes.
    - **`models/`**: Data models.
    - **`database/`**: Database connection and queries.

4. **`scripts/`**: Various build, install, analysis, etc., scripts.

## Getting Started

### Tech Stack

- Go 1.15+ as the programming language.
- gin as the HTTP web framework.
- Docker and Docker Compose for containerization.
- PostgreSQL for the database.
- GORM as the ORM.
- Midtrans for the payment gateway.
- Swagger for API documentation.

### Installation

1. Clone the repository

```bash
git clone https://github.com/mufidu/bookstore-api-go.git
```

2. Navigate to the directory

```bash
cd bookstore-api-go
```

3. Build and run the Docker containers

```bash
make setup && make build && make up
```

### Environment Variables

You can set the environment variables in the `.env` file. Here are some important variables:

- `POSTGRES_HOST`
- `POSTGRES_DB`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_PORT`
- `JWT_SECRET`
- `API_SECRET_KEY`
- `MIDTRANS_SANDBOX_SERVER_KEY`
- `MIDTRANS_SANDBOX_CLIENT_KEY`
