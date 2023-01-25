package calls

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"tuwsp/models"
)

// GetProtocol do a request get to CRUD
func GetProtocol(q string) (*models.Protocol, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/protocols",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var protocol = models.Protocol{}
	err = json.
		Unmarshal(bytesBody, &protocol)
	if err != nil {
		return nil, err
	}
	return &protocol, nil
}

// GetUrl do a request get to CRUD
func GetUrl(q string) (*models.Url, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/urls",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var url = models.Url{}
	err = json.
		Unmarshal(bytesBody, &url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

// GetUser do a request get to CRUD
func GetUser(q string) (*models.User, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/users",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var user = models.User{}
	err = json.
		Unmarshal(bytesBody, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserId do a request get to CRUD
func GetUserId(q string) (*models.User, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/users-id",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var user = models.User{}
	err = json.
		Unmarshal(bytesBody, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetInfoUser do a request get to CRUD
func GetInfoUser(q string) (*models.InfoUser, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/info-users",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var infouser = models.InfoUser{}
	err = json.
		Unmarshal(bytesBody, &infouser)
	if err != nil {
		return nil, err
	}
	return &infouser, nil
}

// GetInfoUserPhone do a request get to CRUD
func GetInfoUserPhone(q int) (*models.InfoUser, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/info-users-phone",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", strconv.Itoa(q))
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var infouser = models.InfoUser{}
	err = json.
		Unmarshal(bytesBody, &infouser)
	if err != nil {
		return nil, err
	}
	return &infouser, nil
}

// GetEndpoint do a request get to CRUD
func GetEndpoint(q string) ([]*models.Endpoint, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/endpoints",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var endpoints []*models.Endpoint
	err = json.
		Unmarshal(bytesBody, &endpoints)
	if err != nil {
		return nil, err
	}
	return endpoints, nil
}

// GetQueryKey do a request get to CRUD
func GetQueryKey(q string) ([]*models.QueryKey, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/query-keys",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var querykeys []*models.QueryKey
	err = json.
		Unmarshal(bytesBody, &querykeys)
	if err != nil {
		return nil, err
	}
	return querykeys, nil
}

// GetQueryValue do a request get to CRUD
func GetQueryValue(q string) ([]*models.QueryValue, error) {
	// creating the client
	client := &http.Client{}
	// parse of url to call
	u, err := url.
		Parse(
			"http://localhost:3000/query-values",
		)
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := client.
		Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.
		ReadAll(resp.Body)
	// unmarshal to struct
	var queryvalues []*models.QueryValue
	err = json.
		Unmarshal(bytesBody, &queryvalues)
	if err != nil {
		return nil, err
	}
	return queryvalues, nil
}
