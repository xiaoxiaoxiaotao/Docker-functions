### README.md

# DOCKER-Longest Common Subsequence

This project contains Go code for Longest Common Subsequence algorithms and a simple web server implemented using the Gin framework. The project is containerized using Docker for easy deployment and management.

## Project Structure

```
DOCKER_GIN_FUNCTIONS/
│
├── functions/
│   ├── LCS.go             # Implementation of Longest Common Subsequence algorithm
│
├── Dockerfile             # Dockerfile for containerizing the Go application
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies file
├── main.go                # Main entry point of the Go application
└── README.md              # Project documentation
```

## Algorithms

### Longest Common Subsequence (LCS)

The `LCS.go` file contains the implementation of the Longest Common Subsequence algorithm, which finds the longest subsequence common to two sequences.

## Web Server

The `app.go` file contains a simple web server built using the Gin framework. The server listens on port 5000 and includes an example endpoint.

## Docker

The `Dockerfile` is used to build a Docker image for the Go application. It installs the necessary dependencies, copies the source code, and sets the entry point for the container.

## Prerequisites

- Docker

## Building and Running the Docker Container

1. **Build and Run the Docker image:**

    You can simply run the script `shell.sh`:

    ```sh
      ./shell.sh
    ```

    Or you can run the docker with following step:

    First build the docker image.
    ```sh
    docker build -t my-go-app:v1.0 .
    ```

    ```sh
    docker run -p 5000:5000 --name my-gin-container -v "$(pwd):/Docker_gin_functions" my-go-app:v1.0
    ```

    This command will start the container and map port 5000 of the host to port 5000 of the container.

    Or you can enter the docker with interactive mode with `-i` mode to manipulate the code. For instance:

    ```sh
    docker run -it --name my-go-container -p 5000:5000  go_gin_functions:v1.0
    go run app.go
    ```

2. **Access the application:**
Open your browser and go to `http://localhost:5000/`. You should see the web server running and responding to requests. You can also access it with the following cmd.
```bash
curl -X GET http://localhost:5000/
```

### Testing APIs

You can test the APIs using CURL commands. Here are examples of how to test each endpoint:

#### 1. Health Check API

```bash
curl -X GET http://localhost:5000/api/health
```

#### 2. Shutdown API

```bash
curl -X POST http://localhost:5000/api/shutdown
```

#### 3. LCS Calculation API

```bash
# Test with length_only set to true
curl -X POST http://localhost:5000/api \
  -H 'Content-Type: application/json' \
  -d '{
    "root": {
        "length_only": false,
        "str1": "abcdef",
        "str2": "abdf"
      }
  }'


# Test with length_only set to false
curl -X POST \
  http://localhost:5000/api \
  -H 'Content-Type: application/json' \
  -d '{
    "root": 
      {
        "length_only": false,
        "str1": "abcdef",
        "str2": "abdf"
      }
  }'
```

### API Responses

- **Health Check**: Returns a JSON response with "result: success" if the server is running.
- **Shutdown**: Gracefully shuts down the server and returns "result: success" upon completion.
- **LCS Calculation**:
  - For `length_only` set to true, returns the length of the LCS.
  - For `length_only` set to false, returns both the LCS and its length.


## Stopping and Removing the Docker Container

1. **Stop the container:**

    ```sh
    docker stop my-go-container
    ```

2. **Remove the container:**

    ```sh
    docker rm my-go-container
    ```

## Example Commands for Docker Management

- **List all containers:**

    ```sh
    docker ps -a
    ```

- **View container logs:**

    ```sh
    docker logs my-go-container
    ```

- **Enter a running container:**

    ```sh
    docker exec -it my-go-container
    ```