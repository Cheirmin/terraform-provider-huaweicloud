package dms

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func getDmsRocketMQInstanceResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	region := acceptance.HW_REGION_NAME
	// getRocketmqInstance: Query DMS rocketmq instance
	var (
		getRocketmqInstanceHttpUrl = "v2/{project_id}/instances/{instance_id}"
		getRocketmqInstanceProduct = "dmsv2"
	)
	getRocketmqInstanceClient, err := cfg.NewServiceClient(getRocketmqInstanceProduct, region)
	if err != nil {
		return nil, fmt.Errorf("error creating DmsRocketMQInstance Client: %s", err)
	}

	getRocketmqInstancePath := getRocketmqInstanceClient.Endpoint + getRocketmqInstanceHttpUrl
	getRocketmqInstancePath = strings.ReplaceAll(getRocketmqInstancePath, "{project_id}", getRocketmqInstanceClient.ProjectID)
	getRocketmqInstancePath = strings.ReplaceAll(getRocketmqInstancePath, "{instance_id}", fmt.Sprintf("%v", state.Primary.ID))

	getRocketmqInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getRocketmqInstanceResp, err := getRocketmqInstanceClient.Request("GET", getRocketmqInstancePath, &getRocketmqInstanceOpt)
	if err != nil {
		return nil, fmt.Errorf("error retrieving DmsRocketMQInstance: %s", err)
	}
	return utils.FlattenResponse(getRocketmqInstanceResp)
}

func TestAccDmsRocketMQInstance_basic(t *testing.T) {
	var obj interface{}

	rName := acceptance.RandomAccResourceNameWithDash()
	updateName := acceptance.RandomAccResourceNameWithDash()
	resourceName := "huaweicloud_dms_rocketmq_instance.test"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getDmsRocketMQInstanceResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testDmsRocketMQInstance_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "engine_version", "4.8.0"),
					resource.TestCheckResourceAttr(resourceName, "enable_acl", "true"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", "0"),
				),
			},
			{
				Config: testDmsRocketMQInstance_update(updateName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "engine_version", "4.8.0"),
					resource.TestCheckResourceAttr(resourceName, "enable_acl", "false"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testDmsRocketMQInstance_basic(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_dms_rocketmq_instance" "test" {
  name              = "%s"
  engine_version    = "4.8.0"
  storage_space     = 300
  vpc_id            = huaweicloud_vpc.test.id
  subnet_id         = huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id

  availability_zones = [
    data.huaweicloud_availability_zones.test.names[0],
    data.huaweicloud_availability_zones.test.names[1],
    data.huaweicloud_availability_zones.test.names[2],
  ]

  flavor_id         = "c6.4u8g.cluster"
  storage_spec_code = "dms.physical.storage.high.v2"
  broker_num        = 1
  enable_acl        = true
}
`, common.TestBaseNetwork(name), name)
}

func testDmsRocketMQInstance_update(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_dms_rocketmq_instance" "test" {
  name              = "%s"
  engine_version    = "4.8.0"
  storage_space     = 300
  vpc_id            = huaweicloud_vpc.test.id
  subnet_id         = huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id

  availability_zones = [
    data.huaweicloud_availability_zones.test.names[2],
    data.huaweicloud_availability_zones.test.names[0],
    data.huaweicloud_availability_zones.test.names[1],
  ]

  flavor_id         = "c6.4u8g.cluster"
  storage_spec_code = "dms.physical.storage.high.v2"
  broker_num        = 1
  enable_acl        = false
}
`, common.TestBaseNetwork(name), name)
}
