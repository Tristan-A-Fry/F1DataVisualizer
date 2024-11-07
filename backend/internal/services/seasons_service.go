package services

import (
    "encoding/json"
    "fmt" //format for I/O i.e strings and printed logs
    "io/ioutil" //used for reading data from an io.Reader to read the response body
    "net/http"
    "net/url" //used to parse and manipulate URLs
    "time"
    "F1DataVisualizer/backend/config" //geting the APIBaseURL from the config file 
)

/*
  Interface allows us to return any data type in this case JSON structure
*/

func GetSeasonsData() (interface{}, error) {
    /*  
      This constructs the full api url by appending /seaons.
      fmt.Sprintf formats the string.
      %w wraps the error.
    */
    apiURL, err := url.Parse(fmt.Sprintf("%s/seasons", config.APIBaseURL))
    if err != nil {
        return nil, fmt.Errorf("failed to parse base URL: %w", err)
    }

    /*
      .Query() -> extracts the existing query parameters from the apiUrl
      query.Set("limit", "75") -> in the third party api, by default the seasons endpoint only returns the first 30 seasons, so we use their limit funcitonality to return the whole list of seasons.

       apiURL.RawQuery = query.Encode() encodes query parameters and sets them back into the url obj, ensuring the url includes the 75 limit 
    */
    query := apiURL.Query()
    query.Set("limit", "75")
    apiURL.RawQuery = query.Encode()

    // Log the final URL for debugging
    fmt.Println("Final URL:", apiURL.String())

    // Create the request, nil indicates that there is no req body 
    req, err := http.NewRequest("GET", apiURL.String(), nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }
    
    // creates client with timout of 10 secs so that the req will be canceled if the req last longer than 10 sec
    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to make request: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %w", err)
    }

    var data interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return data, nil
}


