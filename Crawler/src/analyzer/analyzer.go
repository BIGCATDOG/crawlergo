package analyzer

import "net/http"

type Analyzer interface {
	buildRequest(response http.Response) *http.Request
}
