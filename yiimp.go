package yiimp

import (
	"github.com/dghubble/sling"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"log"
	"strings"
)

type YiimpClient struct {
	sling      *sling.Sling
	httpClient *yiimpHttpClient
	apikey     string
}

// server send the api response with text/html content type
// we fix this: change content type to json
type yiimpHttpClient struct {
	client    *http.Client
	debug     bool
	useragent string
}

func (d yiimpHttpClient) Do(req *http.Request) (*http.Response, error) {
	if d.debug {
		d.dumpRequest(req)
	}
	if d.useragent != "" {
		req.Header.Set("User-Agent", d.useragent)
	}
	client := func() (*http.Client) {
		if d.client != nil {
			return d.client
		} else {
			return http.DefaultClient
		}
	}()
	if client.Transport != nil {
		if transport, ok := client.Transport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	} else {
		if transport, ok := http.DefaultTransport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	}
	resp, err := client.Do(req)
	if d.debug {
		d.dumpResponse(resp)
	}
	if err == nil {
		if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
			resp.Header.Set("Content-Type", "application/json")
		}
	}
	body, err := ioutil.ReadAll(resp.Body);
	body = bytes.Replace(body, []byte(":,"), []byte(":\"\","), -1)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return resp, err
}

func (d yiimpHttpClient) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (d yiimpHttpClient) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func NewYiimpClient(client *http.Client, BaseURL string, ApiKey string, UserAgent string) *YiimpClient {
	httpClient := &yiimpHttpClient{client:client, useragent:UserAgent}
	return &YiimpClient{
		httpClient: httpClient,
		sling: sling.New().Doer(httpClient).Base(BaseURL).Path("/api/"),
		apikey: ApiKey,
	}
}

func (client YiimpClient) SetDebug(debug bool) {
	client.httpClient.debug = debug
}
