package decode

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceConfig_map = `
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
`

func TestDataSource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceConfig_map,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_list"),
					resource.TestCheckResourceAttrSet("data.jsondecode_decode.foo", "result_map"),
					resource.TestCheckNoResourceAttr("data.jsondecode_decode.foo", "result_string"),
				),
			},
		},
	})
}
