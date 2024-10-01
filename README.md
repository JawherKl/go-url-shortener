## URL Shortener in Go

This project is a simple URL shortener built in Go. It allows users to shorten long URLs and provides redirection to the original URLs when the shortened URL is visited.

### Features
- Generate a short URL for any long URL.
- Redirect users to the original URL using the short URL.
- In-memory URL store (no persistence across restarts).

### Project Structure
- `main.go`: The entry point of the application that sets up the HTTP server and routes.
- `store.go`: Contains logic for the in-memory URL storage and URL generation.
- `handlers.go`: Contains the HTTP handler functions for shortening and redirecting URLs.

### Getting Started

#### Prerequisites
- [Go](https://golang.org/doc/install) installed on your machine.

#### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/go-url-shortener.git
    cd go-url-shortener
    ```

2. Run the application:
    ```bash
    go run *.go
    ```

3. The server will start on `localhost:8080`.

### API Endpoints

#### 1. Shorten URL
- **Endpoint**: `/shorten`
- **Method**: `POST`
- **Request Body**: JSON with the original URL:
    ```json
    {
      "url": "https://www.example.com"
    }
    ```

- **Response**: JSON with the shortened URL:
    ```json
    {
      "short_url": "abc123"
    }
    ```

#### Example:
```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://www.example.com"}'


#### 2. Redirect to Original URL
- **Endpoint**: `/r/{short_url}
- **Method**: `GET`
- **Response**: Redirects to the original URL

#### Example:
```bash
http://localhost:8080/r/abc123

### Future Enhancements
- Persistent storage (e.g., using a database).
- Custom short URL creation.
- URL expiration feature.
