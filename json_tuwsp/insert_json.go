package jsontuwsp

import (
	"encoding/json"
	"io/ioutil"

	"tuwsp/calls"
	"tuwsp/models"
)

// InsertProtocolJson insert json/insert_protocol.json to db
func InsertProtocolJson() (err error) {
	var protocol []*models.Protocol
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_protocol.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &protocol)
	if err != nil {
		return
	}
	// range it for insert
	for i := range protocol {
		calls.InsertProtocol(protocol[i])
	}
	return
}

// InsertUrlJson insert json/insert_url.json to db
func InsertUrlJson() (err error) {
	var url []*models.Url
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_url.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &url)
	if err != nil {
		return
	}
	// range it for insert
	for i := range url {
		calls.InsertUrl(url[i])
	}
	return
}

// InsertUserJson insert json/insert_user.json to db
func InsertUserJson() (err error) {
	var user []*models.User
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_user.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &user)
	if err != nil {
		return
	}
	// range it for insert
	for i := range user {
		calls.InsertUser(user[i])
	}
	return
}

// InsertInfoUserJson insert json/insert_info_user.json to db
func InsertInfoUserJson() (err error) {
	var infoUser []*models.InfoUser
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_info_user.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &infoUser)
	if err != nil {
		return
	}
	// range it for insert
	for i := range infoUser {
		calls.InsertInfoUser(infoUser[i])
	}
	return
}

// InsertQueryKeyJson insert json/insert_query_key.json to db
func InsertQueryKeyJson() (err error) {
	var querykey []*models.QueryKey
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_query_key.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &querykey)
	if err != nil {
		return
	}
	// range it for insert
	for i := range querykey {
		// increment foreach Number
		querykey[i].Number = i + 1
		calls.InsertQueryKey(querykey[i])
	}
	return
}

// InsertQueryValueJson insert json/insert_query_value.json to db
func InsertQueryValueJson() (err error) {
	var queryvalue []*models.QueryValue
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_query_value.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &queryvalue)
	if err != nil {
		return
	}
	// range it for insert
	for i := range queryvalue {
		// increment foreach number
		queryvalue[i].Number = i + 1
		calls.InsertQueryValue(queryvalue[i])
	}
	return
}

// InsertEndpointJson insert json/insert_endpoint.json to db
func InsertEndpointJson() (err error) {
	var endpoint []*models.Endpoint
	// Reading files json
	bytes_d, err := ioutil.ReadFile("json_tuwsp/json/insert_endpoint.json")
	if err != nil {
		return
	}
	// Bytes json to struct
	err = json.Unmarshal(bytes_d, &endpoint)
	if err != nil {
		return
	}
	// range it for insert
	for i := range endpoint {
		// increment foreach Number
		endpoint[i].Number = i + 1
		calls.InsertEndpoint(endpoint[i])
	}
	return
}
