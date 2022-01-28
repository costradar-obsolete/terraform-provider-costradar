package provider

var GetCostAndUsageReportSubscriptionQuery = `
query ($id: String!) {
	awsCurSubscription(id: $id) {
		id
		reportName
		bucketName
		bucketRegion
		bucketPathPrefix
		sourceTopicArn
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
	$sourceTopicArn: String!
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
		sourceTopicArn: $sourceTopicArn,
		timeUnit: $timeUnit,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		success
		error
		result {
			id
            sourceTopicArn
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
	$sourceTopicArn: String!
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
	awsUpdateCurSubscription(id: $id, input:{
		bucketName: $bucketName,
		bucketRegion: $bucketRegion,
		bucketPathPrefix: $bucketPathPrefix,
		reportName: $reportName,
		sourceTopicArn: $sourceTopicArn,
		timeUnit: $timeUnit,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
  	}){
		success
		error
		result {
			id
			sourceTopicArn
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

var DeleteCostAndUsageReportSubscriptionQuery = `
mutation($id: String!) {
	awsDeleteCurSubscription(id: $id) {
		success
		error
		result {
			id
		}
	}
}
`

var GetCloudTrailSubscriptionQuery = `
query ($id: String!) {
	awsCloudTrailSubscription(id: $id) {
		id
		trailName
      	bucketName
      	bucketRegion
      	bucketPathPrefix
      	sourceTopicArn
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
	$trailName: String!
	$bucketName: String!
	$bucketRegion: String!
    $bucketPathPrefix: String
	$sourceTopicArn: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsCreateCloudTrailSubscription(input:{
		trailName: $trailName,
      	bucketName: $bucketName,
      	bucketRegion: $bucketRegion,
      	bucketPathPrefix: $bucketPathPrefix,
      	sourceTopicArn: $sourceTopicArn,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		success
		error
		result {
			id
			trailName
			bucketName
			bucketRegion
			bucketPathPrefix
			sourceTopicArn
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
	$trailName: String!
	$bucketName: String!
	$bucketRegion: String!
    $bucketPathPrefix: String
	$sourceTopicArn: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsUpdateCloudTrailSubscription(id: $id, input: {
		trailName: $trailName,
      	bucketName: $bucketName,
      	bucketRegion: $bucketRegion,
      	bucketPathPrefix: $bucketPathPrefix,
      	sourceTopicArn: $sourceTopicArn,
		accessConfig:{
			readerMode: $readerMode
			assumeRoleArn: $assumeRoleArn
			assumeRoleExternalId: $assumeRoleExternalId
			assumeRoleSessionName: $assumeRoleSessionName
		}
  	}) {
		success
		error
		result {
			id
			trailName
			bucketName
			bucketRegion
			bucketPathPrefix
			sourceTopicArn
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
	awsDeleteCloudTrailSubscription(id: $id) {
	success
	error
	result {
		id
	}}
}
`

var AwsIntegrationConfigQuery = `
query {
  awsIntegrationConfig {
    integrationRoleArn
    integrationRoleExternalId
    curSqsArn
    curSqsUrl
    cloudtrailSqsArn
    cloudtrailSqsUrl
  }
}
`

var GetIdentityResolverQuery = `
query {
  	awsIdentityResolver{
		lambdaArn
		accessConfig {
			readerMode
			assumeRoleArn
			assumeRoleExternalId
			assumeRoleSessionName
		}
  	}
}`

var SetIdentityResolverQuery = `
mutation (
	$lambdaArn: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsSetIdentityResolver(input:{
		lambdaArn: $lambdaArn,
		accessConfig:{
			readerMode: $readerMode,
			assumeRoleArn: $assumeRoleArn,
			assumeRoleExternalId: $assumeRoleExternalId,
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		success
		error
		result {
			lambdaArn
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}`

var UpdateIdentityResolverQuery = `
mutation (
	$lambdaArn: String!
	$readerMode: ReaderMode!
	$assumeRoleArn: String
	$assumeRoleExternalId: String
	$assumeRoleSessionName: String
){
	awsUpdateIdentityResolver(input:{
		lambdaArn: $lambdaArn,
		accessConfig:{
			readerMode: $readerMode,
			assumeRoleArn: $assumeRoleArn,
			assumeRoleExternalId: $assumeRoleExternalId,
			assumeRoleSessionName: $assumeRoleSessionName
		}
	}){
		success
		error
		result {
			lambdaArn
			accessConfig {
				readerMode
				assumeRoleArn
				assumeRoleExternalId
				assumeRoleSessionName
			}
		}
	}
}`

var DeleteIdentityResolverQuery = `
mutation{
	awsDeleteIdentityResolver {
		success
		error
	}
}`
