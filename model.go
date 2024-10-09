package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
)

// Prediction represents the input and output data
type Prediction struct {
    Input  []float64 `json:"input"`
    Output int       `json:"output"`
}

// Simulating model training
func trainModel() {
    fmt.Println("Model is being trained...")
    time.Sleep(2 * time.Second) // Simulate training time
    fmt.Println("Model trained and ready!")
}

// Simulate prediction
func predict(inputData []float64) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(3) // Simulate prediction (classification result between 0 and 2)
}

// HTTP handler for prediction
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

    // Set up the HTTP server
    http.HandleFunc("/predict", predictHandler)
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
