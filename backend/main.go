package main

import (
    "fmt"
    "log"
    "net/http"
    "F1DataVisualizer/backend/internal/handlers" 
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

func main() {
    
  r := chi.NewRouter()

  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  r.Use(cors.Handler(cors.Options{
      AllowedOrigins:   []string{"http://localhost:3001"}, // Allow only your frontend origin
      AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
      AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
      AllowCredentials: true,
  }))
  
  r.Get("/seasons", handlers.SeasonsDataHandler)
  r.Get("/circuits", handlers.CircuitsDataHandler)
  r.Get("/races/{year}", handlers.RacesDataHandler)
  r.Get("/constructors/{year}", handlers.ConstructorsDataHandler)
  r.Get("/drivers/{year}", handlers.DriversDataHandler)
  r.Get("/results/{year}", handlers.ResultsDataHandler)
  r.Get("/sprint/{year}", handlers.SprintDataHandler)
  r.Get("/qualifying/{year}", handlers.QualifyingDataHandler)
  // r.Get("/pitstops/{year}/{placeholder_num}", handlers.DriversDataHandler)
  // r.Get("/laps/{year}/{placeholder_num}", handlers.DriversDataHandler)
  r.Get("/driverstandings/{year}", handlers.DriverStandingsDataHandler)
  r.Get("/constructorstandings/{year}", handlers.ConstructorStandingsDataHandler)



  // Start the HTTP server on localhost:8080
  fmt.Println("Server starting on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}


