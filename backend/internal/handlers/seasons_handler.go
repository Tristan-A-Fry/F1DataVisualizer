package handlers

/*
Fucntionality of this handler:

HTTP handler that processes a req to the /seasons endpoint 

We are getting the seasons data from the external api, in this case it will return the format:
  
  0:
    season "1950"
    url: ""http://en.wikipedia.org/wiki/1950_Formula_One_season""

  all the way to season 75 

We do this to allow a user to be able to select the season they want to visualize on the front end.
*/



import(
  "encoding/json" //allows us to decode and encode json data i.e convert Go data structures to json
  "net/http" //http handling
  "F1DataVisualizer/backend/internal/services"
)

/*
  w http.ResponseWriter allows us to write the http response back to the client
  r *http.Request represents the incoming http req, contains all info
*/

func SeasonsDataHandler(w http.ResponseWriter, r *http.Request) {
    /* 
      Fetch all seasons data from the API
      data -> data retireved from the api 
      err -> if nil, then we are successful
    */

    data, err := services.GetSeasonsData()
    if err != nil {
        http.Error(w, "Failed to fetch season data", http.StatusInternalServerError)
        return
    }

    // Set Content-Type to application/json
    // This header informs the client that the response body contains JSON data 
    w.Header().Set("Content-Type", "application/json")

    // Encode data to JSON and send it as the HTTP response
    /*
      json.NewEncoder(w).Encode(data), encodes the data as JSON and writes it ResponseWriter w
      We check if there is an error during the encoding process  
    */
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Failed to encode data", http.StatusInternalServerError)
        return
    }
}
