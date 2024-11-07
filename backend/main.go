package main

import (
    "fmt"
    "log"
    "net/http"
    "F1DataVisualizer/backend/internal/handlers" 
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    
  r := chi.NewRouter()

  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  
  r.Get("/seasons", handlers.SeasonsDataHandler)
  r.Get("/circuits", handlers.CircuitsDataHandler)
  r.Get("/races/{year}", handlers.RacesDataHandler)
  r.Get("/constructors/{year}", handlers.ConstructorsDataHandler)
  r.Get("/drivers/{year}", handlers.DriversDataHandler)
  r.Get("/results/{year}", handlers.ResultsDataHandler)
  // r.Get("/sprint/{year}", handlers.DriversDataHandler)
  // r.Get("/qualifying/{year}", handlers.DriversDataHandler)
  // r.Get("/pitstops/{year}/{placeholder_num}", handlers.DriversDataHandler)
  // r.Get("/laps/{year}/{placeholder_num}", handlers.DriversDataHandler)
  // r.Get("/driverstandings/{year}", handlers.DriversDataHandler)
  // r.Get("/constructorstandings/{year}", handlers.DriversDataHandler)



  // Start the HTTP server on localhost:8080
  fmt.Println("Server starting on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}


