package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	resty "github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/shellhub-io/shellhub/pkg/revdial"
	"github.com/shellhub-io/shellhub/pkg/wsconnadapter"
	log "github.com/sirupsen/logrus"
)

const (
	apiHost   = "ssh.shellhub.io"
	apiPort   = 80
	apiScheme = "https"
)

type Client interface {
	commonAPI
	publicAPI
}

type publicAPI interface {
	GetInfo(agentVersion string) (*models.Info, error)
	Endpoints() (*models.Endpoints, error)
	AuthDevice(req *models.DeviceAuthRequest) (*models.DeviceAuthResponse, error)
	NewReverseListener(ctx context.Context, token string) (*revdial.Listener, error)
	AuthPublicKey(req *models.PublicKeyAuthRequest, token string) (*models.PublicKeyAuthResponse, error)
}

func (c *client) GetInfo(agentVersion string) (*models.Info, error) {
	var info *models.Info

	response, err := c.http.R().
		SetResult(&info).
		Get("/info?agent_version=" + agentVersion)
	if err != nil {
		return nil, err
	}

	if err := ErrorFromResponse(response); err != nil {
		return nil, err
	}

	return info, nil
}

func (c *client) AuthDevice(req *models.DeviceAuthRequest) (*models.DeviceAuthResponse, error) {
	var res *models.DeviceAuthResponse

	response, err := c.http.R().
		AddRetryCondition(func(r *resty.Response, err error) bool {
			identity := func(mac, hostname string) string {
				if mac != "" {
					return mac
				}

				return hostname
			}

			log.WithFields(log.Fields{
				"tenant_id":   req.TenantID,
				"identity":    identity(req.Identity.MAC, req.Hostname),
				"status_code": r.StatusCode(),
			}).Debug("failed to authenticate device")

			return r.IsError()
		}).
		SetBody(req).
		SetResult(&res).
		Post("/api/devices/auth")
	if err != nil {
		return nil, err
	}

	if err := ErrorFromResponse(response); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Endpoints() (*models.Endpoints, error) {
	var endpoints *models.Endpoints

	response, err := c.http.R().
		SetResult(&endpoints).
		Get("/endpoints")
	if err != nil {
		return nil, err
	}

	if err := ErrorFromResponse(response); err != nil {
		return nil, err
	}

	return endpoints, nil
}

func (c *client) NewReverseListener(ctx context.Context, token string) (*revdial.Listener, error) {
	var err error

	req := c.http.R()
	req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	req.URL, err = url.JoinPath(c.http.BaseURL, "/ssh/connection")
	if err != nil {
		return nil, err
	}

	conn, _, err := DialContext(ctx, req.URL, req.Header)
	if err != nil {
		return nil, err
	}

	return revdial.NewListener(wsconnadapter.New(conn),
		func(ctx context.Context, path string) (*websocket.Conn, *http.Response, error) {
			req.URL, err = url.JoinPath(c.http.BaseURL, path)
			if err != nil {
				return nil, nil, err
			}

			return DialContext(ctx, req.URL, nil)
		},
	), nil
}

func (c *client) AuthPublicKey(req *models.PublicKeyAuthRequest, token string) (*models.PublicKeyAuthResponse, error) {
	var res *models.PublicKeyAuthResponse

	response, err := c.http.R().
		SetBody(req).
		SetResult(&res).
		SetAuthToken(token).
		Post("/api/auth/ssh")
	if err != nil {
		return nil, err
	}

	if err := ErrorFromResponse(response); err != nil {
		return nil, err
	}

	return res, nil
}

func tunnelDial(ctx context.Context, protocol, address string, path string) (*websocket.Conn, *http.Response, error) {
	getPortFromProtocol := func(protocol string) int {
		if protocol == "wss" {
			return 443
		}

		return 80
	}

	return websocket.DefaultDialer.DialContext(ctx, strings.Join([]string{fmt.Sprintf("%s://%s:%d", protocol, address, getPortFromProtocol(protocol)), path}, ""), nil)
}
