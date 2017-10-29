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
  "locations": [
    {"name": "Seattle", "state": "WA"},
    {"name": "New York", "state": "NY"},
    {"name": "Bellevue", "state": "WA"},
    {"name": "Olympia", "state": "WA"}
  ]
}
EOS
}
```

## Data Sources

### `jsondecode_decode`

#### Argument Reference

The following arguments are supported:

* `input` - (Required, string) The JSON string that is to be decoded.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `result_list` - (list)
* `result_map` - (map)
* `result_string` - (string)
