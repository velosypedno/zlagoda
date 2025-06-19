# ZLAGODA - Supermarket Management System

A comprehensive REST API for supermarket management built with Go, Gin, and PostgreSQL.

## 🚀 Recent Bug Fixes and Improvements

### Critical Security Fixes
- **Fixed insecure ID generation**: Replaced `math/rand` with `crypto/rand` for secure ID generation
- **Eliminated race conditions**: Fixed concurrent ID generation issues in customer cards, employees, and receipts
- **Added proper error handling**: Improved error handling in update operations to prevent data corruption

### Performance & Reliability Improvements
- **Database connection management**: Added proper connection pooling and graceful shutdown
- **Configurable VAT rates**: Made VAT calculation configurable via environment variables
- **Improved decimal validation**: Enhanced salary and amount validation with proper precision handling
- **Added request ID tracking**: Implemented request tracing for better debugging

### Code Quality Enhancements
- **Centralized utilities**: Created shared utility package for ID generation and validation
- **Middleware support**: Added logging, validation, and error handling middleware
- **Better error messages**: More descriptive error responses throughout the API

## 📋 Features

- **Category Management**: CRUD operations for product categories
- **Product Management**: Complete product lifecycle management
- **Employee Management**: Staff information and role management
- **Customer Cards**: Customer loyalty card system
- **Receipt Management**: Transaction and receipt handling
- **Secure ID Generation**: Cryptographically secure unique identifiers
- **Input Validation**: Comprehensive data validation and sanitization
- **Error Handling**: Robust error handling with detailed logging

## 🛠 Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (>= 1.23.5)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/) (if running without Docker)

### Quick Start with Docker

1. **Clone the repository**:
   ```bash
   git clone git@github.com:velosypedno/zlagoda.git
   cd zlagoda
   ```

2. **Configure environment variables**:
   ```bash
   cp .env.sample .env
   # Edit .env with your configuration
   ```

3. **Start all services**:
   ```bash
   docker compose up
   ```

### Manual Installation

1. **Clone and setup**:
   ```bash
   git clone git@github.com:velosypedno/zlagoda.git
   cd zlagoda
   cp .env.sample .env
   ```

2. **Configure database**:
   - Set `DB_HOST=localhost` in `.env`
   - Configure other database settings as needed

3. **Start database and migrations**:
   ```bash
   docker compose up postgres-zlagoda migrator-zlagoda
   ```

4. **Run the API**:
   ```bash
   go run cmd/api/main.go
   ```

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_DRIVER` | Database driver | `postgres` | Yes |
| `DB_HOST` | Database host | `localhost` | Yes |
| `DB_PORT` | Database port | `5432` | Yes |
| `DB_USER` | Database username | | Yes |
| `DB_PASSWORD` | Database password | | Yes |
| `DB_NAME` | Database name | | Yes |
| `PORT` | Server port | `8080` | Yes |
| `VAT_RATE` | VAT rate (0.2 = 20%) | `0.2` | No |

### Sample Configuration

```env
# Database Configuration
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=zlagoda_user
DB_PASSWORD=secure_password
DB_NAME=zlagoda_db

# Server Configuration
PORT=8080

# Business Configuration
VAT_RATE=0.2

# Optional: Connection Pool Settings
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=5m
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api
```

### Endpoints

#### Categories
- `GET /categories` - List all categories
- `GET /categories/:id` - Get category by ID
- `POST /categories` - Create new category
- `PATCH /categories/:id` - Update category
- `DELETE /categories/:id` - Delete category

#### Products
- `GET /products` - List all products
- `GET /products/:id` - Get product by ID
- `POST /products` - Create new product
- `PATCH /products/:id` - Update product
- `DELETE /products/:id` - Delete product

#### Store Products (Inventory Management)
- `GET /store-products` - List all store products
- `GET /store-products/details` - List store products with product details
- `GET /store-products/promotional` - List promotional products only
- `GET /store-products/by-product/:product_id` - Get store products by product ID
- `GET /store-products/:upc` - Get store product by UPC (12-char)
- `POST /store-products` - Create new store product
- `PATCH /store-products/:upc` - Update store product
- `DELETE /store-products/:upc` - Delete store product
- `PATCH /store-products/:upc/quantity` - Update product quantity
- `GET /store-products/:upc/stock-check` - Check stock availability

#### Sales (Receipt Line Items)
- `GET /sales` - List all sales
- `GET /sales/details` - List sales with product details
- `GET /sales/top-products` - Get top selling products
- `GET /sales/by-receipt/:receipt_number` - Get sales by receipt
- `GET /sales/by-receipt/:receipt_number/details` - Get sales by receipt with details
- `DELETE /sales/by-receipt/:receipt_number` - Delete all sales for receipt
- `GET /sales/by-upc/:upc` - Get sales by UPC
- `GET /sales/stats/product/:product_id` - Get sales statistics for product
- `GET /sales/:upc/:receipt_number` - Get specific sale
- `POST /sales` - Create new sale
- `PATCH /sales/:upc/:receipt_number` - Update sale
- `DELETE /sales/:upc/:receipt_number` - Delete sale

#### Employees
- `GET /employees` - List all employees
- `GET /employees/:id` - Get employee by ID (10-char alphanumeric)
- `POST /employees` - Create new employee
- `PATCH /employees/:id` - Update employee
- `DELETE /employees/:id` - Delete employee

#### Customer Cards
- `GET /customer-cards` - List all customer cards
- `GET /customer-cards/:card_number` - Get card by number (13-char alphanumeric)
- `POST /customer-cards` - Create new customer card
- `PATCH /customer-cards/:card_number` - Update customer card
- `DELETE /customer-cards/:card_number` - Delete customer card

#### Receipts
- `GET /receipts` - List all receipts
- `GET /receipts/:receipt_number` - Get receipt by number (10-char alphanumeric)
- `GET /receipts/:receipt_number/total` - Calculate receipt total from sales
- `POST /receipts` - Create new receipt
- `PATCH /receipts/:receipt_number` - Update receipt
- `DELETE /receipts/:receipt_number` - Delete receipt

### Request/Response Examples

#### Create Employee
```json
POST /api/employees
{
  "empl_surname": "Doe",
  "empl_name": "John",
  "empl_patronymic": "Smith",
  "empl_role": "manager",
  "salary": 50000.00,
  "date_of_birth": "1985-05-15",
  "date_of_start": "2023-01-01",
  "phone_number": "+380123456789",
  "city": "Kyiv",
  "street": "Main Street 123",
  "zip_code": "01001"
}
```

#### Create Receipt
```json
POST /api/receipts
{
  "employee_id": "ABC1234567",
  "card_number": "1234567890123",
  "print_date": "2023-12-01",
  "sum_total": 150.75
}
```

#### Create Store Product
```json
POST /api/store-products
{
  "upc": "123456789012",
  "upc_prom": "123456789013",
  "product_id": 1,
  "selling_price": 15.99,
  "products_number": 100,
  "promotional_product": false
}
```

#### Create Sale
```json
POST /api/sales
{
  "upc": "123456789012",
  "receipt_number": "ABC1234567",
  "product_number": 2,
  "selling_price": 15.99
}
```

#### Check Stock Availability
```
GET /api/store-products/123456789012/stock-check?quantity=5
```

#### Get Sales Statistics
```
GET /api/sales/stats/product/1?start_date=2023-01-01&end_date=2023-12-31
```

## 🔒 Security Features

- **Secure ID Generation**: Uses `crypto/rand` for all ID generation
- **Input Validation**: Comprehensive validation for all inputs
- **SQL Injection Prevention**: Parameterized queries throughout
- **Phone Number Validation**: Ukrainian phone number format validation
- **Decimal Precision**: Proper handling of monetary values
- **CORS Protection**: Configurable CORS settings

## 🏗 Architecture

### Project Structure
```
zlagoda/
├── cmd/api/                 # Application entry point
├── internal/
│   ├── config/             # Configuration management
│   ├── handlers/           # HTTP handlers
│   ├── middleware/         # HTTP middleware
│   ├── models/             # Data models
│   ├── repos/              # Data repositories
│   ├── services/           # Business logic
│   ├── server/             # Server setup
│   ├── utils/              # Shared utilities
│   └── ioc/                # Dependency injection
├── db/migrations/          # Database migrations
├── docker-compose.yml      # Docker configuration
├── .env.sample            # Environment variables template
└── README.md              # This file
```

### Design Patterns
- **Repository Pattern**: Data access abstraction
- **Service Layer**: Business logic separation
- **Dependency Injection**: Loose coupling between components
- **Middleware Pattern**: Cross-cutting concerns handling

## 🧪 Testing

### Run Tests
```bash
go test ./...
```

### Run Tests with Coverage
```bash
go test -cover ./...
```

## 📈 Performance Considerations

- **Connection Pooling**: Configured database connection pool
- **Efficient ID Generation**: Secure random generation with retry logic
- **Proper Error Handling**: Fail-fast approach with detailed logging
- **Graceful Shutdown**: Proper resource cleanup on application termination

## 🔧 Development

### Adding New Features

1. **Model**: Define data structures in `internal/models/`
2. **Repository**: Implement data access in `internal/repos/`
3. **Service**: Add business logic in `internal/services/`
4. **Handler**: Create HTTP handlers in `internal/handlers/`
5. **Routes**: Register routes in `internal/server/`
6. **Wire Up**: Add to IoC container in `internal/ioc/`

### Code Quality

- Follow Go conventions and best practices
- Use meaningful variable and function names
- Add comprehensive error handling
- Include proper logging
- Write tests for new functionality

## 🐛 Known Issues

- None currently reported

## 📝 Changelog

### v1.2.0 (Latest)
- **NEW**: Complete StoreProduct entity implementation for inventory management
- **NEW**: Complete Sale entity implementation for receipt line items
- **NEW**: Advanced analytics endpoints (top products, sales stats)
- **NEW**: Stock management and availability checking
- **NEW**: Promotional product tracking
- Fixed critical security vulnerabilities in ID generation
- Added configurable VAT rates
- Improved error handling and validation
- Added comprehensive middleware support
- Enhanced database connection management
- Added graceful shutdown support

### v1.1.0
- Fixed critical security vulnerabilities in ID generation
- Added configurable VAT rates
- Improved error handling and validation
- Added comprehensive middleware support
- Enhanced database connection management
- Added graceful shutdown support

### v1.0.0
- Initial release with basic CRUD operations
- Docker support
- PostgreSQL integration

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Create an issue on GitHub
- Check the documentation
- Review the code examples

## 🙏 Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- Uses [PostgreSQL](https://www.postgresql.org/) for data persistence
- Containerized with [Docker](https://www.docker.com/)