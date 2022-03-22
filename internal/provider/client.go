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

type TeamPayload struct {
	Status  bool   `json:"success"`
	Error   string `json:"error"`
	Payload Team   `json:"result"`
}

type WorkloadPayload struct {
	Status  bool     `json:"success"`
	Error   string   `json:"error"`
	Payload Workload `json:"result"`
}

type WorkloadSetPayload struct {
	Status  bool              `json:"success"`
	Error   string            `json:"error"`
	Payload []WorkloadSetItem `json:"result"`
}

type TeamMemberSetPayload struct {
	Status  bool                `json:"success"`
	Error   string              `json:"error"`
	Payload []TeamMemberSetItem `json:"result"`
}

type UserPayload struct {
	Status  bool   `json:"success"`
	Error   string `json:"error"`
	Payload User   `json:"result"`
}

type UserIdentitySetPayload struct {
	Status  bool                  `json:"success"`
	Error   string                `json:"error"`
	Payload []UserIdentitySetItem `json:"result"`
}

type AccountPayload struct {
	Status  bool    `json:"success"`
	Error   string  `json:"error"`
	Payload Account `json:"result"`
}

type TenantPayload struct {
	Status  bool   `json:"success"`
	Error   string `json:"error"`
	Payload Tenant `json:"result"`
}

type IntegrationConfig struct {
	IntegrationRoleArn string `json:"integrationRoleArn"`
	IntegrationSqsUrl  string `json:"integrationSqsUrl"`
	IntegrationSqsArn  string `json:"integrationSqsArn"`
}

type Workload struct {
	Name        string                 `json:"name"`
	ID          string                 `json:"id"`
	Description string                 `json:"description"`
	Owners      []interface{}          `json:"owners"`
	Tags        map[string]interface{} `json:"tags"`
}

type Team struct {
	Name        string                 `json:"name"`
	ID          string                 `json:"id"`
	Description string                 `json:"description"`
	Tags        map[string]interface{} `json:"tags"`
}

type WorkloadSetItem struct {
	ServiceVendor string `json:"serviceVendor"`
	ResourceId    string `json:"resourceId"`
}

type TeamMemberSetItem struct {
	Email string `json:"email"`
}

type User struct {
	ID       string                 `json:"id"`
	Email    string                 `json:"email"`
	Name     string                 `json:"name"`
	Initials string                 `json:"initials"`
	IconUrl  string                 `json:"iconUrl"`
	Tags     map[string]interface{} `json:"tags"`
}

type UserIdentitySetItem struct {
	ServiceVendor string `json:"serviceVendor"`
	Identity      string `json:"identity"`
}

type Account struct {
	ID           string                 `json:"id"`
	AccountId    string                 `json:"accountId"`
	Alias        string                 `json:"alias"`
	Owners       []interface{}          `json:"owners"`
	AccessConfig AccessConfig           `json:"accessConfig"`
	Tags         map[string]interface{} `json:"tags"`
}

type TenantAuth struct {
	ClientId          string        `json:"clientId"`
	ClientSecret      string        `json:"clientSecret"`
	ServerMetadataUrl string        `json:"serverMetadataUrl"`
	ClientKwargs      interface{}   `json:"clientKwargs"`
	EmailDomains      []interface{} `json:"emailDomains"`
}

type Tenant struct {
	Description string     `json:"description"`
	Alias       string     `json:"alias"`
	IconUrl     string     `json:"iconUrl"`
	Auth        TenantAuth `json:"auth"`
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

	GetWorkload(id string) (*WorkloadPayload, error)
	CreateWorkload(workload Workload) (*WorkloadPayload, error)
	UpdateWorkload(workload Workload) (*WorkloadPayload, error)
	DeleteWorkload(id string) error

	GetWorkloadSet(workloadId string, setId string) (*WorkloadSetPayload, error)
	CreateWorkloadSet(workloadId string, setId string, workloadSet []WorkloadSetItem) (*WorkloadSetPayload, error)
	UpdateWorkloadSet(workloadId string, setId string, workloadSet []WorkloadSetItem) (*WorkloadSetPayload, error)
	DeleteWorkloadSet(workloadId string, setId string) error

	GetTeam(id string) (*TeamPayload, error)
	CreateTeam(team Team) (*TeamPayload, error)
	UpdateTeam(team Team) (*TeamPayload, error)
	DeleteTeam(id string) error

	GetTeamMemberSet(teamId string, setId string) (*TeamMemberSetPayload, error)
	CreateTeamMemberSet(teamId string, setId string, teamSet []TeamMemberSetItem) (*TeamMemberSetPayload, error)
	UpdateTeamMemberSet(teamId string, setId string, teamSet []TeamMemberSetItem) (*TeamMemberSetPayload, error)
	DeleteTeamMemberSet(teamId string, setId string) error

	GetUser(id string) (*UserPayload, error)
	CreateUser(email string, user User) (*UserPayload, error)
	UpdateUser(user User) (*UserPayload, error)
	DeleteUser(id string) error

	GetUserIdentitySet(userId string, setId string) (*UserIdentitySetPayload, error)
	CreateUserIdentitySet(userId string, setId string, identitySet []UserIdentitySetItem) (*UserIdentitySetPayload, error)
	UpdateUserIdentitySet(userId string, setId string, identitySet []UserIdentitySetItem) (*UserIdentitySetPayload, error)
	DeleteUserIdentitySet(userId string, setId string) error

	GetAwsAccount(id string) (*AccountPayload, error)
	CreateAwsAccount(account Account) (*AccountPayload, error)
	UpdateAwsAccount(account Account) (*AccountPayload, error)
	DeleteAwsAccount(id string) error

	GetTenant() (*TenantPayload, error)
	UpdateTenant(tenant Tenant) (*TenantPayload, error)

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

func (c *ClientGraphql) GetWorkload(id string) (*WorkloadPayload, error) {
	query := GetWorkloadQuery

	variables := map[string]interface{}{
		"workloadId": id,
	}

	data, err := c.graphql(query, variables, "data.getWorkload")
	if err != nil {
		return nil, err
	}

	payload := WorkloadPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateWorkload(workload Workload) (*WorkloadPayload, error) {
	query := CreateWorkloadQuery
	variables := map[string]interface{}{
		"name":        workload.Name,
		"description": workload.Description,
		"owners":      workload.Owners,
		"tags":        workload.Tags,
	}

	data, err := c.graphql(query, variables, "data.createWorkload")
	if err != nil {
		return nil, err
	}
	payload := WorkloadPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateWorkload(workload Workload) (*WorkloadPayload, error) {
	query := UpdateWorkloadQuery
	variables := map[string]interface{}{
		"workloadId":  workload.ID,
		"name":        workload.Name,
		"description": workload.Description,
		"owners":      workload.Owners,
		"tags":        workload.Tags,
	}

	data, err := c.graphql(query, variables, "data.updateWorkload")

	if err != nil {
		return nil, err
	}
	payload := WorkloadPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteWorkload(id string) error {
	query := DeleteWorkloadQuery
	variables := map[string]interface{}{
		"workloadId": id,
	}

	_, err := c.graphql(query, variables, "data.deleteWorkload")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetWorkloadSet(workloadId string, setId string) (*WorkloadSetPayload, error) {
	query := GetWorkloadResourceSetQuery
	variables := map[string]interface{}{
		"workloadId": workloadId,
		"setId":      setId,
	}

	data, err := c.graphql(query, variables, "data.getWorkloadResourceSet")
	if err != nil {
		return nil, err
	}

	payload := WorkloadSetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateWorkloadSet(workloadId string, setId string, workloadSet []WorkloadSetItem) (*WorkloadSetPayload, error) {
	query := CreateWorkloadResourceSetQuery
	variables := map[string]interface{}{
		"workloadId": workloadId,
		"setId":      setId,
		"resources":  workloadSet,
	}

	data, err := c.graphql(query, variables, "data.createWorkloadResourceSet")
	if err != nil {
		return nil, err
	}
	payload := WorkloadSetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateWorkloadSet(workloadId string, setId string, workloadSet []WorkloadSetItem) (*WorkloadSetPayload, error) {
	query := UpdateWorkloadResourceSetQuery
	variables := map[string]interface{}{
		"workloadId": workloadId,
		"setId":      setId,
		"resources":  workloadSet,
	}

	data, err := c.graphql(query, variables, "data.updateWorkloadResourceSet")

	if err != nil {
		return nil, err
	}
	payload := WorkloadSetPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteWorkloadSet(workloadId string, setId string) error {
	query := DeleteWorkloadResourceSetQuery
	variables := map[string]interface{}{
		"workloadId": workloadId,
		"setId":      setId,
	}

	_, err := c.graphql(query, variables, "data.deleteWorkloadResourceSet")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetTeam(id string) (*TeamPayload, error) {
	query := GetTeamQuery

	variables := map[string]interface{}{
		"teamId": id,
	}

	data, err := c.graphql(query, variables, "data.getTeam")
	if err != nil {
		return nil, err
	}

	payload := TeamPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateTeam(team Team) (*TeamPayload, error) {
	query := CreateTeamQuery
	variables := map[string]interface{}{
		"name":        team.Name,
		"description": team.Description,
		"tags":        team.Tags,
	}

	data, err := c.graphql(query, variables, "data.createTeam")
	if err != nil {
		return nil, err
	}
	payload := TeamPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateTeam(team Team) (*TeamPayload, error) {
	query := UpdateTeamQuery
	variables := map[string]interface{}{
		"teamId":      team.ID,
		"name":        team.Name,
		"description": team.Description,
		"tags":        team.Tags,
	}

	data, err := c.graphql(query, variables, "data.updateTeam")

	if err != nil {
		return nil, err
	}
	payload := TeamPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteTeam(id string) error {
	query := DeleteTeamQuery
	variables := map[string]interface{}{
		"teamId": id,
	}

	_, err := c.graphql(query, variables, "data.deleteTeam")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetTeamMemberSet(teamId string, setId string) (*TeamMemberSetPayload, error) {
	query := GetTeamMemberSetQuery
	variables := map[string]interface{}{
		"teamId": teamId,
		"setId":  setId,
	}

	data, err := c.graphql(query, variables, "data.getTeamMemberSet")
	if err != nil {
		return nil, err
	}

	payload := TeamMemberSetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateTeamMemberSet(teamId string, setId string, memberSet []TeamMemberSetItem) (*TeamMemberSetPayload, error) {
	query := CreateTeamMemberSetQuery
	variables := map[string]interface{}{
		"teamId":      teamId,
		"setId":       setId,
		"teamMembers": memberSet,
	}

	data, err := c.graphql(query, variables, "data.createTeamMemberSet")
	if err != nil {
		return nil, err
	}
	payload := TeamMemberSetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateTeamMemberSet(teamId string, setId string, teamMemberSet []TeamMemberSetItem) (*TeamMemberSetPayload, error) {
	query := UpdateTeamMemberSetQuery
	variables := map[string]interface{}{
		"teamId":      teamId,
		"setId":       setId,
		"teamMembers": teamMemberSet,
	}

	data, err := c.graphql(query, variables, "data.updateTeamMemberSet")

	if err != nil {
		return nil, err
	}
	payload := TeamMemberSetPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteTeamMemberSet(teamId string, setId string) error {
	query := DeleteTeamMemberSetQuery
	variables := map[string]interface{}{
		"teamId": teamId,
		"setId":  setId,
	}

	_, err := c.graphql(query, variables, "data.deleteTeamMemberSet")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetUser(id string) (*UserPayload, error) {
	query := GetUserQuery

	variables := map[string]interface{}{
		"userId": id,
	}

	data, err := c.graphql(query, variables, "data.getUser")
	if err != nil {
		return nil, err
	}

	payload := UserPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateUser(email string, user User) (*UserPayload, error) {
	query := CreateUserQuery

	variables := map[string]interface{}{
		"email":    email,
		"name":     user.Name,
		"initials": user.Initials,
		"iconUrl":  user.IconUrl,
		"tags":     user.Tags,
	}

	data, err := c.graphql(query, variables, "data.createUser")
	if err != nil {
		return nil, err
	}
	payload := UserPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateUser(user User) (*UserPayload, error) {
	query := UpdateUserQuery
	variables := map[string]interface{}{
		"userId":   user.ID,
		"name":     user.Name,
		"initials": user.Initials,
		"iconUrl":  user.IconUrl,
		"tags":     user.Tags,
	}

	data, err := c.graphql(query, variables, "data.updateUser")

	if err != nil {
		return nil, err
	}
	payload := UserPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteUser(id string) error {
	query := DeleteUserQuery
	variables := map[string]interface{}{
		"userId": id,
	}

	_, err := c.graphql(query, variables, "data.deleteUser")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetUserIdentitySet(userId string, setId string) (*UserIdentitySetPayload, error) {
	query := GetUserIdentitySetQuery
	variables := map[string]interface{}{
		"userId": userId,
		"setId":  setId,
	}

	data, err := c.graphql(query, variables, "data.getUserIdentitySet")
	if err != nil {
		return nil, err
	}

	payload := UserIdentitySetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) CreateUserIdentitySet(userId string, setId string, identitySet []UserIdentitySetItem) (*UserIdentitySetPayload, error) {
	query := CreateUserIdentitySetQuery
	variables := map[string]interface{}{
		"userId":         userId,
		"setId":          setId,
		"userIdentities": identitySet,
	}

	data, err := c.graphql(query, variables, "data.createUserIdentitySet")
	if err != nil {
		return nil, err
	}
	payload := UserIdentitySetPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateUserIdentitySet(userId string, setId string, identitySet []UserIdentitySetItem) (*UserIdentitySetPayload, error) {
	query := UpdateUserIdentitySetQuery
	variables := map[string]interface{}{
		"userId":         userId,
		"setId":          setId,
		"userIdentities": identitySet,
	}

	data, err := c.graphql(query, variables, "data.updateUserIdentitySet")

	if err != nil {
		return nil, err
	}
	payload := UserIdentitySetPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteUserIdentitySet(userId string, setId string) error {
	query := DeleteUserIdentitySetQuery
	variables := map[string]interface{}{
		"userId": userId,
		"setId":  setId,
	}

	_, err := c.graphql(query, variables, "data.deleteUserIdentitySet")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetAwsAccount(id string) (*AccountPayload, error) {
	query := GetAwsAccountQuery
	variables := map[string]interface{}{
		"id": id,
	}

	data, err := c.graphql(query, variables, "data.awsGetAccount")
	if err != nil {
		return nil, err
	}

	account := Account{}
	err = json.Unmarshal(data, &account)

	payload := AccountPayload{
		Payload: account,
	}

	return &payload, err
}

func (c *ClientGraphql) CreateAwsAccount(account Account) (*AccountPayload, error) {
	query := CreateAwsAccountQuery
	variables := map[string]interface{}{
		"accountId":    account.AccountId,
		"alias":        account.Alias,
		"owners":       account.Owners,
		"accessConfig": account.AccessConfig,
		"tags":         account.Tags,
	}

	data, err := c.graphql(query, variables, "data.awsCreateAccount")
	if err != nil {
		return nil, err
	}
	payload := AccountPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateAwsAccount(account Account) (*AccountPayload, error) {
	query := UpdateAwsAccountQuery
	variables := map[string]interface{}{
		"id":           account.ID,
		"alias":        account.Alias,
		"owners":       account.Owners,
		"accessConfig": account.AccessConfig,
		"tags":         account.Tags,
	}

	data, err := c.graphql(query, variables, "data.awsUpdateAccount")

	if err != nil {
		return nil, err
	}
	payload := AccountPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}

func (c *ClientGraphql) DeleteAwsAccount(id string) error {
	query := DeleteAwsAccountQuery
	variables := map[string]interface{}{
		"id": id,
	}

	_, err := c.graphql(query, variables, "data.awsDeleteAccount")
	return err
}

// ------------------------------------------------------------------------------------------

func (c *ClientGraphql) GetTenant() (*TenantPayload, error) {
	query := GetTenantQuery

	variables := map[string]interface{}{}

	data, err := c.graphql(query, variables, "data.getTenant")
	if err != nil {
		return nil, err
	}

	payload := TenantPayload{}
	err = json.Unmarshal(data, &payload)

	return &payload, err
}

func (c *ClientGraphql) UpdateTenant(tenant Tenant) (*TenantPayload, error) {
	query := UpdateTenantQuery
	variables := map[string]interface{}{
		"alias":       tenant.Alias,
		"auth":        tenant.Auth,
		"description": tenant.Description,
		"iconUrl":     tenant.Alias,
	}

	data, err := c.graphql(query, variables, "data.updateTenant")

	if err != nil {
		return nil, err
	}
	payload := TenantPayload{}
	err = json.Unmarshal(data, &payload)
	return &payload, err
}
