# RESTful API Go Fiber

A RESTful API project using Go, Fiber, and PostgreSQL. This project provides a clean, modular structure for building scalable APIs with built-in features such as request validation, error handling, logging, and testing.

## Quick Start

To create a project, simply run:

```bash
go mod init <project-name>
```

## Manual Installation

Clone the repo:

```bash
git clone https://github.com/gnafhan/intern-go
cd <intern.go>
```

Install the dependencies:

```bash
go mod tidy
```

Set the environment variables:

```bash
cp .env.example .env
# open .env and modify the environment variables (if needed)
```

## Features

- **SQL database**: PostgreSQL with Gorm ORM
- **Database migrations**: with golang-migrate
- **Validation**: request data validation using go-playground/validator
- **Logging**: using Logrus and Fiber-Logger
- **Testing**: unit and integration tests with Testify
- **Error handling**: centralized error handling mechanism
- **Environment variables**: using Viper
- **Security**: HTTP headers with Fiber-Helmet
- **CORS**: enabled with Fiber-CORS
- **Compression**: gzip compression with Fiber-Compress
- **Docker support**
- **Linting**: with golangci-lint

## Commands

Running locally:

```bash
make start
```

Or running with live reload:

```bash
air
```

Testing:

```bash
make tests
make testsum
make tests-TestTaskService
```

Docker:

```bash
make docker
make docker-test
```

Linting:

```bash
make lint
```

## Environment Variables

See `.env.example` for all configuration options.

## Project Structure

```
src\
 |--config\         # Environment variables and configuration
 |--controller\     # Route controllers
 |--database\       # Database connection & migrations
 |--docs\           # Swagger files
 |--middleware\     # Custom fiber middlewares
 |--model\          # Postgres models
 |--response\       # Response models
 |--router\         # Routes
 |--service\        # Business logic
 |--utils\          # Utility classes and functions
 |--validation\     # Request data validation schemas
 |--main.go         # Fiber app
```

## API Documentation

To view the list of available APIs and their specifications, run the server and go to `http://localhost:3000/v1/docs` in your browser.


## API Endpoints

List of available routes:

**Category routes**:
- `POST /v1/categories` - create a category
- `GET /v1/categories` - get all categories
- `GET /v1/categories/:categoryId` - get category by id
- `PUT /v1/categories/:categoryId` - update category
- `DELETE /v1/categories/:categoryId` - delete category

**Task routes**:
- `POST /v1/tasks` - create a task
- `GET /v1/tasks` - get all tasks
- `GET /v1/tasks/:taskId` - get task by id
- `PUT /v1/tasks/:taskId` - update task
- `DELETE /v1/tasks/:taskId` - delete task

## Error Handling

The app includes a custom error handling mechanism, which can be found in the `src/utils/error.go` file.

The error handling process sends an error response in the following format:

```json
{
  "code": 404,
  "status": "error",
  "message": "Not found"
}
```

## Validation

Request data is validated using [Package validator](https://github.com/go-playground/validator). Validation schemas are defined in the `src/validation` directory and are used within the services by passing them to the validation logic.

## Logging

Logging is done using Logrus. See `src/utils/logrus.go` for details.

## Linting

Linting is done using [golangci-lint](https://golangci-lint.run)

## Contributing

Contributions are welcome! Please check out the [contributing guide](CONTRIBUTING.md).

## License

[MIT](LICENSE)
