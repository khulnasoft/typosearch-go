package typosearch

import (
	"context"
	"net/http"
	"testing"

	"github.com/khulnasoft/typosearch-go/v2/typosearch/api"
	"github.com/khulnasoft/typosearch-go/v2/typosearch/api/pointer"
	"github.com/stretchr/testify/assert"
)

func TestStopwordRetrieve(t *testing.T) {
	expectedData := &api.StopwordsSetSchema{
		Id:        "stopwords_set1",
		Locale:    pointer.String("en"),
		Stopwords: []string{"Germany", "France", "Italy", "United States"},
	}

	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/stopwords/stopwords_set1", http.MethodGet)
		data := jsonEncode(t, map[string]any{
			"stopwords": expectedData,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	defer server.Close()

	res, err := client.Stopword(expectedData.Id).Retrieve(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedData, res)
}

func TestStopwordRetrieveOnHttpStatusErrorCodeReturnsError(t *testing.T) {
	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/stopwords/stopwords_set1", http.MethodGet)
		w.WriteHeader(http.StatusConflict)
	})
	defer server.Close()

	_, err := client.Stopword("stopwords_set1").Retrieve(context.Background())
	assert.ErrorContains(t, err, "status: 409")
}

func TestStopwordDelete(t *testing.T) {
	expectedData := &struct {
		Id string "json:\"id\""
	}{
		Id: "stopwords_set1",
	}

	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/stopwords/stopwords_set1", http.MethodDelete)
		data := jsonEncode(t, expectedData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	defer server.Close()

	res, err := client.Stopword(expectedData.Id).Delete(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedData, res)
}

func TestStopwordDeleteOnHttpStatusErrorCodeReturnsError(t *testing.T) {
	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/stopwords/stopwords_set1", http.MethodDelete)
		w.WriteHeader(http.StatusConflict)
	})
	defer server.Close()

	_, err := client.Stopword("stopwords_set1").Delete(context.Background())
	assert.ErrorContains(t, err, "status: 409")
}
