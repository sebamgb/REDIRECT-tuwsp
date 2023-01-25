package calls

import (
	"bytes"
	"encoding/json"
	"net/http"

	"tuwsp/models"
)

// InsertProtocol do a request post to CRUD
func InsertProtocol(protocol *models.Protocol) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(protocol)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/protocols",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertUrl do a request post to CRUD
func InsertUrl(url *models.Url) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(url)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/urls",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertUser do a request post to CRUD
func InsertUser(user *models.User) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(user)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/users",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertInfoUser do a request post to CRUD
func InsertInfoUser(infouser *models.InfoUser) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(infouser)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/info-users",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertEndpoint do a request post to CRUD
func InsertEndpoint(endpoint *models.Endpoint) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(endpoint)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/endpoints",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertQueryKey do a request post to CRUD
func InsertQueryKey(querykey *models.QueryKey) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(querykey)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/query-keys",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// InsertQueryValue do a request post to CRUD
func InsertQueryValue(queryvalue *models.QueryValue) error {
	// Creating the client
	client := &http.Client{}
	// Converting struct to bytes
	requestBody, err := json.Marshal(queryvalue)
	if err != nil {
		return err
	}
	// Create the Request
	resp, err := client.Post(
		"http://localhost:3000/query-values",
		"application/json",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
