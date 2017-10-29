package decode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceDecode() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"input": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					s := v.(string)
					if s != "" && !json.Valid([]byte(s)) {
						errors = append(errors, fmt.Errorf("%q contains invalid JSON", k))
					}

					return
				},
			},

			"result_list": {
				Type:     schema.TypeList,
				Computed: true,
			},

			"result_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},

			"result_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	var input interface{}
	if err := json.Unmarshal([]byte(d.Get("input").(string)), &input); err != nil {
		return err
	}

	switch v := input.(type) {
	case []interface{}:
		d.Set("result_list", v)
		break
	case map[string]interface{}:
		d.Set("result_map", v)
	default:
		d.Set("result_string", v)
		break
	}

	d.SetId(time.Now().UTC().String())

	return nil
}
