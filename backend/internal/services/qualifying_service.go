
package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
    "F1DataVisualizer/backend/config"
    "strconv"  
)

func GetQualifyingData(year string) (interface{}, error) {
    var allData []interface{}  // To store combined data
    offset := 0
    limit := 100               // Set the limit to the maximum allowed by the API
    totalResults := 0          // Variable to track the total number of results

    for {
        apiURL, err := url.Parse(fmt.Sprintf("%s/%s/qualifying", config.APIBaseURL, year))
        if err != nil {
            return nil, fmt.Errorf("failed to parse base URL: %w", err)
        }

        query := apiURL.Query()
        query.Set("limit", fmt.Sprintf("%d", limit))
        query.Set("offset", fmt.Sprintf("%d", offset))
        apiURL.RawQuery = query.Encode()

        fmt.Println("Final URL:", apiURL.String())

        req, err := http.NewRequest("GET", apiURL.String(), nil)
        if err != nil {
            return nil, fmt.Errorf("failed to create request: %w", err)
        }

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

        var data map[string]interface{}
        if err := json.Unmarshal(body, &data); err != nil {
            return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
        }
        

        //COME BACK AND LOOK AT THIS
        totalValue := data["MRData"].(map[string]interface{})["total"]

        // Check the type and convert appropriately
        switch v := totalValue.(type) {
        case string:
            totalResults, err = strconv.Atoi(v)
            if err != nil {
                return nil, fmt.Errorf("failed to convert total to integer: %w", err)
            }
        case float64:
            totalResults = int(v)
        default:
            return nil, fmt.Errorf("unexpected type for total: %T", v)
        }
        //STOP

        // Append data to allData (adjust based on actual JSON structure)
        races := data["MRData"].(map[string]interface{})["RaceTable"].(map[string]interface{})["Races"].([]interface{})
        allData = append(allData, races...)

        // Check if we have retrieved all data
        if offset+limit >= totalResults {
            break
        }

        // Increment offset for the next batch
        offset += limit
    }

    return allData, nil
}

