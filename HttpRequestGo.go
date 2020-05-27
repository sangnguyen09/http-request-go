package HttpRequestGo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

type HttpClient struct {
	Client  *resty.Client
	//Proxy   string
	//Socks5  string
	//Cookies []*http.Cookie
	//Headers map[string]string
}

func (h *HttpClient) SetCookies(cookies []*http.Cookie)  {
	h.Client.SetCookies(cookies)
}
func (h *HttpClient) SetHeaders(headers map[string]string)  {
	h.Client.SetHeaders(headers)
}

func (h *HttpClient) SetProxy(proxyIp string, ) error {
	proxy := fmt.Sprintf("%v%v", "http://", proxyIp);
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	// Set up a custom HTTP transport for client
	customTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	h.Client.SetTransport(customTransport)
	return nil
}

const (
	SOCKS4 = iota
	SOCKS4A
	SOCKS5
)

func (h *HttpClient) SetSocks(socksType int, socksIp string) error {
	socks :="";
	switch socksType {
	case SOCKS4:
		socks = fmt.Sprintf("%v%v", "socks4://", socksIp)
	case SOCKS4A:
		socks = fmt.Sprintf("%v%v", "socks4a://", socksIp)
	case SOCKS5:
		socks = fmt.Sprintf("%v%v", "socks5://", socksIp)
	}

	proxyUrl, err := url.Parse(socks)
	if err != nil {
		return err
	}
	// Set up a custom HTTP transport for client
	customTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	h.Client.SetTransport(customTransport)
	return nil
}

func (h *HttpClient) Get(url string, ) (*resty.Response, error) {
	resp, err := h.Client.R().Get(url)
	return resp, err
}
