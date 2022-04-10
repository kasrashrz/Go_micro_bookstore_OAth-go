package oath

import (
	"encoding/json"
	"fmt"
	"github.com/Go_micro_bookstore_OAth-go/oath/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
	"time"
)

const (
	headerXPublic    = "X-Public"
	headerXClientId  = "X-Client-Id"
	headerXCallerId  = "X-User-Id"
	paramAccessToken = "access_token"
)

var (
	oathRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
}

type OathInterface interface {
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request) *errors.RestErr {
	if request == nil {
		return nil
	}
	accessToken := request.URL.Query().Get(paramAccessToken)
	if accessToken == "" {
		return nil
	}
	return nil
}

func GetAccessToken(accessTokenId string) (*accessToken, *errors.RestErr) {
	response := oathRestClient.Get(fmt.Sprintf("/oath/access_token/%s", accessTokenId))

	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &accessToken{}); err != nil {
		return nil, errors.InternalServerError("error when trying to unmarshal users response")
	}
	return &at, nil
}
