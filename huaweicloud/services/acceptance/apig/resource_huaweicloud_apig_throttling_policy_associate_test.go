package apig

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/apigw/dedicated/v2/throttles"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
)

func getAssociateFunc(conf *config.Config, state *terraform.ResourceState) (interface{}, error) {
	c, err := conf.ApigV2Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating APIG v2 client: %s", err)
	}
	opt := throttles.ListBindOpts{
		InstanceId: state.Primary.Attributes["instance_id"],
		ThrottleId: state.Primary.Attributes["policy_id"],
	}
	return throttles.ListBind(c, opt)
}

func TestAccThrottlingPolicyAssociate_basic(t *testing.T) {
	var (
		apiDetails []throttles.ApiForThrottle

		// The dedicated instance name only allow letters, digits and underscores (_).
		name  = acceptance.RandomAccResourceName()
		rName = "huaweicloud_apig_throttling_policy_associate.test"
	)

	rc := acceptance.InitResourceCheck(
		rName,
		&apiDetails,
		getAssociateFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckEpsID(t) // The creation of APIG instance needs the enterprise project ID.
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccThrottlingPolicyAssociate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName, "instance_id",
						"huaweicloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName, "policy_id",
						"huaweicloud_apig_throttling_policy.test", "id"),
					resource.TestCheckResourceAttr(rName, "publish_ids.#", "1"),
				),
			}, {
				Config: testAccThrottlingPolicyAssociate_update(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName, "instance_id",
						"huaweicloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName, "policy_id",
						"huaweicloud_apig_throttling_policy.test", "id"),
					resource.TestCheckResourceAttr(rName, "publish_ids.#", "1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccThrottlingPolicyAssociate_base(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_apig_instance" "test" {
  name                  = "%[2]s"
  edition               = "BASIC"
  vpc_id                = huaweicloud_vpc.test.id
  subnet_id             = huaweicloud_vpc_subnet.test.id
  security_group_id     = huaweicloud_networking_secgroup.test.id
  enterprise_project_id = "0"

  availability_zones = try(slice(data.huaweicloud_availability_zones.test.names, 0, 1), null)
}

resource "huaweicloud_compute_instance" "test" {
  name               = "%[2]s"
  image_id           = data.huaweicloud_images_images.test.images[0].id
  flavor_id          = data.huaweicloud_compute_flavors.test.ids[0]
  security_group_ids = [huaweicloud_networking_secgroup.test.id]
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  system_disk_type   = "SSD"

  network {
    uuid = huaweicloud_vpc_subnet.test.id
  }
}

resource "huaweicloud_apig_group" "test" {
  name        = "%[2]s"
  instance_id = huaweicloud_apig_instance.test.id
}

resource "huaweicloud_apig_vpc_channel" "test" {
  name        = "%[2]s"
  instance_id = huaweicloud_apig_instance.test.id
  port        = 80
  algorithm   = "WRR"
  protocol    = "HTTP"
  path        = "/"
  http_code   = "201"

  members {
    id = huaweicloud_compute_instance.test.id
  }
}

resource "huaweicloud_apig_api" "test" {
  instance_id             = huaweicloud_apig_instance.test.id
  group_id                = huaweicloud_apig_group.test.id
  name                    = "%[2]s"
  type                    = "Public"
  request_protocol        = "HTTP"
  request_method          = "GET"
  request_path            = "/user_info/{user_age}"
  security_authentication = "APP"
  matching                = "Exact"
  success_response        = "Success response"
  failure_response        = "Failed response"
  description             = "Created by script"

  request_params {
    name     = "user_age"
    type     = "NUMBER"
    location = "PATH"
    required = true
    maximum  = 200
    minimum  = 0
  }
  
  backend_params {
    type     = "REQUEST"
    name     = "userAge"
    location = "PATH"
    value    = "user_age"
  }

  web {
    path             = "/getUserAge/{userAge}"
    vpc_channel_id   = huaweicloud_apig_vpc_channel.test.id
    request_method   = "GET"
    request_protocol = "HTTP"
    timeout          = 30000
  }

  web_policy {
    name             = "%[2]s_policy1"
    request_protocol = "HTTP"
    request_method   = "GET"
    effective_mode   = "ANY"
    path             = "/getUserAge/{userAge}"
    timeout          = 30000
    vpc_channel_id   = huaweicloud_apig_vpc_channel.test.id

    backend_params {
      type     = "REQUEST"
      name     = "userAge"
      location = "PATH"
      value    = "user_age"
    }

    conditions {
      source     = "param"
      param_name = "user_age"
      type       = "Equal"
      value      = "28"
    }
  }
}

resource "huaweicloud_apig_environment" "test" {
  count = 2

  name        = "%[2]s_${count.index}"
  instance_id = huaweicloud_apig_instance.test.id
}

resource "huaweicloud_apig_api_publishment" "test" {
  count = 2

  instance_id = huaweicloud_apig_instance.test.id
  api_id      = huaweicloud_apig_api.test.id
  env_id      = huaweicloud_apig_environment.test[count.index].id
}

resource "huaweicloud_apig_throttling_policy" "test" {
  instance_id       = huaweicloud_apig_instance.test.id
  name              = "%[2]s"
  type              = "API-based"
  period            = 15000
  period_unit       = "SECOND"
  max_api_requests  = 100
  max_user_requests = 60
  max_app_requests  = 60
  max_ip_requests   = 60
}
`, common.TestBaseComputeResources(name), name)
}

func testAccThrottlingPolicyAssociate_basic(name string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_apig_throttling_policy_associate" "test" {
  instance_id = huaweicloud_apig_instance.test.id
  policy_id   = huaweicloud_apig_throttling_policy.test.id

  publish_ids = [
    huaweicloud_apig_api_publishment.test[0].publish_id
  ]
}
`, testAccThrottlingPolicyAssociate_base(name))
}

func testAccThrottlingPolicyAssociate_update(name string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_apig_throttling_policy_associate" "test" {
  instance_id = huaweicloud_apig_instance.test.id
  policy_id   = huaweicloud_apig_throttling_policy.test.id

  publish_ids = [
    huaweicloud_apig_api_publishment.test[1].publish_id
  ]
}
`, testAccThrottlingPolicyAssociate_base(name))
}
