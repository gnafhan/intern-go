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

## Deployment on Ubuntu VPS

To deploy this application on an Ubuntu VPS, follow these steps:

1. **Run as a System Service**
   - Use `make start` to run the application. For production, it is recommended to run the app as a systemd service to ensure it stays running and restarts automatically if it crashes.
   - Example systemd service file (`/etc/systemd/system/intern-go.service`):
     ```ini
     [Unit]
     Description=Intern Go Fiber API
     After=network.target

     [Service]
     User=ubuntu
     WorkingDirectory=/path/to/your/project
     ExecStart=/usr/bin/make start
     Restart=always
     Environment=PATH=/usr/local/bin:/usr/bin:/bin
     Environment=GO_ENV=production

     [Install]
     WantedBy=multi-user.target
     ```
   - Reload systemd and start the service:
     ```bash
     sudo systemctl daemon-reload
     sudo systemctl enable intern-go
     sudo systemctl start intern-go
     ```

2. **Set Up Nginx Reverse Proxy**
   - Install nginx:
     ```bash
     sudo apt update && sudo apt install nginx
     ```
   - Configure nginx to proxy requests to your Go app (listening on port 3000):
     ```nginx
     server {
         listen 80;
         server_name your-domain.com;

         location / {
             proxy_pass http://localhost:3000;
             proxy_set_header Host $host;
             proxy_set_header X-Real-IP $remote_addr;
             proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
             proxy_set_header X-Forwarded-Proto $scheme;
         }
     }
     ```
   - Test and reload nginx:
     ```bash
     sudo nginx -t
     sudo systemctl reload nginx
     ```

3. **Enable Free SSL with Certbot**
   - Install Certbot:
     ```bash
     sudo apt install certbot python3-certbot-nginx
     ```
   - Obtain and install SSL certificate:
     ```bash
     sudo certbot --nginx -d your-domain.com
     ```
   - Certbot will automatically configure SSL in your nginx config and set up auto-renewal.

4. **Deployment Link**
   - The API documentation is deployed and available at: [https://intern-go.nafhan.com/v1/docs/index.html](https://intern-go.nafhan.com/v1/docs/index.html)

---

*Deployment instructions are provided for running the API securely and reliably in a production environment using Ubuntu VPS, systemd, nginx, and free SSL from Let's Encrypt.*
