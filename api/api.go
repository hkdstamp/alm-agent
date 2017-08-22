package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/mobingi/alm-agent/config"
)

var logregion = "ap-northeast-1"

type clientInterface interface {
	buildURI(string) string
	getHTTPClient() *http.Client
	setConfig(*config.Config) error
	getConfig() *config.Config
}

var c clientInterface

type apiToken struct {
	tokenType string
	token     string
}

var apitoken apiToken

type StsToken struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

var stsToken StsToken

type client struct {
	config *config.Config
}

func (c *client) getHTTPClient() *http.Client {
	return &http.Client{}
}

func (c *client) buildURI(path string) string {
	return c.getConfig().APIHost + path
}

func SetConfig(conf *config.Config) error {
	c.setConfig(conf)
	return nil
}

func (c *client) setConfig(conf *config.Config) error {
	c.config = conf
	return nil
}

func (c *client) getConfig() *config.Config {
	return c.config
}

func init() {
	log.Debug("Initializing api client...")
	c = &client{}
}

var Get = func(path string, values url.Values, target interface{}) error {
	log.Debug(c.getConfig())
	req, err := http.NewRequest("GET", c.buildURI(path), nil)

	if apitoken.token != "" && apitoken.tokenType != "" {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", apitoken.tokenType, apitoken.token))
	}

	req.URL.RawQuery = values.Encode()
	httpClient := c.getHTTPClient()
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = json.Unmarshal(res, &target)
	if err != nil {
		return err
	}

	return nil
}

var Post = func(path string, values url.Values, target interface{}) error {
	log.Debug(c.getConfig())
	req, err := http.NewRequest("POST", c.buildURI(path), strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if apitoken.token != "" && apitoken.tokenType != "" {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", apitoken.tokenType, apitoken.token))
	}

	httpClient := c.getHTTPClient()
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = json.Unmarshal(res, &target)
	if err != nil {
		return err
	}

	return nil
}
