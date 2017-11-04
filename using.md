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

* `input` - (Required, string) The JSON string that is to be decoded. The subset of JSON that can be decoded is limited - boolean, number, string, object with string values, array of strings and array of objects with string values.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `boolean` - (boolean) Boolean
* `number` - (float) Number
* `string` - (string) String
* `object` - (map) Object with string values
* `string_array` - (list) Array of strings
* `object_array` - (list) Array of objects with string values
