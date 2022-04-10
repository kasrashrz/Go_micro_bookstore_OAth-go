package oath

import "net/http"

const (
	headerXPublic = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallertId = "X-User-Id"
)

type OathClient struct {

}

type OathInterface interface {

}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request){
	if request == nil {
		return
	}
}
