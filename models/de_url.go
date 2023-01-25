package models

type Protocol struct {
	Id       string `json:"id"`
	Protocol string `json:"protocol"`
}

type Url struct {
	Id         string `json:"id"`
	Domain     string `json:"domain"`
	ProtocolId string `json:"protocol_id"`
}

type QueryKey struct {
	Number   int    `json:"number"`
	Id       int    `json:"id"`
	KeyParam string `json:"key_param"`
	UrlId    string `json:"url_id"`
}

// GetId gets an queryKeyId
func (qk *QueryKey) GetId() int {
	return qk.Id
}

type QueryValue struct {
	Number     int    `json:"number"`
	Id         int    `json:"id"`
	ValueParam string `json:"value_param"`
	UserId     string `json:"user_id"`
}

// GetId gets an queryValueId
func (qv *QueryValue) GetId() int {
	return qv.Id
}

type Endpoint struct {
	Number   int    `json:"number"`
	Id       int    `json:"id"`
	Endpoint string `json:"endpoint"`
	UrlId    string `json:"url_id"`
}

// GetId gets an endpointId
func (e *Endpoint) GetId() int {
	return e.Id
}
