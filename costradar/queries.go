package costradar

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
		status
		error
		payload {
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
	awsUpdateCurSubscription(subscriptionId: $id, input:{
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
		status
		error
		payload {
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

var DestroyCostAndUsageReportSubscriptionQuery = `
mutation($id: String!) {
	awsDeleteCurSubscription(subscriptionId: $id) {
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
		status
		error
		payload {
			id
			tenant
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
	awsUpdateCloudTrailSubscription(subscriptionId: $id, input: {
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
		status
		error
		payload {
			id
			tenant
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
		awsDeleteCloudTrailSubscription(subscriptionId: $id) {
		status
		error
		payload {
		  	id
		}
	}
}
`
