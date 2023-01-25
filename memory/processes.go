package mem

import (
	"context"
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"tuwsp/calls"
	"tuwsp/models"
)

// endId recept to users in uri
func endId(ctx context.Context, JSON string) (bytes []byte, err error) {
	var (
		// instance of params
		values = &Parameters{
			// initialization of map
			Info: make(map[string]interface{}),
		}
		// slice for urlInfo
		urlInfo []interface{}
		// required instances of models
		infoUser *models.InfoUser
		user     *models.User
	)
	// unmarshal parameter JSON yo instance values
	err = json.Unmarshal([]byte(JSON), values)
	if err != nil {
		return
	}
	// getting protocol from db
	protocol, err := calls.GetProtocol(values.ProtocolId)
	if err != nil {
		return
	} else if protocol == nil {
		return
	}
	// try convert values.user to int
	phone, err := strconv.Atoi(values.User)
	if err != nil {
		// getting user by nick name from db
		user, err = calls.GetUser(values.User)
		if err != nil {
			return
		} else if user == nil {
			return
		}
		// getting info user by userid from db
		infoUser, err = calls.GetInfoUser(user.Id)
		if err != nil {
			return
		}
	} else {
		// value.user phone
		// getting info user by phone
		infoUser, err = calls.GetInfoUserPhone(phone)
		if err != nil {
			return
		}
		// getting user by infouser.userid
		user, err = calls.GetUserId(infoUser.UserId)
		if err != nil {
			return
		} else if user == nil {
			return
		}
	}
	// getting url from db
	url, err := calls.GetUrl(user.UrlId)
	if err != nil {
		return
	} else if url == nil {
		return
	}
	// preparing domain where same user.urLId
	if url.Id == user.UrlId {
		urlInfo = append(urlInfo, url.Domain)
		urlInfo = append(urlInfo, url.Id)
	}
	// add protocol to urlInfo
	urlInfo = append(urlInfo, protocol.Protocol)
	// filling map info with prepares
	values.Info = map[string]interface{}{
		"url_info":  urlInfo,
		"user_info": user.NickName,
	}
	// translate info to json bytes
	bytes, err = json.Marshal(values.Info)
	if err != nil {
		return
	}
	return
}

// url_w make the url to redirect
func url_w(ctx context.Context, JSON string) (bytes []byte, err error) {
	var (
		// builder for concatenation
		sb = &strings.Builder{}
		// instance of parameters
		values = &Parameters{
			// initialization of map
			Info: make(map[string]interface{}),
		}
		// package url to make the url
		u = &url.URL{}
		// queries wraper
		query = u.Query()
		// slice of values of params
		val []string
	)
	// getting and wrap them values translating from json
	err = json.Unmarshal([]byte(JSON), &values)
	if err != nil {
		return
	}
	// get endpoint(s)
	endpoints, err := calls.GetEndpoint(values.UrlId)
	if err != nil {
		return
	}
	// order of endpoints by id
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].GetId() < endpoints[j].GetId()
	})
	// adding protocol
	u.Scheme = values.Protocol
	// adding domain
	u.Host = values.Domain
	// range endpoints
	for _, v := range endpoints {
		// if empty getout
		if v.Endpoint == "" {
			break
		}
		// separator of endpoint
		sb.WriteString("/")
		// add sb endpoint
		sb.WriteString(v.Endpoint)
		// concatenation of endpoint in url.Path
		u.Path = sb.String()
	}
	// getting queryvalues
	queryValues, err := calls.GetQueryValue(values.UserId)
	if err != nil {
		return
	}
	// order of queryvalues by id
	sort.Slice(queryValues, func(i, j int) bool {
		return queryValues[i].GetId() < queryValues[j].GetId()
	})
	// rage queryvalues
	for _, v := range queryValues {
		// if empty getout
		if v.ValueParam == "" {
			break
		}
		// adding to a slice value params
		val = append(val, v.ValueParam)
	}
	// getting querykey(s)
	queryKeys, err := calls.GetQueryKey(values.UrlId)
	if err != nil {
		return
	}
	// order of querykeys by id
	sort.Slice(queryKeys, func(i, j int) bool {
		return queryKeys[i].GetId() < queryKeys[j].GetId()
	})
	// range querykeys
	for i, v := range queryKeys {
		// if empty getout
		if v.KeyParam == "" {
			break
		}
		// assing slice de strings value params foreach to its key in query
		query[v.KeyParam] = []string{val[i]}
	}
	// endcode query into RawQuery
	u.RawQuery = query.Encode()
	// adding url complete to map
	values.Info["url_path"] = u.String()
	// translate info to json byte
	bytes, err = json.Marshal(values.Info)
	if err != nil {
		return
	}
	return
}
