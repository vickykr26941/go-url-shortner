## URL Shortener Service (Under Development)
This project is a self learning exercise in Go, Gin, Mysql,PostgreSQL,and Redis. It is a URL shortener service that allows users to shorten long URLs, track analytics, and manage user accounts. The project is structured to follow best practices in Go development, including modular design, dependency management, and testing.

#### Go Learning (got to learn)
1. Gin
2. Wire
3. Cobra
4. Postgres
5. Mysql
6. Go Cuncurrency
7. Go package architecture

![Go](https://img.shields.io/badge/Go-1.24-blue.svg)
This is a URL shortener service built with Go, PostgreSQL, and Redis. It provides features such as URL shortening, analytics, user authentication, and caching.
## Features
- URL shortening with custom aliases
- Analytics for shortened URLs
- User authentication and management
- Rate limiting
- Caching with Redis
- Structured logging
- API documentation
- Configuration management
- Middleware for authentication, rate limiting, and logging
- Multi-environment configuration (development, production)
- Support for multiple database config(postgres, mysql etc)
- Support for custom error handling
- Support for custom URL encoding/decoding


```azure
url-shortener/
├── cmd/
│   ├── server/
│   │   └── main.go                 # Application entry point
│   └── migrate/
│       └── main.go                 # Database migration tool
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── database/
│   │   ├── postgres.go             # PostgreSQL connection
│   │   ├── redis.go                # Redis connection
│   │   └── migrations/
│   │       ├── 001_create_urls.sql
│   │       └── 002_create_analytics.sql
│   ├── models/
│   │   ├── url.go                  # URL model
│   │   ├── analytics.go            # Analytics model
│   │   └── user.go                 # User model
│   ├── repository/
│   │   ├── interfaces.go           # Repository interfaces
│   │   ├── url_repo.go             # URL repository
│   │   ├── analytics_repo.go       # Analytics repository
│   │   └── user_repo.go            # User repository
│   ├── service/
│   │   ├── url_service.go          # URL business logic
│   │   ├── analytics_service.go    # Analytics business logic
│   │   ├── auth_service.go         # Authentication service
│   │   └── cache_service.go        # Caching service
│   ├── handler/
│   │   ├── url_handler.go          # URL HTTP handlers
│   │   ├── analytics_handler.go    # Analytics HTTP handlers
│   │   ├── auth_handler.go         # Authentication handlers
│   │   └── middleware/
│   │       ├── auth.go             # Authentication middleware
│   │       ├── rate_limit.go       # Rate limiting middleware
│   │       └── logging.go          # Logging middleware
│   ├── utils/
│   │   ├── hash.go                 # URL encoding/decoding
│   │   ├── validator.go            # Input validation
│   │   └── response.go             # HTTP response helpers
│   └── router/
│       └── router.go               # Route definitions
├── pkg/
│   ├── logger/
│   │   └── logger.go               # Structured logging
│   └── metrics/
│       └── metrics.go              # Prometheus metrics
├── web/
│   ├── templates/
│   │   ├── index.html              # Landing page
│   │   ├── dashboard.html          # User dashboard
│   │   └── analytics.html          # Analytics page
│   └── static/
│       ├── css/
│       ├── js/
│       └── images/
├── scripts/
│   ├── build.sh                    # Build script
│   ├── test.sh                     # Test script
│   └── deploy.sh                   # Deployment script
├── docker/
│   ├── Dockerfile                  # Multi-stage Docker build
│   ├── docker-compose.yml          # Local development setup
│   └── docker-compose.prod.yml     # Production setup
├── k8s/
│   ├── deployment.yaml             # Kubernetes deployment
│   ├── service.yaml                # Kubernetes service
│   ├── configmap.yaml              # Configuration
│   └── ingress.yaml                # Ingress configuration
├── tests/
│   ├── integration/
│   │   └── url_test.go
│   ├── unit/
│   │   ├── service_test.go
│   │   └── handler_test.go
│   └── fixtures/
│       └── test_data.sql
├── configs/
│   ├── config.yaml                 # Default configuration
│   ├── config.dev.yaml             # Development config
│   └── config.prod.yaml            # Production config
├── docs/
│   ├── API.md                      # API documentation
│   ├── ARCHITECTURE.md             # Architecture overview
│   └── DEPLOYMENT.md               # Deployment guide
├── .github/
│   └── workflows/
│       ├── ci.yml                  # CI pipeline
│       └── cd.yml                  # CD pipeline
├── go.mod
├── go.sum
├── Makefile                        # Build automation
├── README.md
└── .env.example                    # Environment variables template
```
