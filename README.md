# Machine Learning Model Deployment with Go and Docker

This project demonstrates how to build, containerize, and deploy a simple machine learning model using Go and Docker. We trained a simple model, exposed it via an HTTP API, and deployed the whole application inside a Docker container for consistent and portable execution across different environments.

## Table of Contents

- [Overview](#overview)
- [Technologies Used](#technologies-used)
- [Project Setup](#project-setup)
  - [Step 1: Writing the Model Logic in Go](#step-1-writing-the-model-logic-in-go)
  - [Step 2: Dockerizing the Application](#step-2-dockerizing-the-application)
  - [Step 3: Building and Running the Docker Container](#step-3-building-and-running-the-docker-container)
  - [Step 4: Exposing the Model via an HTTP API](#step-4-exposing-the-model-via-an-http-api)
  - [Step 5: Pushing to DockerHub](#step-5-pushing-to-dockerhub)
- [Running the Application Locally](#running-the-application-locally)
- [Sending Requests to the API](#sending-requests-to-the-api)

## Overview

This project is an example of how to deploy a simple machine learning model using the Go programming language and Docker. The key objectives were to:
- Build a lightweight Go application that simulates a machine learning model.
- Containerize the application using Docker to ensure portability and consistent execution across various environments.
- Expose the prediction logic via an HTTP API for easy integration with other services.

## Technologies Used

- **Go (Golang)**: Programming language used to write the model logic and expose it via an HTTP API.
- **Docker**: Used to containerize the Go application for easy deployment and portability.
- **Curl**: To send test requests to the HTTP API and validate predictions.

## Project Setup

### Step 1: Writing the Model Logic in Go

We wrote a simple Go application that simulates training a machine learning model and predicting output based on input data. The model logic is encapsulated in the `model.go` file.

Key parts of the Go code:
- **`trainModel()`**: Simulates model training.
- **`predict()`**: Generates a random prediction for the given input.
- **HTTP Server**: The model’s prediction logic is exposed via an HTTP POST API at `/predict`.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
)

type Prediction struct {
    Input  []float64 `json:"input"`
    Output int       `json:"output"`
}

func trainModel() {
    fmt.Println("Model is being trained...")
    time.Sleep(2 * time.Second) 
    fmt.Println("Model trained and ready!")
}

func predict(inputData []float64) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(3) 
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
    var input []float64
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    output := predict(input)
    response := Prediction{Input: input, Output: output}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    trainModel()
    http.HandleFunc("/predict", predictHandler)
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Step 2: Dockerizing the Application

We created a `Dockerfile` to package the Go application and its dependencies into a Docker container. This makes the application portable and ensures consistent execution across different environments.

```dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod download
RUN go build -o model-app .
EXPOSE 8080
CMD ["/app/model-app"]
```

### Step 3: Building and Running the Docker Container

To build and run the Docker container locally, follow these steps:

1. **Build the Docker image**:
   ```bash
   docker build -t mesutoezdil/machine-learning-model-logic-in-go .
   ```

2. **Run the Docker container**:
   ```bash
   docker run -p 8080:8080 mesutoezdil/machine-learning-model-logic-in-go
   ```

### Step 4: Exposing the Model via an HTTP API

We exposed the prediction logic via an HTTP API using Go's `net/http` package. This allows us to send POST requests to the `/predict` endpoint, and the server responds with predictions.

### Step 5: Pushing to DockerHub

We pushed the Docker image to DockerHub to make it accessible to others. The image can be pulled and run from DockerHub.

1. **Tag the image**:
   ```bash
   docker tag mesutoezdil/machine-learning-model-logic-in-go mesutoezdil/machine-learning-model-logic-in-go
   ```

2. **Push the image to DockerHub**:
   ```bash
   docker push mesutoezdil/machine-learning-model-logic-in-go
   ```

## Running the Application Locally

1. **Pull the image from DockerHub**:
   ```bash
   docker pull mesutoezdil/machine-learning-model-logic-in-go
   ```

2. **Run the container**:
   ```bash
   docker run -p 8080:8080 mesutoezdil/machine-learning-model-logic-in-go
   ```

3. The server will start and listen for requests on port `8080`.

## Sending Requests to the API

You can test the model’s prediction by sending a POST request to the `/predict` endpoint. Here’s an example using `curl`:

```bash
curl -X POST http://localhost:8080/predict -d '[5.1, 3.5, 1.4, 0.2]' -H "Content-Type: application/json"
```

The server will respond with a JSON object containing the input and the predicted output, like this:

```json
{
  "input": [5.1, 3.5, 1.4, 0.2],
  "output": 1
}
```

## Conclusion

This project shows how to containerize and expose a simple machine learning model using Go and Docker. The API provides a way to send requests and receive predictions, making the model easy to integrate into other applications.
