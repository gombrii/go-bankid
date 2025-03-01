package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	bidURL     string
}

type Config struct {
	ca         []byte
	clientCert []byte
	Timeout    time.Duration
}

func New(cfg Config) (Client, error) {
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cfg.ca); !ok {
		return Client{}, errors.New("error parsing CA cert from PEM")
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            certPool,
				InsecureSkipVerify: false,
				Certificates: []tls.Certificate{{
					Certificate: [][]byte{cfg.clientCert},
				}},
			},
		},
		Timeout: cfg.Timeout,
	}

	return Client{httpClient: &client}, nil
}

func (c Client) Auth(ctx context.Context, req authReq) {

}

func (c Client) Sign() {

}

func (c Client) Payment() {

}

func (c Client) PhoneAuth() {

}
func (c Client) PhoneSign() {

}

func (c Client) Collect() {

}
func (c Client) Cancel() {

}
