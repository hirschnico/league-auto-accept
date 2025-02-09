package riotapi

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	BaseURL  string
	Password string
	HTTP     *http.Client
}

type ReadyCheck struct {
	State string `json:"state"`
}

type GameSession struct {
	Phase string `json:"phase"`
}

func NewClient(creds *Credentials) *Client {
	return &Client{
		BaseURL:  "https://127.0.0.1:" + creds.Port,
		Password: creds.Password,
		HTTP: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (c *Client) CheckReadyCheck() (string, error) {
	url := c.BaseURL + "/lol-matchmaking/v1/ready-check"
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("riot", c.Password)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var readyCheck ReadyCheck
	if err := json.Unmarshal(body, &readyCheck); err != nil {
		return "", err
	}

	return readyCheck.State, nil
}

func (c *Client) AcceptMatch() error {
	url := c.BaseURL + "/lol-matchmaking/v1/ready-check/accept"
	req, _ := http.NewRequest("POST", url, nil)
	req.SetBasicAuth("riot", c.Password)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) IsInGame() (bool, error) {
	url := c.BaseURL + "/lol-gameflow/v1/session"
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("riot", c.Password)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var session GameSession
	if err := json.Unmarshal(body, &session); err != nil {
		return false, err
	}

	return session.Phase == "InProgress", nil
}
