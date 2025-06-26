# Restaurant Management System Backend

A robust backend system for restaurant management built with Go (Golang) and MongoDB.

## 🚀 Features

- User Authentication and Authorization
- Food Item Management
- Menu Management
- Order Processing
- Table Management
- Order Item Tracking
- Invoice Generation

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Database:** MongoDB
- **Authentication:** JWT (JSON Web Tokens)

## 📋 Prerequisites

- Go 1.16 or higher
- MongoDB
- Git

## 🔧 Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd resManagementSystem_golang
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
export PORT=8080
export MONGODB_URI="your_mongodb_connection_string"
```

4. Run the application:
```bash
go run main.go
```

## 📁 Project Structure

```
.
├── controllers/     # Request handlers
├── database/       # Database connection and configuration
├── helpers/        # Utility functions
├── middleware/     # Custom middleware
├── models/         # Data models
├── routes/         # API routes
├── main.go         # Application entry point
└── go.mod          # Go module file
```

## 🔐 API Endpoints

The API includes the following main routes:
- User Routes
- Food Routes
- Menu Routes
- Order Routes
- Table Routes
- Order Item Routes
- Invoice Routes

## 🔒 Authentication

The system uses JWT-based authentication. Protected routes require a valid JWT token in the Authorization header.

## 🚀 Running Tests

```bash
go test ./...
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📧 Contact

For any queries or support, please open an issue in the repository. 
