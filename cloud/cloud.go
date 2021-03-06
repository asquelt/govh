package cloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/toorop/govh"
)

// Client is a REST client for cloud API
type Client struct {
	*govh.OVHClient
}

// New return a new Cloud API Client
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// GetPassports returns cloud passports
func (c *Client) GetPassports() (passports []string, err error) {
	r, err := c.GET("cloud")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &passports)
	return
}

// GetPrices return cloud prices
func (c *Client) GetPrices() (prices GetPriceResponse, err error) {
	var r govh.APIResponse
	r, err = c.GET("cloud/price")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &prices)
	return
}

// GetProjectIDs returns clouds projects IDs
func (c *Client) GetProjectIDs() (projectIDs []string, err error) {
	r, err := c.GET("cloud/project")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &projectIDs)
	return
}

// GetProject return a project
func (c *Client) GetProject(id string) (p Project, err error) {
	r, err := c.GET("cloud/project/" + url.QueryEscape(id))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &p)
	return
}

// GetInstances return instance of project projectID
func (c *Client) GetInstances(projectID string) (instances []Instance, err error) {
	r, err := c.GET("cloud/project/" + url.QueryEscape(projectID) + "/instance")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &instances)
	return
}

// GetSnapshots return snapshots( private image) of project projectID
func (c *Client) GetSnapshots(projectID string) (images []Image, err error) {
	r, err := c.GET("cloud/project/" + url.QueryEscape(projectID) + "/snapshot")
	if err != nil {
		return
	}
	fmt.Println(string(r.Body))
	err = json.Unmarshal(r.Body, &images)
	return
}
