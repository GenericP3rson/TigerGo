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
[ ] Parse the results (possibly via a parameter)
[ ] Add more functions
[ ] Add more error-catching functions based on JSON output
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

func (conn TigerGraphConnection) GetToken() (string, error) {
	data := strings.NewReader(`{"graph": ` + conn.GraphName + `}`)            // Data is Graph
	req, err := http.NewRequest("POST", conn.Host+":9000/requesttoken", data) // Set up POST reqest
	if err != nil {                                                           // Check for errors
		return "", err
	}

	req.SetBasicAuth(conn.Username, conn.Password)                      // Create authentication
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // Add headers

	response, err := http.DefaultClient.Do(req) // Create request
	if err != nil {                             // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	type Res struct { // Result structure
		Code string 
		Expiration int
		Error bool
		Message string
		Results map[string]string
	}

	var resp Res	
	json.Unmarshal([]byte(sb), &resp) // Maps response to structure

	if resp.Error {
		return "", fmt.Errorf(resp.Message) // If there is an error, return the message
	} else {
		return resp.Results["token"], nil // Otherwise, return the token
	}
}

func (conn TigerGraphConnection) GetEndpoints(builtin bool, dynamic bool, static bool) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/endpoints?builtin=%t&dynamic=%t&static=%t", conn.Host, builtin, dynamic, static), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil // Return the endpoints

}

func (conn TigerGraphConnection) GetStatistics(seconds int) (string, error) {

	if seconds < 0 || seconds > 60 {
		return "", fmt.Errorf("Seconds value invalid, must be 0-60 inclusive")
	}

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/statistics?seconds=%d", conn.Host, seconds), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil // Return the endpoints

}

func (conn TigerGraphConnection) Echo() (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", conn.Host+":9000/echo", nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	var jsonMap map[string]interface{} // Create map
	json.Unmarshal([]byte(sb), &jsonMap)

	mess := jsonMap["message"] // Grab the value of "message"

	return fmt.Sprintf("%v", mess), nil // Return message contents

}

func (conn TigerGraphConnection) GetVersion() (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", conn.Host+":9000/version", nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

/*
VERTEX FUNCTIONS:
[√] GetVertices
[√] GetVerticesById
[√] GetVertexCount
[√] DelVertices
[√] DelVerticesById
[√] UpsertVertex
*/

func (conn TigerGraphConnection) UpsertVertex(vertexType string, vertexId string, attributes map[string]interface{}) (string, error) {

	modified_attributes := map[string]map[string]interface{}{}
	for key, value := range attributes { // Convert attributes to proper format
		modified_attributes[key] = map[string]interface{}{"value": value}
	}

	json_map := map[string]map[string]map[string]map[string]map[string]interface{}{"vertices": map[string]map[string]map[string]map[string]interface{}{vertexType: map[string]map[string]map[string]interface{}{vertexId: modified_attributes}}}

	params_json, err := json.Marshal(json_map) // Convert map to JSON
	if err != nil { // Check for errors
		return "", err
	}

	data := strings.NewReader(string(params_json)) // Convert JSON to useable data
	req, err := http.NewRequest("POST", fmt.Sprintf("%s:9000/graph/%s", conn.Host, conn.GraphName), data) // Makes POST request

	if err != nil {
		return "", err // Check for errors
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token)

	response, err := http.DefaultClient.Do(req) // Executes POST request

	if err != nil {
		return "", err // Check for error
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

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

func (conn TigerGraphConnection) DelVertices(vertexType string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s:9000/graph/%s/vertices/%s", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

func (conn TigerGraphConnection) GetVertices(vertexType string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

func (conn TigerGraphConnection) GetVerticesById(vertexType string, vertexId string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s/%s", conn.Host, conn.GraphName, vertexType, vertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

func (conn TigerGraphConnection) GetVertexCount(vertexType string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/vertices/%s?count_only=true", conn.Host, conn.GraphName, vertexType), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

/*
EDGE FUNCTIONS:
[√] GetEdges
[√] DelEdges
[√] UpsertEdge
*/

func (conn TigerGraphConnection) UpsertEdge(sourceVertexType string, sourceVertexId string, edgeType string, targetVertexType string, targetVertexId string, attributes map[string]interface{}) (string, error) {

	modified_attributes := map[string]map[string]interface{}{}

	for key, value := range attributes { // Change attributes to proper format
		modified_attributes[key] = map[string]interface{}{"value": value}
	}

	json_map := map[string]map[string]map[string]map[string]map[string]map[string]map[string]map[string]interface{}{"edges": map[string]map[string]map[string]map[string]map[string]map[string]map[string]interface{}{sourceVertexType: map[string]map[string]map[string]map[string]map[string]map[string]interface{}{sourceVertexId: map[string]map[string]map[string]map[string]map[string]interface{}{edgeType: map[string]map[string]map[string]map[string]interface{}{targetVertexType: map[string]map[string]map[string]interface{}{targetVertexId: modified_attributes}}}}}} // Create map

	params_json, err := json.Marshal(json_map) // Convert map to JSON
	if err != nil { // Check for errors
		return "", err
	}

	data := strings.NewReader(string(params_json)) // Convert JSON to readable data
	req, err := http.NewRequest("POST", fmt.Sprintf("%s:9000/graph/%s", conn.Host, conn.GraphName), data) // Makes POST request

	if err != nil {
		return "", err // Check for errors
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token)

	response, err := http.DefaultClient.Do(req) // Executes POST request

	if err != nil {
		return "", err // Check for error
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

func (conn TigerGraphConnection) DelEdges(sourceVertexType string, sourceVertexId string, edgeType string, targetVertexType string, targetVertexId string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s:9000/graph/%s/edges/%s/%s/%s/%s/%s", conn.Host, conn.GraphName, sourceVertexType, sourceVertexId, edgeType, targetVertexType, targetVertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

func (conn TigerGraphConnection) GetEdges(sourceVertexType string, sourceVertexId string) (string, error) {

	client := &http.Client{ // Creates client
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:9000/graph/%s/edges/%s/%s/_", conn.Host, conn.GraphName, sourceVertexType, sourceVertexId), nil) // Makes GET Request

	if err != nil { // Checks for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := client.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}

/*
QUERY FUNCTIONS:
[√] Run Query
*/

func (conn TigerGraphConnection) RunInstalledQuery(queryName string, params map[string]interface{}) (string, error) {

	params_json, err := json.Marshal(params) // Convert map to JSON
	if err != nil { // Check for errors
		return "", err 
	}

	data := strings.NewReader(fmt.Sprintf("%s", string(params_json))) // Convert JSON to readable data

	req, err := http.NewRequest("POST", fmt.Sprintf("%s:9000/query/%s/%s", conn.Host, conn.GraphName, queryName), data) // Make request
	if err != nil { // Check for errors
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+conn.Token) // Add authorisation header
	response, err := http.DefaultClient.Do(req)                       // Make request
	if err != nil {                                       // Check for errors
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body) // Read the response body
	if err != nil {                            // Check for errors
		return "", err
	}

	sb := string(body) // Save response as a string

	defer response.Body.Close() // Close request

	return sb, nil
}
