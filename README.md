# IN THE NAME OF ALLAH

**Store Product Management Microservice**

This repository hosts a microservice designed for managing products within a retail store environment. It utilizes Gin for API management, GORM for database interaction with PostgreSQL, and Redis for caching. 

## Technologies Used

- **Gin**: High-performance web framework for building HTTP services.
- **GORM**: Object-relational mapping (ORM) library for Golang, enabling seamless communication with PostgreSQL.
- **Redis**: In-memory data structure store used for caching.

## Features

- **Authentication**: Requires a login with predefined credentials (`Chek` as username and `123456` as password).
- **CRUD Operations**: Capabilities for creating, reading, updating, and deleting product information based on business name (BN) or business ID (BI).

## Getting Started

Before proceeding, ensure you have Go installed on your system.

### Setup

1. Clone the repository to your local machine:   
git clone https://github.com/yourusername/store-product-management.git
   
2. Navigate to the project directory:
   
cd store-product-management

3. Build the application:
   
go build .

4. Run the application:
   
./store-product-management


## Authentication

To authenticate, send a request to the `/login` endpoint with the required credentials in the header:
- **User**: `Chek`
- **Password**: `123456`

Successful authentication will grant access to the other endpoints.

## API Endpoints

### Authentication
- **POST /login**
  - Required for accessing other endpoints.

### Product Management
- **POST /insert**
  - Insert a new product.
- **POST /updateBN/:name**
  - Update product information by business name.
- **POST /updateBI/:id**
  - Update product information by business ID.
- **GET /getBN/:name**
  - Retrieve product information by business name.
- **GET /getBI/:id**
  - Retrieve product information by business ID.
- **GET /removeBI/:id**
  - Remove a product by business ID.
- **GET /removeBN/:name**
  - Remove a product by business name.

## Contributing

Contributions are welcome If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.
