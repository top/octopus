package facebook

import (
	"net/http"
)

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PUT    Method = "PUT"
)

type (
	DebugMode string // https://developers.facebook.com/docs/graph-api/using-graph-api/debugging
	Method    string
)

var (
	// https://developers.facebook.com/docs/apps/versions
	Version           string    // v11.0
	Debug             DebugMode // all, info, warning
	RFC3339Timestamps bool

	defaultSession = &Session{}
)

func Api(path string, method Method, params Params) (Result, error) {
	return defaultSession.Api(path, method, params)
}

func Get(path string, params Params) (Result, error) {
	return Api(path, GET, params)
}

func Post(path string, params Params) (Result, error) {
	return Api(path, POST, params)
}

func Delete(path string, params Params) (Result, error) {
	return Api(path, DELETE, params)
}

func Put(path string, params Params) (Result, error) {
	return Api(path, PUT, params)
}

func BatchApi(accessToken string, params ...Params) ([]Result, error) {
	return Batch(Params{"access_token": accessToken}, params...)
}

func Batch(batchParams Params, params ...Params) ([]Result, error) {
	return defaultSession.Batch(batchParams, params...)
}

func Request(request *http.Request) (Result, error) {
	return defaultSession.Request(request)
}

func DefaultHttpClient() HttpClient {
	return defaultSession.HttpClient
}

func SetHttpClient(client HttpClient) {
	defaultSession.HttpClient = client
}
