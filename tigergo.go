package tigergo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Echo(Host string, Token string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://"+Host+".i.tgcloud.io:9000/echo", nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+Token) // Add authorisation header
	response, err := client.Do(req)                  // Make request
	if err != nil {                                  // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	var jsonMap map[string]interface{} // Create map
	json.Unmarshal([]byte(sb), &jsonMap)

	mess := jsonMap["message"] // Grab the value of "message"

	return fmt.Sprintf("%v", mess) // Return message contents

}
