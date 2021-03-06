package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccCloudTrailSubscription(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cloudtrail_subscription.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTrailSubscriptionTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "trail_name", "trail-name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "prefix"),
					resource.TestCheckResourceAttr(resourceName, "source_topic_arn", "topic-arn"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudTrailSubscriptionUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "trail_name", "trail-name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "prefix_updated"),
					resource.TestCheckResourceAttr(resourceName, "source_topic_arn", "topic-arn"),
				),
			},
		},
	})
}

func TestAccCloudTrailSubscriptionNoPrefix(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cloudtrail_subscription.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTrailSubscriptionNoPrefixTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "trail_name", "trail-name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "region"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckNoResourceAttr(resourceName, "bucket_path_prefix"),
					resource.TestCheckResourceAttr(resourceName, "source_topic_arn", "topic-arn"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudTrailSubscriptionNoPrefixUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "trail_name", "trail-name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "region_updated"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckNoResourceAttr(resourceName, "bucket_path_prefix"),
					resource.TestCheckResourceAttr(resourceName, "source_topic_arn", "topic-arn"),
				),
			},
		},
	})
}

func testAccCloudTrailSubscriptionTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
		  trail_name         = "trail-name"
		  bucket_name        = "test-costradar-bucket"
		  bucket_region      = "region"
		  bucket_path_prefix = "prefix"
          source_topic_arn   = "topic-arn"
		  access_config {
			reader_mode              = "direct"
			assume_role_arn          = "assume_role_arn_value"
			assume_role_external_id  = "assume_role_external_id_value"
			assume_role_session_name = "assume_role_session_name_value"
		  }
		}
	`
}

func testAccCloudTrailSubscriptionUpdateTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
		  trail_name         = "trail-name"
		  bucket_name        = "test-costradar-bucket"
		  bucket_region      = "region"
		  bucket_path_prefix = "prefix_updated"
          source_topic_arn   = "topic-arn"
          //include_global_service_events = true
          //is_multi_region_trail = true
          //is_organization_trail = false
		  access_config {
			reader_mode              = "direct"
			assume_role_arn          = "assume_role_arn_value"
			assume_role_external_id  = "assume_role_external_id_value"
			assume_role_session_name = "assume_role_session_name_value"
		  }
		}
	`
}

func testAccCloudTrailSubscriptionNoPrefixTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
		  trail_name         = "trail-name"
		  bucket_name        = "test-costradar-bucket"
		  bucket_region      = "region"
          source_topic_arn   = "topic-arn"
		  access_config {
			reader_mode              = "direct"
			assume_role_arn          = "assume_role_arn_value"
			assume_role_external_id  = "assume_role_external_id_value"
			assume_role_session_name = "assume_role_session_name_value"
		  }
		}
	`
}

func testAccCloudTrailSubscriptionNoPrefixUpdateTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
		  trail_name         = "trail-name"
		  bucket_name        = "test-costradar-bucket"
		  bucket_region      = "region_updated"
          source_topic_arn   = "topic-arn"
		  access_config {
			reader_mode              = "direct"
			assume_role_arn          = "assume_role_arn_value"
			assume_role_external_id  = "assume_role_external_id_value"
			assume_role_session_name = "assume_role_session_name_value"
		  }
		}
	`
}
