package fmc

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	insecureSkipVerify bool
	user               string
	password           string
	host               string
	domainBaseURL      string
	accessToken        string
	DomainUUID         string
}

type ErrorResponse struct {
	Error struct {
		Category string `json:"category"`
		Messages []struct {
			Description string `json:"description"`
		} `json:"messages"`
		Severity string `json:"severity"`
	} `json:"error"`
}

func NewClient(user, password, host string, insecureSkipVerify bool) *Client {
	return &Client{
		insecureSkipVerify: insecureSkipVerify,
		user:               user,
		password:           password,
		host:               host,
	}
}

func (v *Client) Login() error {

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/fmc_platform/v1/auth/generatetoken", v.host), nil)
	if err != nil {
		return (err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(v.user, v.password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: v.insecureSkipVerify},
	}
	c := &http.Client{Transport: tr}

	res, err := c.Do(req)
	if err != nil {
		return (err)
	}
	defer res.Body.Close()
	if res.StatusCode == 401 {
		return fmt.Errorf("wrong username or password %d %v", res.StatusCode, req.URL)
	} else if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("cannot login unknown error, status code: %d %v", res.StatusCode, req.URL)
	}

	v.accessToken = res.Header.Get("X-Auth-Access-Token")
	v.DomainUUID = res.Header.Get("DOMAIN_UUID")
	v.domainBaseURL = fmt.Sprintf("https://%s/api/fmc_config/v1/domain/%s", v.host, v.DomainUUID)
	return nil
}

func (v *Client) DoRequest(req *http.Request, item interface{}, status int) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: v.insecureSkipVerify},
	}
	c := &http.Client{Transport: tr}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)

	if err != nil {
		return err
	}
	if status == 0 {
		status = http.StatusOK
	}
	if r.StatusCode != status {
		defer r.Body.Close()

		errorRes := ErrorResponse{}
		err = json.NewDecoder(r.Body).Decode(&errorRes)
		if err != nil {
			if body, err := ioutil.ReadAll(r.Body); err != nil {
				return fmt.Errorf("wrong status code: %d, could not read error body as error json and string", r.StatusCode)
			} else {
				return fmt.Errorf("wrong status code: %d, could not read error body as error json, error body: %s", r.StatusCode, body)
			}
		}
		return fmt.Errorf("wrong status code: %d, error category: %s, error severity: %s, error messages: %v", r.StatusCode, errorRes.Error.Category, errorRes.Error.Severity, errorRes.Error.Messages)
	}
	//TODO: Handle 429 if any
	log.Printf("Status code: %d", r.StatusCode)
	if item != nil {
		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(item)
		if err != nil {
			return err
		}
	}
	return nil
}
