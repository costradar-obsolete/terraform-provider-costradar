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

type CostAndUsageReportSubscription struct {
	ID               string                         `json:"id"`
	ReportName       string                         `json:"reportName"`
	BucketName       string                         `json:"bucketName"`
	BucketRegion     string                         `json:"bucketRegion"`
	BucketPathPrefix string                         `json:"bucketPathPrefix"`
	TimeUnit         string                         `json:"timeUnit"`
	AccessConfig     CostAndUsageReportAccessConfig `json:"accessConfig"`
}

type CostAndUsageReportAccessConfig struct {
	ReaderMode            string `json:"readerMode"`
	AssumeRoleArn         string `json:"assumeRoleArn"`
	AssumeRoleExternalId  string `json:"assumeRoleExternalId"`
	AssumeRoleSessionName string `json:"assumeRoleSessionName"`
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
	DestroyCostAndUsageReportSubscription(id string) error
}

type ClientGraphqlClient struct {
	graphqlEndpoint string
	accessToken     string
	httpClient      *http.Client
}

func NewCostRadarClient(graphqlEndpoint, accessToken string) Client {
	return &ClientGraphqlClient{
		graphqlEndpoint: graphqlEndpoint,
		accessToken:     accessToken,
		httpClient:      &http.Client{},
	}
}

func (c *ClientGraphqlClient) graphql(query string, variables map[string]interface{}, dataPath string) (data interface{}, err error) {

	payload := new(bytes.Buffer)

	json.NewEncoder(payload).Encode(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})

	req, _ := http.NewRequest("POST", c.graphqlEndpoint, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.accessToken))
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

func (c *ClientGraphqlClient) DestroyCostAndUsageReportSubscription(id string) error {
	var query = DestroyCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"id": id,
	}

	_, err := c.graphql(query, variables, "data.awsDeleteCurSubscription")
	return err
}
