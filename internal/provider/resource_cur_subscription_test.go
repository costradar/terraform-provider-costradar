package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccCurSubscription(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cur_subscription.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCurSubscriptionTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "bucket_region"),
					resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "xxx!!"),
					resource.TestCheckResourceAttr(resourceName, "time_unit", "hour"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
				),
			},
		},
	})
}

func TestAccCurSubscriptionEmptyPrefix(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cur_subscription.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCurSubscriptionEmptyPrefixTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "bucket_region"),
					//resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "xxx!!"),
					resource.TestCheckNoResourceAttr(resourceName, "bucket_path_prefix"),
					resource.TestCheckResourceAttr(resourceName, "time_unit", "hour"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCurSubscriptionEmptyPrefixUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "test-costradar-bucket"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "bucket_region_update"),
					resource.TestCheckNoResourceAttr(resourceName, "bucket_path_prefix"),
					//resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "xxx!!"),
					resource.TestCheckResourceAttr(resourceName, "time_unit", "hour"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
				),
			},
		},
	})
}

func testAccCurSubscriptionTF() string {
	return `
	  resource "costradar_cur_subscription" "test" {
	  	report_name        = "report_name"
	  	bucket_name        = "test-costradar-bucket"
	  	bucket_region      = "bucket_region"
	  	time_unit          = "hour"
		source_topic_arn   = "topic-arn"
		bucket_path_prefix = "xxx!!"
	  	access_config {
		  reader_mode              = "direct"
 		  assume_role_arn          = "assume_role_arn_value"
		  assume_role_external_id  = "assume_role_external_id_value"
		  assume_role_session_name = "assume_role_session_name_value"
	    }
	  }`
}

func testAccCurSubscriptionEmptyPrefixTF() string {
	return `
	  resource "costradar_cur_subscription" "test" {
	  	report_name        = "report_name"
	  	bucket_name        = "test-costradar-bucket"
	  	bucket_region      = "bucket_region"
	  	time_unit          = "hour"
		source_topic_arn   = "topic-arn"
	  	access_config {
		  reader_mode              = "direct"
 		  assume_role_arn          = "assume_role_arn_value"
		  assume_role_external_id  = "assume_role_external_id_value"
		  assume_role_session_name = "assume_role_session_name_value"
	    }
	  }`
}

func testAccCurSubscriptionEmptyPrefixUpdateTF() string {
	return `
	  resource "costradar_cur_subscription" "test" {
	  	report_name        = "report_name"
	  	bucket_name        = "test-costradar-bucket"
	  	bucket_region      = "bucket_region_update"
	  	time_unit          = "hour"
		source_topic_arn   = "topic-arn"
	  	access_config {
		  reader_mode              = "direct"
 		  assume_role_arn          = "assume_role_arn_value"
		  assume_role_external_id  = "assume_role_external_id_value"
		  assume_role_session_name = "assume_role_session_name_value"
	    }
	  }`
}
