package decode

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceConfig_boolean = `
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
	true
EOS
}
`

const testDataSourceConfig_string = `
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
	"SSS"
EOS
}
`

const testDataSourceConfig_number = `
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
	123
EOS
}
`

const testDataSourceConfig_list = `
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
	["ABC", "zyxwv"]
EOS
}
`

const testDataSourceConfig_map = `
provider "jsondecode" {}

data "jsondecode_decode" "foo" {
  input =<<EOS
{
	"name": "Washington",
	"state": "DC"
}
EOS
}
`

func TestDataSource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceConfig_boolean,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.jsondecode_decode.foo", "result_boolean", "true"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_list"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_map"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_number"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_string"),
				),
			},
			{
				Config: testDataSourceConfig_string,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_boolean"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_list"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_map"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_number"),
					resource.TestCheckResourceAttr("data.jsondecode_decode.foo", "result_string", "SSS"),
				),
			},
			{
				Config: testDataSourceConfig_number,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_boolean"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_list"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_map"),
					resource.TestCheckResourceAttr("data.jsondecode_decode.foo", "result_number", "123"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_string"),
				),
			},
			{
				Config: testDataSourceConfig_list,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_boolean"),
					resource.TestCheckResourceAttr("data.jsondecode_decode.foo", "result_list.#", "2"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_map"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_number"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_string"),
				),
			},
			{
				Config: testDataSourceConfig_map,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_boolean"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_list"),
					resource.TestCheckResourceAttr("data.jsondecode_decode.foo", "result_map.state", "DC"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_number"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_string"),
				),
			},
		},
	})
}
