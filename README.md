# Go Backend API Development

This project contains a set of API endpoints for user registration, login, task management, and more. Below are the details for each API endpoint.

## Base URL
`http://localhost:8080`

## Endpoints

### 1. **Register user**
- **Method**: POST
- **URL**: `/register`
- **Request Body**:
  ```json
  {
      "email": "test@123.com",
      "password": "4234fhjfhef"
  }
  ```
- **Response**:
  - **Code**: 201
  - **Body**:
    ```json
    {
      "message": "User created",
      "user": {
        "id": 1,
        "email": "test@123.com"
      }
    }
    ```

### 2. **Login user**
- **Method**: POST
- **URL**: `/login`
- **Request Body**:
  ```json
  {
      "email": "test@123.com",
      "password": "4234fhjfhef"
  }
  ```
- **Response**:
  - **Code**: 200
  - **Body**:
    ```json
    {
      "message": "Token generated successfully",
      "token": "your_jwt_token"
    }
    ```
- **Test Script**: 
  The JWT token is saved in the environment variable `jwt_token`.
  ```javascript
  pm.environment.set("jwt_token", pm.response.json().token);
  ```

### 3. **Create task**
- **Method**: POST
- **URL**: `/tasks`
- **Authorization**: Bearer token
- **Request Body**:
  ```json
  {
      "title": "my random task",
      "description": "some random description",
      "due_date": "10/01/2025 15:40",
      "completed": true
  }
  ```
- **Response**:
  - **Code**: 201
  - **Body**:
    ```json
    {
      "message": "task created successfully",
      "task": {
        "id": 17,
        "title": "my random task",
        "description": "some random description",
        "due_date": "2025-01-10T15:40:00Z",
        "completed": true,
        "user_id": 1
      }
    }
    ```

### 4. **Update task**
- **Method**: PUT
- **URL**: `/tasks/{task_id}`
- **Authorization**: Bearer token
- **Request Body**:
  ```json
  {
      "title": "my first task (edited)",
      "description": "some random description (edited)",
      "due_date": "12/01/2025",
      "completed": false
  }
  ```
- **Response**:
  - **Code**: 200
  - **Body**:
    ```json
    {
      "message": "task updated successfully",
      "task": {
        "id": 3,
        "title": "my first task (edited)",
        "description": "some random description (edited)",
        "due_date": "12/01/2025",
        "completed": false,
        "user_id": 1
      }
    }
    ```

### 5. **Delete task**
- **Method**: DELETE
- **URL**: `/tasks/{task_id}`
- **Authorization**: Bearer token
- **Response**:
  - **Code**: 200
  - **Body**:
    ```json
    {
      "message": "task deleted successfully"
    }
    ```

### 6. **Get all tasks**
- **Method**: GET
- **URL**: `/tasks?page=1&order=desc&sortBy=due_date`
- **Authorization**: Bearer token
- **Response**:
  - **Code**: 200
  - **Body**:
    ```json
    {
      "message": "tasks retrieved successfully",
      "tasks": [
        {
          "id": 17,
          "title": "my random task",
          "description": "some random description",
          "due_date": "2025-01-10T15:40:00Z",
          "completed": true,
          "user_id": 1
        },
        {
          "id": 16,
          "title": "my random task",
          "description": "some random description",
          "due_date": "2025-01-10T00:00:00Z",
          "completed": true,
          "user_id": 1
        }
      ]
    }
    ```

### 7. **Get one task by ID**
- **Method**: GET
- **URL**: `/tasks/{task_id}`
- **Authorization**: Bearer token
- **Response**:
  - **Code**: 200
  - **Body**:
    ```json
    {
      "id": 3,
      "title": "my first task (edited)",
      "description": "some random description (edited)",
      "due_date": "12/01/2025",
      "completed": false,
      "user_id": 1
    }
    ```

## Authentication

All endpoints (except registration and login) require a JWT token. You can get the token by logging in with valid user credentials. Once logged in, the token will be stored in the environment variable `jwt_token` for further requests.

## Local Development

1. Clone this repository.
2. Run the backend service:
   - Install dependencies: `go mod tidy`
   - Run the service: `go run main.go`
3. Use Postman or cURL to test the API endpoints as described above.
