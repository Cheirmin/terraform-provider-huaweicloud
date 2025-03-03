---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud_dms_rocketmq_broker

Use this data source to get the list of DMS rocketMQ broker.

## Example Usage

```hcl
var "instance_id" {}

data "huaweicloud_dms_rocketmq_broker" "test" {
  instance_id = var.instance_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the data source.
  If omitted, the provider-level region will be used.

* `instance_id` - (Optional, String) Specifies the ID of the rocketMQ instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `brokers` - Indicates the list of the brokers.
