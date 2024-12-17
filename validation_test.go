package gojsonschema

import (
	"fmt"
	"testing"
)

func TestValidation(t *testing.T) {
	yapiSchema := `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "type": "object",
    "properties":
    {
        "retcode":
        {
            "type": "number"
        },
        "message":
        {
            "type": "string"
        },
        "data":
        {
            "type": "object",
            "properties":
            {
                "count":
                {
                    "type": "number"
                },
                "total":
                {
                    "type": "number"
                },
                "yapi_item_list":
                {
                    "type": "array",
                    "items":
                    {
                        "type": "object",
                        "properties":
                        {
                            "joint_id":
                            {
                                "type": "number"
                            },
                            "path":
                            {
                                "type": "string"
                            }
                        },
                        "required":
                        [
                            "joint_id",
                            "path"
                        ]
                    }
                }
            },
            "required":
            [
                "yapi_item_list",
                "count"
            ]
        }
    },
    "required":
    [
        "retcode",
        "data"
    ]
}
`
	jsonData := `
{
    "retcode": 0,
    "message": "success",
    "data":
    {
		"count": 3,
        "has_wave_running": 0,
        "yapi_item_list": [
			{
				"joint_id": 1,
				"path": null
			}
		]
    }
}
`
	schemaLoader := NewStringLoader(yapiSchema)
	jsonDataLoader := NewStringLoader(jsonData)
	result, e := Validate(schemaLoader, jsonDataLoader)
	if e != nil {
		fmt.Println(e)
	}
	if result.Valid() {
		fmt.Println("The document is valid")
	} else {
		var errMsgs []string
		for _, desc := range result.Errors() {
			errMsgs = append(errMsgs, desc.String())
		}
		if len(errMsgs) == 0 {
			fmt.Println("The document is valid")
		}
		fmt.Println("The document is not valid. see errors:", errMsgs)
	}
}
