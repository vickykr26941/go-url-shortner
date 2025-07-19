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