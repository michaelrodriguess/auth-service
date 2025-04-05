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
│   ├── model/              # Domain entities and business rules
│   ├── service/            # Business logic
│   ├── repository/         # Persistence layer
│   └── handler/           # Endpoints: login, register, validate
├── pkg/                    # Public code, shared libraries
│   ├── auth/               # Authentication/authorization logic
    ├── db/                 # Database logic
│   ├── jwt/                # JWT implementation
│   └── crypto/             # Cryptographic functions
├── config/                 # Application configuration

