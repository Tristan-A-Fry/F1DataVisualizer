

package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
    "F1DataVisualizer/backend/config"
)

func GetResultsData(year string, limit string) (interface{}, error) {
    if limit == "" {
        limit = "100" // Use the maximum allowed by the API
    }

    offset := 0
    allResults := []interface{}{}
    client := &http.Client{Timeout: 10 * time.Second}
    totalResults := -1

    for {
        apiURL := fmt.Sprintf("%s/%s/results?limit=%s&offset=%d", config.APIBaseURL, year, limit, offset)
        fmt.Println("Requesting URL:", apiURL)

        req, err := http.NewRequest("GET", apiURL, nil)
        if err != nil {
            return nil, fmt.Errorf("failed to create request: %w", err)
        }

        resp, err := client.Do(req)
        if err != nil {
            return nil, fmt.Errorf("failed to make request: %w", err)
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, fmt.Errorf("failed to read response: %w", err)
        }

        var data map[string]interface{}
        if err := json.Unmarshal(body, &data); err != nil {
            return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
        }

        races := data["MRData"].(map[string]interface{})["RaceTable"].(map[string]interface{})["Races"].([]interface{})
        allResults = append(allResults, races...)

        // Safely extract and convert the total number of results
        totalField := data["MRData"].(map[string]interface{})["total"]
        switch v := totalField.(type) {
        case float64:
            totalResults = int(v)
        case string:
            totalResults, err = strconv.Atoi(v)
            if err != nil {
                return nil, fmt.Errorf("failed to convert total to int: %w", err)
            }
        default:
            return nil, fmt.Errorf("unexpected type for total field: %T", v)
        }

        fmt.Println("Total results available:", totalResults)

        // If the offset plus the limit exceeds the total results, break the loop
        if offset+100 >= totalResults {
            break
        }

        offset += 100 // Increment offset for the next request
    }

    // Return the results in the expected structure
    return map[string]interface{}{
        "MRData": map[string]interface{}{
            "RaceTable": map[string]interface{}{
                "Races": allResults,
            },
        },
    }, nil
}



// func GetResultsData(year string, limit string) (interface{}, error) {
//     // Default limit if not provided
//     if limit == "" {
//         limit = "440" // Adjust this default as needed
//     }
//
//     fmt.Println("Using Limit: ", limit)
//
//     // Construct the API URL with the year and limit as part of the query
//     apiURL := fmt.Sprintf("%s/%s/results?limit=%s", config.APIBaseURL, year, limit)
//
//     // Log the final URL for debugging
//     fmt.Println("Final URL:", apiURL)
//
//     // Create the request
//     req, err := http.NewRequest("GET", apiURL, nil)
//     if err != nil {
//         return nil, fmt.Errorf("failed to create request: %w", err)
//     }
//
//     client := &http.Client{Timeout: 10 * time.Second}
//     resp, err := client.Do(req)
//     if err != nil {
//         return nil, fmt.Errorf("failed to make request: %w", err)
//     }
//     defer resp.Body.Close()
//
//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return nil, fmt.Errorf("failed to read response: %w", err)
//     }
//
//     var data interface{}
//     if err := json.Unmarshal(body, &data); err != nil {
//         return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
//     }
//
//     return data, nil
// }

