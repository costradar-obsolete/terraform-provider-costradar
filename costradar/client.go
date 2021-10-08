package costradar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
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
	SourceTopicArn   string       `json:"sourceTopicArn"`
	AccessConfig     AccessConfig `json:"accessConfig"`
}

type CloudTrailSubscription struct {
	ID                         string `json:"id"`
	TrailName                  string `json:"trailName"`
	BucketName                 string `json:"bucketName"`
	BucketRegion               string `json:"bucketRegion"`
	BucketPathPrefix           string `json:"bucketPathPrefix"`
	SourceTopicArn             string `json:"sourceTopicArn"`
	AccessConfig     AccessConfig     `json:"accessConfig"`
}

type CloudTrailSubscriptionPayload struct {
	Status  bool                   `json:"status"`
	Error   string                 `json:"error"`
	Payload CloudTrailSubscription `json:"payload"`
}

type CostAndUsageReportSubscriptionPayload struct {
	Status  bool                           `json:"status"`
	Error   string                         `json:"error"`
	Payload CostAndUsageReportSubscription `json:"payload"`
}

type Client interface {
	GetCostAndUsageReportSubscription(id string) (*CostAndUsageReportSubscriptionPayload, error)
	CreateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportSubscriptionPayload, error)
	UpdateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportSubscriptionPayload, error)
	DeleteCostAndUsageReportSubscription(id string) error

	GetCloudTrailSubscription(id string) (*CloudTrailSubscriptionPayload, error)
	CreateCloudTrailSubscription(subscription CloudTrailSubscription) (*CloudTrailSubscriptionPayload, error)
	UpdateCloudTrailSubscription(subscription CloudTrailSubscription) (*CloudTrailSubscriptionPayload, error)
	DeleteCloudTrailSubscription(id string) error
}

type ClientGraphql struct {
	endpoint   string
	token      string
	httpClient *http.Client
}

func NewCostRadarClient(endpoint, token string) Client {
	return &ClientGraphql{
		endpoint:   endpoint,
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *ClientGraphql) graphql(query string, variables map[string]interface{}, dataPath string) (data interface{}, err error) {

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

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Status code: " + strconv.Itoa(resp.StatusCode) + ". Message: " + gjson.GetBytes(body, "error").String())
		return nil, err
	}

	errorMessage := getErrorFromBody(body, dataPath)

	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	data = gjson.GetBytes(body, dataPath).Value()
	return data, err
}

func (c *ClientGraphql) GetCostAndUsageReportSubscription(id string) (*CostAndUsageReportSubscriptionPayload, error) {
	query := GetCostAndUsageReportSubscriptionQuery

	variables := map[string]interface{}{
		"id": id,
	}
	subscription := CostAndUsageReportSubscription{}

	data, err := c.graphql(query, variables, "data.awsCurSubscription")
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(data, &subscription)
	payload := CostAndUsageReportSubscriptionPayload{
		Payload: subscription,
	}
	return &payload, err
}

func (c *ClientGraphql) CreateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportSubscriptionPayload, error) {

	query := CreateCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"sourceTopicArn":        subscription.SourceTopicArn,
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

	var payload CostAndUsageReportSubscriptionPayload

	data, err := c.graphql(query, variables, "data.awsCreateCurSubscription")
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) UpdateCostAndUsageReportSubscription(subscription CostAndUsageReportSubscription) (*CostAndUsageReportSubscriptionPayload, error) {

	query := UpdateCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"id":                    subscription.ID,
		"bucketName":            subscription.BucketName,
		"bucketRegion":          subscription.BucketRegion,
		"bucketPathPrefix":      subscription.BucketPathPrefix,
		"reportName":            subscription.ReportName,
		"timeUnit":              subscription.TimeUnit,
		"sourceTopicArn":        subscription.SourceTopicArn,
		"readerMode":            subscription.AccessConfig.ReaderMode,
		"assumeRoleArn":         subscription.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  subscription.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": subscription.AccessConfig.AssumeRoleSessionName,
	}

	var payload CostAndUsageReportSubscriptionPayload

	data, err := c.graphql(query, variables, "data.awsUpdateCurSubscription")
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteCostAndUsageReportSubscription(id string) error {
	var query = DestroyCostAndUsageReportSubscriptionQuery
	variables := map[string]interface{}{
		"id": id,
	}

	_, err := c.graphql(query, variables, "data.awsDeleteCurSubscription")
	return err
}

func (c *ClientGraphql) GetCloudTrailSubscription(id string) (*CloudTrailSubscriptionPayload, error) {
	var query = GetCloudTrailSubscriptionQuery

	variables := map[string]interface{}{
		"id": id,
	}

	subscription := CloudTrailSubscription{}

	data, err := c.graphql(query, variables, "data.awsCloudTrailSubscription")
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(data, &subscription)
	payload := CloudTrailSubscriptionPayload{
		Payload: subscription,
	}

	return &payload, err
}

func (c *ClientGraphql) CreateCloudTrailSubscription(subscription CloudTrailSubscription) (*CloudTrailSubscriptionPayload, error) {

	query := CreateCloudTrailSubscriptionQuery
	variables := map[string]interface{}{
		"id":                    subscription.ID,
		"bucketName":            subscription.BucketName,
		"bucketRegion":          subscription.BucketRegion,
		"bucketPathPrefix":      subscription.BucketPathPrefix,
		"trailName":             subscription.TrailName,
		"sourceTopicArn":        subscription.SourceTopicArn,
		"readerMode":            subscription.AccessConfig.ReaderMode,
		"assumeRoleArn":         subscription.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  subscription.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": subscription.AccessConfig.AssumeRoleSessionName,
	}

	var payload CloudTrailSubscriptionPayload

	data, err := c.graphql(query, variables, "data.awsCreateCloudTrailSubscription")
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) UpdateCloudTrailSubscription(subscription CloudTrailSubscription) (*CloudTrailSubscriptionPayload, error) {

	query := UpdateCloudTrailSubscriptionQuery
	variables := map[string]interface{}{
		"id":                    subscription.ID,
		"bucketName":            subscription.BucketName,
		"bucketRegion":          subscription.BucketRegion,
		"bucketPathPrefix":      subscription.BucketPathPrefix,
		"trailName":             subscription.TrailName,
		"sourceTopicArn":        subscription.SourceTopicArn,
		"readerMode":            subscription.AccessConfig.ReaderMode,
		"assumeRoleArn":         subscription.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  subscription.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": subscription.AccessConfig.AssumeRoleSessionName,
	}

	var payload CloudTrailSubscriptionPayload

	data, err := c.graphql(query, variables, "data.awsUpdateCloudTrailSubscription")
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteCloudTrailSubscription(id string) error {
	var query = DeleteCloudTrailSubscriptionQuery
	variables := map[string]interface{}{
		"id": id,
	}

	_, err := c.graphql(query, variables, "data.awsDeleteCloudTrailSubscription")
	return err
}
