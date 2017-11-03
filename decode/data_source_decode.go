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

			"result_boolean": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"result_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"result_map": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"result_number": {
				Type:     schema.TypeFloat,
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
		l, err := stringList(v)
		if err != nil {
			return err
		}
		d.Set("result_list", l)
	case map[string]interface{}:
		m, err := stringToStringMap(v)
		if err != nil {
			return err
		}
		d.Set("result_map", m)
	case string:
		d.Set("result_string", v)
	case float64:
		d.Set("result_number", v)
	case bool:
		d.Set("result_boolean", v)
	default:
		return fmt.Errorf("unsupported top-level type: %T", v)
	}

	d.SetId(time.Now().UTC().String())

	return nil
}

func stringList(v []interface{}) ([]string, error) {
	l := make([]string, 0)
	for _, value := range v {
		if s, ok := value.(string); ok {
			l = append(l, s)
		} else {
			return nil, fmt.Errorf("unsupported array element type: %T", value)
		}
	}

	return l, nil
}

func stringToStringMap(v map[string]interface{}) (map[string]string, error) {
	m := make(map[string]string)
	for key, value := range v {
		if s, ok := value.(string); ok {
			m[key] = s
		} else {
			return nil, fmt.Errorf("unsupported map value type: %T", value)
		}
	}

	return m, nil
}
