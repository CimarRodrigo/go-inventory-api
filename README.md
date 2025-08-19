# API Project

A clean architecture Go API built with Gin framework.

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/CimarRodrigo/go-inventory-api
   cd go-inventory-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

The server will start by default on `http://localhost:8080`

## 📋 Environment Configuration

Create a `.env` file in the root directory (optional, defaults will be used):

```bash
# Server Configuration
PORT=8080
GIN_MODE=debug  # debug, release, test

# Application Configuration
APP_NAME=My API
APP_VERSION=1.0.0
APP_ENV=development  # development, staging, production
```

### Environment Variables

| Variable | Description | Default | Options |
|----------|-------------|---------|---------|
| `PORT` | Server port | `8080` | Any valid port number |
| `GIN_MODE` | Gin framework mode | `debug` | `debug`, `release`, `test` |
| `APP_NAME` | Application name | `My API` | Any string |
| `APP_VERSION` | Application version | `1.0.0` | Semantic versioning |
| `APP_ENV` | Environment name | `development` | `development`, `staging`, `production` |

## 🛣️ API Endpoints

### Health Checks
- `GET /api/v1/health` - Basic health check
