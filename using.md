# `jsondecode` Provider

The Terraform [jsondecode](https://github.com/EvilSuperstars/terraform-provider-jsondecode) provider decodes a JSON string.

It provides a stopgap until this is supported in the Terraform configuration language; See https://github.com/hashicorp/terraform/issues/10363.

This provider requires no configuration.

### Example Usage

```hcl
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
{
  "name": "Seattle",
  "state": "WA"
}
EOS
}
```

## Data Sources

### `jsondecode_decode`

#### Argument Reference

The following arguments are supported:

* `input` - (Required, string) The JSON string that is to be decoded. The subset of JSON that can be decoded is limited - boolean, number, string, list of strings or string-to-string map.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `result_boolean` - (boolean) Boolean
* `result_list` - (list) List of strings
* `result_map` - (map) String-to-string map
* `result_number` - (float) Number
* `result_string` - (string) String
