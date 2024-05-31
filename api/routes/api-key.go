package routes

import (
	"net/http"
	"strconv"

	"github.com/shellhub-io/shellhub/api/pkg/gateway"
	"github.com/shellhub-io/shellhub/api/pkg/guard"
	"github.com/shellhub-io/shellhub/pkg/api/requests"
	"github.com/shellhub-io/shellhub/pkg/api/responses"
)

const (
	CreateAPIKeyURL = "/namespaces/api-key"
	ListAPIKeysURL  = "/namespaces/api-key"
	UpdateAPIKeyURL = "/namespaces/api-key/:name"
	DeleteAPIKeyURL = "/namespaces/api-key/:name"
)

func (h *Handler) CreateAPIKey(c gateway.Context) error {
	req := new(requests.CreateAPIKey)

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	res := new(responses.CreateAPIKey)
	if err := guard.EvaluatePermission(c.Role(), guard.Actions.APIKey.Create, func() error {
		var err error
		res, err = h.service.CreateAPIKey(c.Ctx(), req)

		return err
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) ListAPIKeys(c gateway.Context) error {
	req := new(requests.ListAPIKey)

	if err := c.Bind(req); err != nil {
		return err
	}

	req.Paginator.Normalize()

	if req.Sorter.By == "" {
		req.Sorter.By = "expires_in"
	}

	if req.Sorter.Order == "" {
		req.Sorter.Order = "desc"
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	res, count, err := h.service.ListAPIKeys(c.Ctx(), req)
	if err != nil {
		return err
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(count))

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateAPIKey(c gateway.Context) error {
	req := new(requests.UpdateAPIKey)

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := guard.EvaluatePermission(c.Role(), guard.Actions.APIKey.Edit, func() error {
		return h.service.UpdateAPIKey(c.Ctx(), req) // TODO: name
	}); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteAPIKey(c gateway.Context) error {
	req := new(requests.DeleteAPIKey)

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := guard.EvaluatePermission(c.Role(), guard.Actions.APIKey.Delete, func() error {
		return h.service.DeleteAPIKey(c.Ctx(), req)
	}); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
