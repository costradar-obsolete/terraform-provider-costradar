package costradar

var GetCostAndUsageReportSubscriptionQuery = `
		query ($id: String!)
			{
			  awsCURSubscription(id: $id) {
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
