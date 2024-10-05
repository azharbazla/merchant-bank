# merchant-bank API

API between merchant & bank using Golang.

## How to Run
#### Install dependencies:
  ```bash
  go mod tidy
  ```
#### Run the server:
  ```bash
  go run main.go
  ```

## API Endpoint

#### Customer Login

- **POST** `/login`

  Request Body :

  ```json
  {
    "email": "customer@example.com",
    "password": "password123"
  }
  ```

#### Payment

- **POST** `/payment`

  Request Body :

  ```json
  {
    "amount": 100.0
  }
  ```

#### Logout

- **POST** `/logout`

#### Fetch All History

- **POST** `/history`

#### Fetch Customer History

- **POST** `/history/customer`

## Sample Requests
#### Login Request
  ```bash
  curl -X POST http://localhost:8080/login -d '{"email":"test@test.com", "password":"1234"}' -H "Content-Type: application/json"
  ```
#### Payment Request
  ```bash
  curl -X POST http://localhost:8080/payment -d '{"amount": 100.0}' -H "Content-Type: application/json" -H "Authorization: Bearer <JWT>"
  ```
#### Fetch All History Request
  ```bash
  curl -X GET http://localhost:8080/history
  ```
#### Fetch Customer History Request
  ```bash
  curl -X GET http://localhost:8080/history/customer -H "Authorization: Bearer <JWT>"
  ```
## Testing the Application
#### Run the tests using Go’s testing framework:
  ```bash
  go test ./tests
  ```