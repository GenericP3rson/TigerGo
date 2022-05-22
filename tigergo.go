package TigerGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
OVERALL TODO
[ ] Parse the results
[ ] Add more functions
*/

type TigerGraphConnection struct {
	Token     string
	Host      string
	GraphName string
	Username  string
	Password  string
}

/*
GENERAL FUNCTIONS
*/

func (conn TigerGraphConnection) GetToken() string {
	data := strings.NewReader(`{"graph": ` + conn.GraphName + `}`)            // Data is Graph
	req, err := http.NewRequest("POST", conn.Host+":9000/requesttoken", data) // Set up POST reqest
	if err != nil {                                                           // Check for errors
		return err.Error()
	}

	req.SetBasicAuth(conn.Username, conn.Password)                      // Create authentication
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // Add headers

	response, err := http.DefaultClient.Do(req) // Create request
	if err != nil {                             // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb // Return the endpoints

}

func (conn TigerGraphConnection) GetEndpoints(builtin bool, dynamic bool, static bool) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/endpoints?builtin=%t&dynamic=%t&static=%t", conn.Host, builtin, dynamic, static), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb // Return the endpoints

}

func (conn TigerGraphConnection) GetStatistics(seconds int) string {

	if seconds < 0 || seconds > 60 {
		return "Seconds value invalid, must be 0-60 inclusive"
	}

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/statistics?seconds=%d", conn.Host, seconds), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb // Return the endpoints

}

func (conn TigerGraphConnection) Echo() string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", conn.Host+":9000/echo", nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
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

func (conn TigerGraphConnection) GetVersion() string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", conn.Host+":9000/version", nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

/*
VERTEX FUNCTIONS:
[√] GetVertices
[√] GetVerticesById
[√] GetVertexCount
[√] DelVertices
[√] DelVerticesById
[ ] UpsertVertex
*/

func (conn TigerGraphConnection) DelVerticesById(vertexType string, vertexId string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s:9000/graph/%s/vertices/%s/%s", conn.Host, conn.GraphName, vertexType, vertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

func (conn TigerGraphConnection) DelVertices(vertexType string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s:9000/graph/%s/vertices/%s", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

func (conn TigerGraphConnection) GetVertices(vertexType string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

func (conn TigerGraphConnection) GetVerticesById(vertexType string, vertexId string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s/%s", conn.Host, conn.GraphName, vertexType, vertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

func (conn TigerGraphConnection) GetVertexCount(vertexType string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s?count_only=true", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

/*
EDGE FUNCTIONS:
[√] GetEdges
[√] DelEdges
[ ] UpsertEdge
*/

func (conn TigerGraphConnection) DelEdges(sourceVertexType string, sourceVertexId string, edgeType string, targetVertexType string, targetVertexId string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s:9000/graph/%s/edges/%s/%s/%s/%s/%s", conn.Host, conn.GraphName, sourceVertexType, sourceVertexId, edgeType, targetVertexType, targetVertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

func (conn TigerGraphConnection) GetEdges(sourceVertexType string, sourceVertexId string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/edges/%s/%s/_", conn.Host, conn.GraphName, sourceVertexType, sourceVertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}

/*
QUERY FUNCTIONS:
[√] Run Query -> TODO: ADD PARAMS
*/

func (conn TigerGraphConnection) RunInstalledQuery(queryName string) string {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/query/%s/%s", conn.Host, conn.GraphName, queryName), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return err.Error()
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb
}
