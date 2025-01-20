This project is a template with a Clean Architecture, written in the Go programming language, designed to provide the foundational capabilities for developing software. My goal with this project is to implement a comprehensive set of essential features required for a backend project. This project includes the following:

Infrastructure for developing RESTful services with documentation support using Swagger
Token-based authentication via JWT, including user roles
Dependency injection infrastructure using the Dig framework
Support for connecting to four powerful databases: MongoDB, PostgreSQL, MSSQL, and Neo4J, using the Repository Pattern
Multilingual error management and message generation system
Other foundational features such as Jalali and Gregorian date management, list handling, and more.

The project structure is based on Clean Architecture and follows the hierarchy shown below.

GoClean
├── GoClean.Api
│   ├── docs
│   └── Middlewares
├── GoClean.Application
│   ├── Commons
│   ├── Contract
│   ├── Database
│   │   ├── MongoDB
│   │   ├── Neo4J
│   │   └── Sql
│   ├── Services
│   └── Features
│       ├── Api
│       │   └── Sample
│       └── Model
├── GoClean.Common
├── GoClean.Domain
│   ├── Entities
│   │   ├── MongoDB
│   │   ├── Neo4J
│   │   └── Sql
│   └── Enums
├── GoClean.Identity
├── GoClean.Infrastructure
│   ├── Database
│   │   ├── MongoDB
│   │   ├── Neo4J
│   │   └── Sql
│   ├── ServiceImpl
│   │   └── SMSIR
│   └── Services

