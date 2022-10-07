package client

import (
	"reflect"
	
	"github.com/pkg/errors"
	"github.com/google/jsonapi"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

type Cause struct {
	ID string `jsonapi:"primary,causes"`
	Name string `jsonapi:"attr,name,omitempty"`
  Slug string `jsonapi:"attr,slug,omitempty"`
  Description string `jsonapi:"attr,description,omitempty"`
}

func (c *Client) ListCauses(params *rootlygo.ListCausesParams) ([]interface{}, error) {
	req, err := rootlygo.NewListCausesRequest(c.Rootly.Server, params)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request: %s", err.Error())
	}

	causes, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(new(Cause)))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling: %s", err.Error())
	}

	return causes, nil
}

func (c *Client) CreateCause(d *Cause) (*Cause, error) {
	buffer, err := MarshalData(d)
	if err != nil {
		return nil, errors.Errorf("Error marshaling cause: %s", err.Error())
	}

	req, err := rootlygo.NewCreateCauseRequestWithBody(c.Rootly.Server, c.ContentType, buffer)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to perform request to create cause: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Cause))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling cause: %s", err.Error())
	}

	return data.(*Cause), nil
}

func (c *Client) GetCause(id string) (*Cause, error) {
	req, err := rootlygo.NewGetCauseRequest(c.Rootly.Server, id)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request to get cause: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Cause))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling cause: %s", err.Error())
	}

	return data.(*Cause), nil
}

func (c *Client) UpdateCause(id string, cause *Cause) (*Cause, error) {
	buffer, err := MarshalData(cause)
	if err != nil {
		return nil, errors.Errorf("Error marshaling cause: %s", err.Error())
	}

	req, err := rootlygo.NewUpdateCauseRequestWithBody(c.Rootly.Server, id, c.ContentType, buffer)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request to update cause: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Cause))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling cause: %s", err.Error())
	}

	return data.(*Cause), nil
}

func (c *Client) DeleteCause(id string) error {
	req, err := rootlygo.NewDeleteCauseRequest(c.Rootly.Server, id)
	if err != nil {
		return errors.Errorf("Error building request: %s", err.Error())
	}

	_, err = c.Do(req)
	if err != nil {
		return errors.Errorf("Failed to make request to delete cause: %s", err.Error())
	}

	return nil
}
