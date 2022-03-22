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
	integrationSqsUrl
    integrationSqsArn
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

var CreateWorkloadQuery = `
mutation (
	$name: String!
	$description: String
	$owners: [String!]
	$tags: JSON
){
  createWorkload(
    input: {
      name: $name
 	  description: $description
      owners: $owners
      tags: $tags
  })
  {
    success
    error
    result {
	  id
      name
      tags
      description
      owners
    }
  }
}`

var GetWorkloadQuery = `
query ($workloadId: String!){
  getWorkload(workloadId: $workloadId){
    result {
	  id
      name
      tags
      description
      owners
    }
    success
    error
  }
}`

var UpdateWorkloadQuery = `
mutation (
	$workloadId: String!
	$name: String!
	$description: String
	$owners: [String!]
	$tags: JSON
){
  updateWorkload(
    workloadId: $workloadId,
    input: {
      owners: $owners
	  name: $name
	  description: $description
	  tags: $tags
  	}){
      result {
		id
        name
        tags
        description
        owners
      }
      success
      error
  }
}`

var DeleteWorkloadQuery = `
mutation ($workloadId: String!){
  deleteWorkload(workloadId: $workloadId){
      success
      error
  }
}`

var GetWorkloadResourceSetQuery = `
query ($workloadId: String!, $setId: String!) {
  getWorkloadResourceSet(workloadId: $workloadId, setId: $setId)
  {
    result {
	  setId
      serviceVendor
      resourceId
    }
	success
    error
  }
}
`

var CreateWorkloadResourceSetQuery = `
mutation (
  	$workloadId: String!,
	$setId: String!,
	$resources: [WorkloadResourceInput!]
  ) {
  createWorkloadResourceSet(
    workloadId: $workloadId,
    setId: $setId,
    input: $resources
  )
  {
    result {
	  setId
      resourceId
      serviceVendor
    }
    success
    error
  }
}
`

var UpdateWorkloadResourceSetQuery = `
mutation (
  	$workloadId: String!,
	$setId: String!,
	$resources: [WorkloadResourceInput!]
  ) {
  updateWorkloadResourceSet(
    workloadId: $workloadId,
    setId: $setId,
    input: $resources
  )
  {
    result {
	  setId
      resourceId
      serviceVendor
    }
    success
    error
  }
}
`

var DeleteWorkloadResourceSetQuery = `
mutation ($workloadId: String!, $setId: String!) {
  deleteWorkloadResourceSet(workloadId: $workloadId, setId: $setId) 
  {
    result {
	  setId
      serviceVendor
      resourceId
    }
    success
    error
  }
}`

var GetTeamQuery = `
query ($teamId: String!){
  getTeam(teamId: $teamId){
    result {
      name
      id
      tags
      description
    }
    success
    error
  }
}`

var CreateTeamQuery = `
mutation (
	$name: String!
	$description: String
	$tags: JSON
  ){
  createTeam(
    input: {
      name: $name
      tags: $tags
	  description: $description
    })
    {
      success
      error
      result {
        id
        tags
        description
        name
      }
    }
}`

var UpdateTeamQuery = `
mutation (
	$teamId: String!
	$name: String!
	$description: String
	$tags: JSON
  ){
  updateTeam(
	teamId: $teamId
    input: {
      name: $name
	  description: $description
      tags: $tags
    })
    {
      success
      error
      result {
        id
        tags
        description
        name
      }
    }
}`

var DeleteTeamQuery = `
mutation ($teamId: String!) {
  deleteTeam(teamId: $teamId){
    result {
      name
      id
      tags
      description
      members {
        email
        setId
      }
    }
    success
    error
  }
}`

var GetTeamMemberSetQuery = `
query ($teamId: String!, $setId: String!) {
  getTeamMemberSet(teamId: $teamId, setId: $setId)
  {
    success
    result {
	  setId
      email
    }
    error
  }
}
`

var CreateTeamMemberSetQuery = `
mutation ($teamId: String!, $setId: String!, $teamMembers: [TeamMemberInput!]) {
  createTeamMemberSet(
    teamId: $teamId,
    setId: $setId,
    input: $teamMembers)
  {
    success
    result {
      setId
      email
    }
    error
  }
}
`

var UpdateTeamMemberSetQuery = `
mutation ($teamId: String!, $setId: String!, $teamMembers: [TeamMemberInput!]) {
  updateTeamMemberSet(
    teamId: $teamId,
    setId: $setId,
    input: $teamMembers)
  {
    success
    result {
    	setId
      email
    }
    error
  }
}
`

var DeleteTeamMemberSetQuery = `
mutation ($teamId: String!, $setId: String!) {
  deleteTeamMemberSet(teamId: $teamId, setId: $setId) 
  {
    success
    result {
	  setId
      email
    }
    error
  }
}
`

var GetUserQuery = `
query ($userId: String!) {
	getUser(userId: $userId){
    success
    result {
      id
      name
      email
      initials
	  iconUrl
	  tags
    }
    error
  }
}
`

var CreateUserQuery = `
mutation (
	$email: String!,
	$name: String!,
	$initials: String,
	$iconUrl: String,
	$tags: JSON
	){
	createUser(
    email: $email
    input: {
      name: $name
	  initials: $initials
	  iconUrl: $iconUrl
      tags: $tags
    })
  {
    success
    result {
      id
      name
      email
      initials
	  iconUrl
	  tags
    }
    error
  }
}
`

var UpdateUserQuery = `
mutation (
	$userId: String!,
	$name: String!,
	$initials: String,
	$iconUrl: String,
	$tags: JSON
	){
  	updateUser(userId: $userId,
  	input: {
      name: $name
	  initials: $initials
	  iconUrl: $iconUrl
      tags: $tags
    })
  {
    success
    result {
      id
      name
      email
      initials
	  iconUrl
	  tags
    }
    error
  }
}
`

var DeleteUserQuery = `
mutation ($userId: String!) {
  deleteUser(userId: $userId) {
    success
    result {
      id
      name
      email
      initials
	  iconUrl
	  tags
    }
    error
  }
}
`

var GetUserIdentitySetQuery = `
query ($userId: String!, $setId: String!) {
  getUserIdentitySet(userId: $userId, setId: $setId)
  {
    success
    result {
	  setId
      serviceVendor
      identity
    }
    error
  }
}
`

var CreateUserIdentitySetQuery = `
mutation ($userId: String!, $setId: String!, $userIdentities: [UserIdentityInput!]) {
  createUserIdentitySet(
    userId: $userId,
    setId: $setId,
    input: $userIdentities)
  {
    success
    result {
	  setId
      serviceVendor
      identity
    }
    error
  }
}
`

var UpdateUserIdentitySetQuery = `
mutation ($userId: String!, $setId: String!, $userIdentities: [UserIdentityInput!]) {
  updateUserIdentitySet(
    userId: $userId,
    setId: $setId,
    input: $userIdentities)
  {
    success
    result {
	  setId
      serviceVendor
      identity
    }
    error
  }
}
`

var DeleteUserIdentitySetQuery = `
mutation ($userId: String!, $setId: String!) {
  deleteUserIdentitySet(userId: $userId, setId: $setId)
  {
    success
    error
    result {
      setId
      serviceVendor
      identity
    }
  }
}
`

var GetAwsAccountQuery = `
query ($id: String!){
  awsGetAccount(id: $id){
	id
	accountId
	alias
	tags
	owners
	accessConfig {
	  readerMode
	  assumeRoleArn
	  assumeRoleExternalId
	  assumeRoleSessionName
	}
  }
}
`

var CreateAwsAccountQuery = `
mutation (
  	$accountId: String!,
	$alias: String!
	$owners: [String!]
    $accessConfig: AWSAccessConfigInput!
	$tags: JSON
  ){
  awsCreateAccount(
    accountId: $accountId,
    input: {
      alias: $alias
      owners: $owners
      accessConfig: $accessConfig
	  tags: $tags
  }){
    result {
      id
      accountId
      alias
      owners
	  tags
      accessConfig {
        readerMode
		assumeRoleArn
		assumeRoleExternalId
		assumeRoleSessionName
      }
    }
    success
  }
}
`

var UpdateAwsAccountQuery = `
mutation (
	$id: String!
	$alias: String!
	$owners: [String!]
    $accessConfig: AWSAccessConfigInput!
	$tags: JSON
  ){
  awsUpdateAccount(
	id: $id, 
	input: {
	  alias: $alias
	  owners: $owners
	  accessConfig: $accessConfig
	  tags: $tags
  }){
    result {
      id
      accountId
      alias
      owners
	  tags
      accessConfig {
        readerMode
		assumeRoleArn
		assumeRoleExternalId
		assumeRoleSessionName
      }
    }
    success
  }
}
`

var DeleteAwsAccountQuery = `
mutation ($id: String!) {
  awsDeleteAccount(id: $id){
    result {
      id
      accountId
      alias
      owners
      accessConfig {
        readerMode
		assumeRoleArn
		assumeRoleExternalId
		assumeRoleSessionName
      }
    }
    success
  }
}
`

var GetTenantQuery = `
query {
  getTenant
  {
    result {
      alias
      description
      iconUrl
      auth {
        clientId
        clientSecret
        serverMetadataUrl
        clientKwargs
        emailDomains
      }
    }
    success
  }
}`

var UpdateTenantQuery = `
mutation (
	$alias: String!
	$auth: OIDCAuthConfigInput
	$iconUrl: String
	$description: String
  ){
  updateTenant(
	input: {
   	  alias: $alias
	  iconUrl: $iconUrl
	  description: $description
	  auth: $auth
	}
  ){
	result {
	  alias
	  description
	  iconUrl
	  auth {
		clientId
		clientSecret
		serverMetadataUrl
		clientKwargs
		emailDomains
	  }
	}
	success
  }
}`
