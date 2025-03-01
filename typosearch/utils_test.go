package typosearch

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestServerAndClient(handler func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	return server, NewClient(WithServer(server.URL))
}

func validateRequestMetadata(t *testing.T, r *http.Request, expectedEndpoint string, expectedMethod string) {
	if r.RequestURI != expectedEndpoint {
		t.Fatal("Invalid request endpoint!")
	}
	if r.Method != expectedMethod {
		t.Fatal("Invalid HTTP method!")
	}
}

func jsonEncode(t *testing.T, v any) []byte {
	t.Helper()
	data, err := json.Marshal(v)
	assert.NoError(t, err)
	return data
}
