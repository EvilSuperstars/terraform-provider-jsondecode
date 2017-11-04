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

			"boolean": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"number": {
				Type:     schema.TypeFloat,
				Computed: true,
			},

			"string": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"string_array": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"object_array": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
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
		if len(v) == 0 {
			d.Set("string_array", make([]string, 0))
			break
		}

		switch v0 := v[0].(type) {
		case string:
			a, err := stringArray(v)
			if err != nil {
				return err
			}
			d.Set("string_array", a)
		case map[string]interface{}:
			a, err := objectArray(v)
			if err != nil {
				return err
			}
			d.Set("object_array", a)
		default:
			return fmt.Errorf("unsupported array element type: %T", v0)
		}
	case map[string]interface{}:
		o, err := object(v)
		if err != nil {
			return err
		}
		d.Set("object", o)
	case string:
		d.Set("string", v)
	case float64:
		d.Set("number", v)
	case bool:
		d.Set("boolean", v)
	default:
		return fmt.Errorf("unsupported top-level type: %T", v)
	}

	d.SetId(time.Now().UTC().String())

	return nil
}

func stringArray(v []interface{}) ([]string, error) {
	a := make([]string, 0)
	for _, value := range v {
		if s, ok := value.(string); ok {
			a = append(a, s)
		} else {
			return nil, fmt.Errorf("unsupported array element type in string array: %T", value)
		}
	}

	return a, nil
}

func objectArray(v []interface{}) ([]map[string]string, error) {
	a := make([]map[string]string, 0)
	for _, value := range v {
		if m, ok := value.(map[string]interface{}); ok {
			o, err := object(m)
			if err != nil {
				return nil, err
			}
			a = append(a, o)
		} else {
			return nil, fmt.Errorf("unsupported array element type in object array: %T", value)
		}
	}

	return a, nil
}

func object(v map[string]interface{}) (map[string]string, error) {
	o := make(map[string]string)
	for key, value := range v {
		if s, ok := value.(string); ok {
			o[key] = s
		} else {
			return nil, fmt.Errorf("unsupported object value type: %T", value)
		}
	}

	return o, nil
}
