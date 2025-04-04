# Auth Service 
An authentication service developed in Go to provide authentication and authorization functionality for applications.

## Features  
- [ ] **User authentication**  
- [ ] **Session management**  
- [ ] **JWT (JSON Web Token) for authentication**  
- [ ] **Credential validation**  
- [ ] **Token management**  

## Project Structure  

```plaintext
auth-service/
├── cmd/                    # Command line interface
│   └── server/             # Server HTTP
│       └── main.go
├── internal/               # Private code
│   ├── domain/             # Domain entities and business rules
│   │   ├── model/          # Domain models (User, Role, etc.)
│   │   └── service/        # Business logic
│   ├── repository/         # Persistence layer
│   ├── usecase/            # Application use cases
│   └── delivery/           # Presentation layer
│       └── http/           # HTTP handlers
├── pkg/                    # Public code, shared libraries
│   ├── auth/               # Authentication/authorization logic
│   ├── jwt/                # JWT implementation
│   └── crypto/             # Cryptographic functions
├── config/                 # Application configuration
└── migrations/             # Database migrations
