package requests

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"

	"github.com/deemount/discogs/api/utils"
)

// HTTPClient is a struct
type HTTPClient struct {
	client http.Client
}

// Send is a method
func (c *HTTPClient) Send(url string, params url.Values, headers map[string]string) ([]byte, error) {

	var err error

	log.Print("send request uri: %v", url+"?"+params.Encode())

	// initialize request
	req, err := http.NewRequest("GET", url+"?"+params.Encode(), nil)
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
	}

	// add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte("0"), err
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return []byte("0"), utils.ErrUnauthorized
		case http.StatusNotFound:
			return []byte("0"), utils.ErrNotFound
		default:
			return []byte("0"), utils.ErrDefaultResponse // resp.Status
		}
	}

	// check response mime type
	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal("Error on parsing mime type. ", err.Error())
	}
	if mimeType != "application/json" {
		log.Fatalf("Error on mime type. Response Content-Type is '%s', but should be 'application/json'.", mimeType)
	}

	// return response body
	return ioutil.ReadAll(resp.Body)

}
