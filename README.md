# Product Inventory System

## Overview

The Product Inventory System is a simple CRUD application built using React for the frontend and Go for the backend. It allows users to manage a product inventory by adding, updating, deleting, and viewing products. The application uses MongoDB for storing product data.

## Features

- **Add Product**: Add new products to the inventory.
- **Update Product**: Edit details of existing products.
- **Delete Product**: Remove products from the inventory.
- **View Products**: Display a list of all products in the inventory.

## Tech Stack

- **Frontend**: React
- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **Styling**: CSS

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js 18+
- MongoDB instance

### Setup

#### Backend

1. **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd go-inventory
    ```

2. **Navigate to the backend directory:**

    ```bash
    cd backend
    ```

3. **Set up environment variables:**

    Create a `.env` file in the `backend` directory and add the following environment variables:

    ```plaintext
    MONGO_URI=mongodb://<username>:<password>@localhost:27017/inventory
    ```

4. **Install dependencies:**

    ```bash
    go mod tidy
    ```

5. **Run the backend server:**

    ```bash
    go run cmd/main.go
    ```

#### Frontend

1. **Navigate to the frontend directory:**

    ```bash
    cd frontend
    ```

2. **Set up environment variables:**

    Create a `.env` file in the `frontend` directory and add the following environment variable:

    ```plaintext
    REACT_APP_API_URL=http://localhost:8080
    ```

3. **Install dependencies:**

    ```bash
    npm install
    ```

4. **Start the frontend development server:**

    ```bash
    npm start
    ```

## Usage

1. Open your browser and navigate to `http://localhost:3000` to access the frontend application.
2. Use the form to add new products or update existing ones.
3. View the list of products and perform actions like edit or delete.

## API Endpoints

- **GET /products**: Retrieve all products.
- **POST /products**: Add a new product.
- **PUT /products/:id**: Update a product by ID.
- **DELETE /products/:id**: Delete a product by ID.

## Troubleshooting

- **CORS Issues**: Ensure the backend is configured to handle CORS requests from the frontend.

    Example CORS setup in Go:

    ```go
    import (
        "github.com/gin-contrib/cors"
        "github.com/gin-gonic/gin"
    )

    func main() {
        r := gin.Default()

        r.Use(cors.Default())

        // Your routes here

        r.Run(":8080")
    }
    ```

- **Network Errors**: Verify that the backend server is running and accessible at the specified API URL.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to submit issues or pull requests to improve the application.

## Contact

For any questions or feedback, please reach out to [your-email@example.com](mailto:your-email@example.com).

