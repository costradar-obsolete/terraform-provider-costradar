package costradar

var GetCostAndUsageReportSubscriptionQuery = `
query ($id: String!) {
	awsCurSubscription(id: $id) {
		id
		reportName
		bucketName
		bucketRegion
		bucketPathPrefix
		timeUnit
		accessConfig {
			readerMode
			assumeRoleArn
			assumeRoleExternalId
			assumeRoleSessionName
		}
	}
}
`

var CreateCostAndUsageReportSubscriptionQuery = `
mutation (
	$bucketName: String!
	$bucketRegion: String!
	$bucketPathPrefix: String
	$reportName: String!
	$timeUnit: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsCreateCurSubscription(input:{
		bucketName: $bucketName,
		bucketRegion: $bucketRegion,
		bucketPathPrefix: $bucketPathPrefix,
		reportName: $reportName,
		timeUnit: $timeUnit,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		status
		error
		payload {
			id
			reportName
			bucketName
			bucketRegion
			bucketPathPrefix
			timeUnit
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}
`

var UpdateCostAndUsageReportSubscriptionQuery = `
mutation (
	$id: String!
	$bucketName: String!
	$bucketRegion: String!
	$bucketPathPrefix: String
	$reportName: String!
	$timeUnit: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsUpdateCurSubscription(subscriptionId: $id, input:{
		bucketName: $bucketName,
		bucketRegion: $bucketRegion,
		bucketPathPrefix: $bucketPathPrefix,
		reportName: $reportName,
		timeUnit: $timeUnit,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
  	}){
		status
		error
		payload {
			id
			reportName
			bucketName
			bucketRegion
			bucketPathPrefix
			timeUnit
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}
`

var DestroyCostAndUsageReportSubscriptionQuery = `
mutation($id: String!) {
		awsDeleteCurSubscription(subscriptionId: $id){
		status
		error
		payload {
		  	id
		}
	}
}
`

var GetCloudTrailSubscriptionQuery = `
query ($id: String!) {
	awsCloudTrailSubscription(id: $id) {
		id
		tenant
		sourceArn
		subscriptionArn
		accountId
		bucketName
		accessConfig {
			readerMode
			assumeRoleArn
			assumeRoleExternalId
			assumeRoleSessionName
		}
	}
}
`

var CreateCloudTrailSubscriptionQuery = `
mutation (
	$sourceArn: String!
	$subscriptionArn: String
	$bucketName: String!
	$accountId: String
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsCreateCloudTrailSubscription(input:{
		sourceArn: $sourceArn,
		subscriptionArn: $subscriptionArn,
		bucketName: $bucketName,
		accountId: $accountId,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		status
		error
		payload {
			id
			sourceArn
			subscriptionArn
			bucketName
			accountId
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}
`

var UpdateCloudTrailSubscriptionQuery = `
mutation (
	$id: String!
	$sourceArn: String!
	$subscriptionArn: String
	$bucketName: String!
	$accountId: String
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsUpdateCloudTrailSubscription(subscriptionId: $id, input:{
		sourceArn: $sourceArn,
		subscriptionArn: $subscriptionArn,
		bucketName: $bucketName,
		accountId: $accountId
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
  	}){
		status
		error
		payload {
			id
			sourceArn
			subscriptionArn
			bucketName
			accountId
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}
`

var DeleteCloudTrailSubscriptionQuery = `
mutation($id: String!) {
		awsDeleteCloudTrailSubscription(subscriptionId: $id){
		status
		error
		payload {
		  	id
		}
	}
}
`
