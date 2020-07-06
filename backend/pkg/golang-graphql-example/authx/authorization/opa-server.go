package authorization

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

type inputOPA struct {
	Input *inputDataOPA `json:"input"`
}

type inputDataOPA struct {
	User    *models.OIDCUser  `json:"user"`
	Request *requestDataOPA   `json:"request"`
	Tags    map[string]string `json:"tags"`
}

type requestDataOPA struct {
	Method     string            `json:"method"`
	Protocol   string            `json:"protocol"`
	Headers    map[string]string `json:"headers"`
	RemoteAddr string            `json:"remoteAddr"`
	Scheme     string            `json:"scheme"`
	Host       string            `json:"host"`
	ParsedPath []string          `json:"parsed_path"`
	Path       string            `json:"path"`
}

type opaAnswer struct {
	Result bool `json:"result"`
}

func isOPAServerAuthorized(req *http.Request, oidcUser *models.OIDCUser, opaServerCfg *config.OPAServerAuthorization) (bool, error) {
	// Get trace from request
	trace := tracing.GetTraceFromContext(req.Context())
	// Generate child trace
	childTrace := trace.GetChildTrace("opa-server.request")
	defer childTrace.Finish()
	// Add data
	childTrace.SetTag("opa.uri", opaServerCfg.URL)

	// Transform headers into map
	headers := make(map[string]string)
	for k, v := range req.Header {
		headers[strings.ToLower(k)] = v[0]
	}
	// Parse path
	parsedPath := deleteEmpty(strings.Split(req.RequestURI, "/"))
	// Calculate scheme
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	// Generate OPA Server input data
	input := &inputOPA{
		Input: &inputDataOPA{
			User: oidcUser,
			Tags: opaServerCfg.Tags,
			Request: &requestDataOPA{
				Method:     req.Method,
				Protocol:   req.Proto,
				Headers:    headers,
				RemoteAddr: req.RemoteAddr,
				Scheme:     scheme,
				Host:       req.Host,
				ParsedPath: parsedPath,
				Path:       req.RequestURI,
			},
		},
	}
	// Json encode body
	bb, err := json.Marshal(input)
	if err != nil {
		return false, err
	}

	// Making request to OPA server
	resp, err := http.Post(opaServerCfg.URL, "application/json", bytes.NewBuffer(bb))
	if err != nil {
		return false, err
	}
	// Defer closing body
	defer resp.Body.Close()

	// Prepare answer
	var answer opaAnswer
	// Decode answer
	err = json.NewDecoder(resp.Body).Decode(&answer)
	if err != nil {
		return false, err
	}

	return answer.Result, nil
}

func deleteEmpty(s []string) []string {
	var r []string

	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}
