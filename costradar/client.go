package costradar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type AccessConfig struct {
	ReaderMode            string `json:"readerMode"`
	AssumeRoleArn         string `json:"assumeRoleArn"`
	AssumeRoleExternalId  string `json:"assumeRoleExternalId"`
	AssumeRoleSessionName string `json:"assumeRoleSessionName"`
}

type CostAndUsageReportSubscription struct {
	ID               string       `json:"id"`
	ReportName       string       `json:"reportName"`
	BucketName       string       `json:"bucketName"`
	BucketRegion     string       `json:"bucketRegion"`
	BucketPathPrefix string       `json:"bucketPathPrefix"`
	TimeUnit         string       `json:"timeUnit"`
	AccessConfig     AccessConfig `json:"accessConfig"`
}

type CostAndUsageReportPayload struct {
	Status  bool                           `json:"status"`
	Error   string                         `json:"error"`
	Payload CostAndUsageReportSubscription `json:"payload"`
}

type Client interface {
	GetCostAndUsageReportSubscription(id string) (*CostAndUsageReportPayload, error)
	CreateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportPayload, error)
	UpdateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportPayload, error)
	DeleteCostAndUsageReportSubscription(id string) error
}

type ClientGraphqlClient struct {
	endpoint   string
	token      string
	httpClient *http.Client
}

func NewCostRadarClient(endpoint, token string) Client {
	return &ClientGraphqlClient{
		endpoint:   endpoint,
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *ClientGraphqlClient) graphql(query string, variables map[string]interface{}, dataPath string) (data interface{}, err error) {

	payload := new(bytes.Buffer)

	json.NewEncoder(payload).Encode(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})

	req, _ := http.NewRequest("POST", c.endpoint, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.token))
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	data = gjson.GetBytes(body, dataPath).Value()
	return data, err
}

func (c *ClientGraphqlClient) GetCostAndUsageReportSubscription(id string) (*CostAndUsageReportPayload, error) {
	query := GetCostAndUsageReportSubscriptionQuery

	variables := map[string]interface{}{
		"id": id,
	}

	subscription := CostAndUsageReportSubscription{}

	data, err := c.graphql(query, variables, "data.awsCURSubscription")
	mapstructure.Decode(data, &subscription)
	payload := CostAndUsageReportPayload{
		Payload: subscription,
	}
	return &payload, err
}

func (c *ClientGraphqlClient) CreateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportPayload, error) {

	query := CreateCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"bucketName":            subscription.BucketName,
		"bucketRegion":          subscription.BucketRegion,
		"bucketPathPrefix":      subscription.BucketPathPrefix,
		"reportName":            subscription.ReportName,
		"timeUnit":              subscription.TimeUnit,
		"readerMode":            subscription.AccessConfig.ReaderMode,
		"assumeRoleArn":         subscription.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  subscription.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": subscription.AccessConfig.AssumeRoleSessionName,
	}

	var payload CostAndUsageReportPayload

	data, err := c.graphql(query, variables, "data.awsCreateCurSubscription")
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphqlClient) UpdateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportPayload, error) {

	query := UpdateCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"id":                    subscription.ID,
		"bucketName":            subscription.BucketName,
		"bucketRegion":          subscription.BucketRegion,
		"bucketPathPrefix":      subscription.BucketPathPrefix,
		"reportName":            subscription.ReportName,
		"timeUnit":              subscription.TimeUnit,
		"readerMode":            subscription.AccessConfig.ReaderMode,
		"assumeRoleArn":         subscription.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  subscription.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": subscription.AccessConfig.AssumeRoleSessionName,
	}

	var payload CostAndUsageReportPayload

	data, err := c.graphql(query, variables, "data.awsUpdateCurSubscription")
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphqlClient) DeleteCostAndUsageReportSubscription(id string) error {
	var query = DestroyCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"id": id,
	}

	_, err := c.graphql(query, variables, "data.awsDeleteCurSubscription")
	return err
}
