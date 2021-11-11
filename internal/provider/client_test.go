package provider

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

var _ = Describe("Costradar http client", func() {
	var server *ghttp.Server
	var costradarClient Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		costradarClient = NewCostRadarClient(server.URL()+"/graphql", "api_Uu94jTpg7UOw5vDFcHUsqFo0pkYoZiL8")
	})

	AfterEach(func() {
		server.Close()
	})

	It("Get cost and usage report subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, getCostAndUsageReportSubscriptionSuccess()),
		))
		subscription, err := costradarClient.GetCostAndUsageReportSubscription("test")
		Expect(err).To(BeNil())
		Expect(subscription.Payload.ID).ToNot(Equal(""))
		Expect(subscription.Payload.ID).To(Equal("61487ac0151407505898f1a3"))
		Expect(subscription.Payload.ReportName).To(Equal("report_name"))
		Expect(subscription.Payload.BucketName).To(Equal("bucket_name"))
		Expect(subscription.Payload.BucketRegion).To(Equal("bucket_region"))
		Expect(subscription.Payload.BucketPathPrefix).To(Equal("xxx!!"))
		Expect(subscription.Payload.TimeUnit).To(Equal("hour"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("assumeRole"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleArn).To(Equal("assume_role_arn_value"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleExternalId).To(Equal("assume_role_external_id_value"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleSessionName).To(Equal("assume_role_session_name_value"))

	})

	It("Update cost and usage report subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, updateCostAndUsageReportSubscriptionSuccess()),
		))
		updateSubscription := CostAndUsageReportSubscription{
			ID:               "61487ac0151407505898f1a3",
			ReportName:       "report",
			BucketName:       "bucket",
			BucketRegion:     "region",
			BucketPathPrefix: "path_prefix",
			TimeUnit:         "daily",
			AccessConfig: AccessConfig{
				ReaderMode: "direct",
			},
		}
		subscription, err := costradarClient.UpdateCostAndUsageReportSubscription(updateSubscription)
		Expect(err).To(BeNil())
		Expect(subscription).ToNot(BeNil())
		Expect(subscription.Payload.ID).To(Equal("61487ac0151407505898f1a3"))
		Expect(subscription.Payload.ReportName).To(Equal("report"))
		Expect(subscription.Payload.BucketName).To(Equal("bucket"))
		Expect(subscription.Payload.BucketRegion).To(Equal("region"))
		Expect(subscription.Payload.BucketPathPrefix).To(Equal("path_prefix"))
		Expect(subscription.Payload.TimeUnit).To(Equal("daily"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("direct"))
	})

	It("Cost and usage report subscription not found", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, getCostAndUsageReportSubscriptionNotExist()),
		))

		subscription, err := costradarClient.GetCostAndUsageReportSubscription("test")
		Expect(err).ToNot(BeNil())
		Expect(subscription).To(BeNil())
	})

	It("Create cost and usage report subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, createCostAndUsageReportSubscriptionSuccess()),
		))
		createSubscription := CostAndUsageReportSubscription{
			ReportName:       "report",
			BucketName:       "bucket",
			BucketRegion:     "region",
			BucketPathPrefix: "path_prefix",
			TimeUnit:         "daily",
			AccessConfig: AccessConfig{
				ReaderMode:            "assumeRole",
				AssumeRoleArn:         "ARN",
				AssumeRoleExternalId:  "ID",
				AssumeRoleSessionName: "SESSIONNAME",
			},
		}
		subscription, err := costradarClient.CreateCostAndUsageReportSubscription(createSubscription)
		Expect(err).To(BeNil())
		Expect(subscription).ToNot(BeNil())
		Expect(subscription.Payload.ID).ToNot(BeNil())
		Expect(subscription.Payload.ReportName).To(Equal("SomeName"))
		Expect(subscription.Payload.BucketName).To(Equal("bucketName"))
		Expect(subscription.Payload.BucketRegion).To(Equal("bucketRegion"))
		Expect(subscription.Payload.BucketPathPrefix).To(Equal("bucketPathPrefix"))
		Expect(subscription.Payload.TimeUnit).To(Equal("hour"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("assumeRole"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleArn).To(Equal("ARN"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleExternalId).To(Equal("ID"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleSessionName).To(Equal("SESSIONNAME"))
	})

	It("Delete cost report subscription success", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, deleteCostAndUsageReportSubscriptionSuccess()),
		))

		err := costradarClient.DeleteCostAndUsageReportSubscription("123")
		Expect(err).To(BeNil())
	})

	It("Delete cost report subscription does not exist", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, deleteCostAndUsageReportSubscriptionNotFound()),
		))

		err := costradarClient.DeleteCostAndUsageReportSubscription("123")
		Expect(err).NotTo(BeNil())
	})

	It("Create Cloud Trail subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, createCloudTrailSubscriptionSuccess()),
		))
		createSubscription := CloudTrailSubscription{
			TrailName:      "trail-name",
			SourceTopicArn: "topic-arn",
			BucketName:     "bucket",
			BucketRegion:   "region",
			AccessConfig: AccessConfig{
				ReaderMode:            "assumeRole",
				AssumeRoleArn:         "ARN",
				AssumeRoleExternalId:  "ID",
				AssumeRoleSessionName: "SESSIONNAME",
			},
		}
		subscription, err := costradarClient.CreateCloudTrailSubscription(createSubscription)
		Expect(err).To(BeNil())
		Expect(subscription).ToNot(BeNil())
		Expect(subscription.Payload.ID).To(Equal("614ae6fc151407505898f1af"))
		Expect(subscription.Payload.TrailName).To(Equal("trail"))
		Expect(subscription.Payload.SourceTopicArn).To(Equal("topic-arn"))
		Expect(subscription.Payload.BucketName).To(Equal("bucket"))
		Expect(subscription.Payload.BucketRegion).To(Equal("region"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("assumeRole"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleArn).To(Equal("ARN"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleExternalId).To(Equal("ID"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleSessionName).To(Equal("SESSIONNAME"))
	})

	It("Get Cloud Trail subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, getCloudTrailSubscriptionSuccess()),
		))

		subscription, err := costradarClient.GetCloudTrailSubscription("123")
		Expect(err).To(BeNil())
		Expect(subscription).ToNot(BeNil())
		Expect(subscription.Payload.ID).To(Equal("614ae6fc151407505898f1af"))
		Expect(subscription.Payload.SourceTopicArn).To(Equal("topic-arn"))
		Expect(subscription.Payload.BucketName).To(Equal("bucket"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("assumeRole"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleArn).To(Equal("ARN"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleExternalId).To(Equal("ID"))
		Expect(subscription.Payload.AccessConfig.AssumeRoleSessionName).To(Equal("SESSIONNAME"))
	})

	It("Update Cloud Trail subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, updateCloudTrailSubscriptionSuccess()),
		))

		updateSubscription := CloudTrailSubscription{
			TrailName:      "trail-name",
			SourceTopicArn: "topic-arn",
			BucketName:     "bucket",
			BucketRegion:   "region",
			AccessConfig: AccessConfig{
				ReaderMode: "direct",
			},
		}
		subscription, err := costradarClient.UpdateCloudTrailSubscription(updateSubscription)
		Expect(err).To(BeNil())
		Expect(subscription.Payload.ID).To(Equal("614ae6fc151407505898f1af"))
		Expect(subscription.Payload.TrailName).To(Equal("trail-u"))
		Expect(subscription.Payload.SourceTopicArn).To(Equal("topic-arn-u"))
		Expect(subscription.Payload.BucketRegion).To(Equal("region-u"))
		Expect(subscription.Payload.BucketName).To(Equal("bucket-u"))
		Expect(subscription.Payload.AccessConfig.ReaderMode).To(Equal("direct"))
	})

	It("Delete Cloud Trail subscription", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, deleteCloudTrailSubscriptionSuccess()),
		))

		err := costradarClient.DeleteCloudTrailSubscription("123")

		Expect(err).To(BeNil())
	})

	It("Delete Cloud Trail subscription does not found", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, deleteCloudTrailSubscriptionNotFound()),
		))

		err := costradarClient.DeleteCloudTrailSubscription("123")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("CloudTrail subscription not found."))
	})

	It("Get Cloud Trail subscription does not exist", func() {
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/graphql"),
			ghttp.RespondWith(http.StatusOK, getCloudTrailSubscriptionNotExist()),
		))

		_, err := costradarClient.GetCloudTrailSubscription("123")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("CloudTrailSubscription matching query does not exist. "))
	})
})

func createCostAndUsageReportSubscriptionSuccess() string {
	return fmt.Sprintf(`
		{
		  "data": {
			"awsCreateCurSubscription": {
			  "status": true,
			  "error": null,
			  "payload": {
				"id": "61489a2c151407505898f1a5",
				"reportName": "SomeName",
				"bucketName": "bucketName",
				"bucketRegion": "bucketRegion",
				"bucketPathPrefix": "bucketPathPrefix",
				"timeUnit": "hour",
				"accessConfig": {
				  "readerMode": "assumeRole",
				  "assumeRoleArn": "ARN",
				  "assumeRoleExternalId": "ID",
				  "assumeRoleSessionName": "SESSIONNAME"
				}
			  }
			}
		  }
		}
`)
}

func updateCostAndUsageReportSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsUpdateCurSubscription": {
		  "status": true,
		  "error": null,
		  "payload": {
			"id": "61487ac0151407505898f1a3",
			"reportName": "report",
			"bucketName": "bucket",
			"bucketRegion": "region",
			"bucketPathPrefix": "path_prefix",
			"timeUnit": "daily",
			"accessConfig": {
			  "readerMode": "direct",
			  "assumeRoleArn": null,
			  "assumeRoleExternalId": null,
			  "assumeRoleSessionName": null
			}
		  }
		}
	  }
	}`)
}

func getCostAndUsageReportSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsCurSubscription": {
		  "id": "61487ac0151407505898f1a3",
		  "reportName": "report_name",
		  "bucketName": "bucket_name",
		  "bucketRegion": "bucket_region",
		  "bucketPathPrefix": "xxx!!",
		  "timeUnit": "hour",
		  "accessConfig": {
			"readerMode": "assumeRole",
			"assumeRoleArn": "assume_role_arn_value",
			"assumeRoleExternalId": "assume_role_external_id_value",
			"assumeRoleSessionName": "assume_role_session_name_value"
		  }
		}
	  }
	}`)
}

func getCostAndUsageReportSubscriptionNotExist() string {
	return fmt.Sprintf(`
		{
		  "data": {
			"awsCurSubscription": null
		  },
		  "errors": [
			{
			  "message": "CostAndUsageReportSubscription matching query does not exist.",
			},
		  ],
		}
	`)
}

func deleteCostAndUsageReportSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsDeleteCurSubscription": {
		  "error": null,
		  "payload": null,
		  "status": true
		}
	  }
	}`)
}

func deleteCostAndUsageReportSubscriptionNotFound() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsDeleteCurSubscription": {
		  "error": "CUR subscription not found.",
		  "payload": null,
		  "status": false
		}
	  }
	}`)
}

func getCloudTrailSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsCloudTrailSubscription": {
		  "id": "614ae6fc151407505898f1af",
		  "sourceTopicArn": "topic-arn",
          "trailName": "trail",
		  "bucketName": "bucket",
		  "bucketRegion": "region",
		  "bucketPathPrefix": "prefix",
		  "accessConfig": {
			"readerMode": "assumeRole",
			"assumeRoleArn": "ARN",
			"assumeRoleExternalId": "ID",
			"assumeRoleSessionName": "SESSIONNAME"
		  }
		}
	  }
	}`)
}

func createCloudTrailSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsCreateCloudTrailSubscription": {
		  "status": true,
		  "error": null,
		  "payload": {
			"id": "614ae6fc151407505898f1af",
			"sourceTopicArn": "topic-arn",
            "trailName": "trail",
		    "bucketName": "bucket",
		    "bucketRegion": "region",
		    "bucketPathPrefix": "prefix",
			"accessConfig": {
			  "readerMode": "assumeRole",
			  "assumeRoleArn": "ARN",
			  "assumeRoleExternalId": "ID",
			  "assumeRoleSessionName": "SESSIONNAME"
			}
		  }
		}
	  }
	}`)
}

func updateCloudTrailSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsUpdateCloudTrailSubscription": {
		  "status": true,
		  "error": null,
		  "payload": {
			"id": "614ae6fc151407505898f1af",
			"sourceTopicArn": "topic-arn-u",
		    "trailName": "trail-u",
		    "bucketName": "bucket-u",
		    "bucketRegion": "region-u",
		    "bucketPathPrefix": "prefix-u",
			"accessConfig": {
			  "readerMode": "direct",
			  "assumeRoleArn": null,
			  "assumeRoleExternalId": null,
			  "assumeRoleSessionName": null
			}
		  }
		}
	  }
	}`)
}

func deleteCloudTrailSubscriptionSuccess() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsDeleteCloudTrailSubscription": {
		  "error": null,
		  "payload": null,
		  "status": true
		}
	  }
	}`)
}

func deleteCloudTrailSubscriptionNotFound() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsDeleteCloudTrailSubscription": {
		  "error": "CloudTrail subscription not found.",
		  "payload": null,
		  "status": false
		}
	  }
	}`)
}

func getCloudTrailSubscriptionNotExist() string {
	return fmt.Sprintf(`{
	  "data": {
		"awsCloudTrailSubscription": null
	  },
	  "errors": [
		{
		  "message": "CloudTrailSubscription matching query does not exist."
		}
	  ]
	}`)
}
