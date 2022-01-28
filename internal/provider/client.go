package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	ID               string       `json:"id"`
	TrailName        string       `json:"trailName"`
	BucketName       string       `json:"bucketName"`
	BucketRegion     string       `json:"bucketRegion"`
	BucketPathPrefix string       `json:"bucketPathPrefix"`
	SourceTopicArn   string       `json:"sourceTopicArn"`
	AccessConfig     AccessConfig `json:"accessConfig"`
}

type IdentityResolver struct {
	LambdaArn    string       `json:"lambdaArn"`
	AccessConfig AccessConfig `json:"accessConfig"`
}

type IdentityResolverPayload struct {
	Status  bool             `json:"success"`
	Error   bool             `json:"error"`
	Payload IdentityResolver `json:"result"`
}

type CloudTrailSubscriptionPayload struct {
	Status  bool                   `json:"success"`
	Error   string                 `json:"error"`
	Payload CloudTrailSubscription `json:"result"`
}

type CostAndUsageReportSubscriptionPayload struct {
	Status  bool                           `json:"success"`
	Error   string                         `json:"error"`
	Payload CostAndUsageReportSubscription `json:"result"`
}

type IntegrationConfig struct {
	IntegrationRoleArn        string `json:"integrationRoleArn"`
	IntegrationRoleExternalId string `json:"integrationRoleExternalId"`
	CurSqsArn                 string `json:"curSqsArn"`
	CurSqsUrl                 string `json:"curSqsUrl"`
	CloudTrailSqsArn          string `json:"cloudtrailSqsArn"`
	CloudTrailSqsUrl          string `json:"cloudtrailSqsUrl"`
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

	GetIdentityResolver() (*IdentityResolverPayload, error)
	CreateIdentityResolver(resolver IdentityResolver) error
	UpdateIdentityResolver(resolver IdentityResolver) (*IdentityResolverPayload, error)
	DeleteIdentityResolver() error
	GetIntegrationConfig() (*IntegrationConfig, error)
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

func (c *ClientGraphql) graphql(query string, variables map[string]interface{}, dataPath string) (data []byte, err error) {

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
	data, err = json.Marshal(gjson.GetBytes(body, dataPath).Value())
	return data, err
}

func (c *ClientGraphql) GetCostAndUsageReportSubscription(id string) (*CostAndUsageReportSubscriptionPayload, error) {
	query := GetCostAndUsageReportSubscriptionQuery

	variables := map[string]interface{}{
		"id": id,
	}

	data, err := c.graphql(query, variables, "data.awsCurSubscription")
	if err != nil {
		return nil, err
	}

	subscription := CostAndUsageReportSubscription{}
	err = json.Unmarshal(data, &subscription)
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
	data, err := c.graphql(query, variables, "data.awsCreateCurSubscription")
	if err != nil {
		return nil, err
	}
	payload := CostAndUsageReportSubscriptionPayload{}
	err = json.Unmarshal(data, &payload)
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

	data, err := c.graphql(query, variables, "data.awsUpdateCurSubscription")
	if err != nil {
		return nil, err
	}

	payload := CostAndUsageReportSubscriptionPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteCostAndUsageReportSubscription(id string) error {
	var query = DeleteCostAndUsageReportSubscriptionQuery
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

	data, err := c.graphql(query, variables, "data.awsCloudTrailSubscription")
	if err != nil {
		return nil, err
	}

	subscription := CloudTrailSubscription{}
	err = json.Unmarshal(data, &subscription)

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

	data, err := c.graphql(query, variables, "data.awsCreateCloudTrailSubscription")
	if err != nil {
		return nil, err
	}
	payload := CloudTrailSubscriptionPayload{}
	err = json.Unmarshal(data, &payload)
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

	data, err := c.graphql(query, variables, "data.awsUpdateCloudTrailSubscription")
	if err != nil {
		return nil, err
	}
	payload := CloudTrailSubscriptionPayload{}
	err = json.Unmarshal(data, &payload)
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

func (c *ClientGraphql) GetIdentityResolver() (*IdentityResolverPayload, error) {
	query := GetIdentityResolverQuery

	data, err := c.graphql(query, nil, "data.awsIdentityResolver")
	if err != nil {
		return nil, err
	}

	resolver := IdentityResolver{}
	err = json.Unmarshal(data, &resolver)
	payload := IdentityResolverPayload{
		Payload: resolver,
	}
	return &payload, err
}

func (c *ClientGraphql) CreateIdentityResolver(resolver IdentityResolver) error {
	query := SetIdentityResolverQuery
	variables := map[string]interface{}{
		"lambdaArn":             resolver.LambdaArn,
		"readerMode":            resolver.AccessConfig.ReaderMode,
		"assumeRoleArn":         resolver.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  resolver.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": resolver.AccessConfig.AssumeRoleSessionName,
	}

	_, err := c.graphql(query, variables, "data.awsSetIdentityResolver")
	return err
}

func (c *ClientGraphql) UpdateIdentityResolver(resolver IdentityResolver) (*IdentityResolverPayload, error) {
	query := UpdateIdentityResolverQuery
	variables := map[string]interface{}{
		"lambdaArn":             resolver.LambdaArn,
		"readerMode":            resolver.AccessConfig.ReaderMode,
		"assumeRoleArn":         resolver.AccessConfig.AssumeRoleArn,
		"assumeRoleExternalId":  resolver.AccessConfig.AssumeRoleExternalId,
		"assumeRoleSessionName": resolver.AccessConfig.AssumeRoleSessionName,
	}

	data, err := c.graphql(query, variables, "data.awsUpdateIdentityResolver")
	if err != nil {
		return nil, err
	}
	payload := IdentityResolverPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteIdentityResolver() error {
	var query = DeleteIdentityResolverQuery

	_, err := c.graphql(query, nil, "data.awsDeleteIdentityResolver")
	return err
}

func (c *ClientGraphql) GetIntegrationConfig() (*IntegrationConfig, error) {
	var query = AwsIntegrationConfigQuery
	data, err := c.graphql(query, nil, "data.awsIntegrationConfig")
	if err != nil {
		return nil, err
	}
	integrationConfig := IntegrationConfig{}
	err = json.Unmarshal(data, &integrationConfig)
	return &integrationConfig, err
}
