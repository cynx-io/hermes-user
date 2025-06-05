# Hermes User Service

A gRPC-based user management service written in Go.

## Features

- User management (create, get, list)
- Username availability checking
- Pagination support
- Standardized response format
- gRPC communication

## Prerequisites

- Go 1.21 or higher
- Protocol Buffers compiler (protoc)
- MySQL/MariaDB

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/hermes-user.git
cd hermes-user
```

2. Install dependencies:
```bash
go mod download
```

3. Generate proto files:
```bash
make proto
```

4. Build the application:
```bash
make build
```

## Configuration

Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
GRPC_PORT=50051
```

## Running the Service

```bash
make run
```

## API Documentation

### gRPC Endpoints

1. CheckUsername
   - Request: `username`
   - Response: `exists` (boolean)

2. GetUser
   - Request: `username`
   - Response: User details (id, username, coin, timestamps)

3. CreateUser
   - Request: `username`, `password`
   - Response: Created user details

4. PaginateUsers
   - Request: `page`, `limit`, `sort_by`, `sort_order`
   - Response: List of users with pagination info

## Response Codes

- `00`: Success
- `DBU`: Database error
- `NFU`: Not found
- `VDU`: Validation error

## Development

1. Generate proto files:
```bash
make proto
```

2. Build the application:
```bash
make build
```

3. Run the application:
```bash
make run
```

4. Clean generated files:
```bash
make clean
```

## License

MIT License