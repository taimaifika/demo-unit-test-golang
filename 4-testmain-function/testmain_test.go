package testMain

import (
	"net/http"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func TestHttpBinGet(t *testing.T) {
	t.Run("Testing a get request", func(t *testing.T) {
		t.Log("sending get request to https://httpbin.org/get")
		resp, err := client.R().Get("https://httpbin.org/get")
		if err != nil {
			t.Error(err.Error())
		}
		if resp.StatusCode() != http.StatusOK {
			t.Errorf("An unexpected status code has been returned: %d", resp.StatusCode())
		}

	})
}

func TestHttpBinPost(t *testing.T) {
	t.Run("Testing a post request", func(t *testing.T) {
		t.Log("sending post request to https://httpbin.org/post")
		resp, err := client.R().Post("https://httpbin.org/post")
		if err != nil {
			t.Error(err.Error())
		}
		if resp.StatusCode() != http.StatusOK {
			t.Errorf("An unexpected status code has been returned: %d", resp.StatusCode())
		}

	})
}

func TestMain(m *testing.M) {
	// setup()
	client = resty.New()
	exitVal := m.Run()
	if exitVal == 0 {
		// teardown()
	}
	os.Exit(exitVal)
}
