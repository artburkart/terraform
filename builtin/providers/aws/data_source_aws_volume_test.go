package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSVolumeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckAwsAmiDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsVolumeDataSourceID("data.aws_volume.ebs_volume"),
					resource.TestCheckResourceAttr("data.aws_ami.nat_ami", "virtualization_type", "hvm"),
				),
			},
		},
	})
}

func testAccCheckAwsVolumeDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckAwsVolumeDataSourceID(n string) resource.TestCheckFunc {
	// Wait for IAM role
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Volume data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Volume data source ID not set")
		}
		return nil
	}
}

const testAccCheckAwsAmiDataSourceConfig = `
data "aws_ami" "nat_ami" {
	most_recent = true
	filter {
		name = "owner-alias"
		values = ["amazon"]
	}
	filter {
		name = "name"
		values = ["amzn-ami-vpc-nat*"]
	}
	filter {
		name = "virtualization-type"
		values = ["hvm"]
	}
	filter {
		name = "root-device-type"
		values = ["ebs"]
	}
	filter {
		name = "block-device-mapping.volume-type"
		values = ["standard"]
	}
}
`

// Windows image test.
const testAccCheckAwsAmiDataSourceWindowsConfig = `
data "aws_ami" "windows_ami" {
	most_recent = true
	filter {
		name = "owner-alias"
		values = ["amazon"]
	}
	filter {
		name = "name"
		values = ["Windows_Server-2012-R2*"]
	}
	filter {
		name = "virtualization-type"
		values = ["hvm"]
	}
	filter {
		name = "root-device-type"
		values = ["ebs"]
	}
	filter {
		name = "block-device-mapping.volume-type"
		values = ["gp2"]
	}
}
`

// Instance store test - using Ubuntu images
const testAccCheckAwsAmiDataSourceInstanceStoreConfig = `
data "aws_ami" "instance_store_ami" {
	most_recent = true
	filter {
		name = "owner-id"
		values = ["099720109477"]
	}
	filter {
		name = "name"
		values = ["ubuntu/images/hvm-instance/ubuntu-trusty-14.04-amd64-server*"]
	}
	filter {
		name = "virtualization-type"
		values = ["hvm"]
	}
	filter {
		name = "root-device-type"
		values = ["instance-store"]
	}
}
`

// Testing owner parameter
const testAccCheckAwsAmiDataSourceOwnersConfig = `
data "aws_ami" "amazon_ami" {
	most_recent = true
	owners = ["amazon"]
}
`

// Testing name_regex parameter
const testAccCheckAwsAmiDataSourceNameRegexConfig = `
data "aws_ami" "name_regex_filtered_ami" {
	most_recent = true
	owners = ["amazon"]
	filter {
		name = "name"
		values = ["amzn-ami-*"]
	}
	name_regex = "^amzn-ami-\\d{3}[5].*-ecs-optimized"
}
`
