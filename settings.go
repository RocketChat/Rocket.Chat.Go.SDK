package goRocket

import (
	"log"

	"github.com/Jeffail/gabs"
)

// GetPublicSettings gets public settings
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-public-settings
func (c *LiveService) GetPublicSettings() ([]Setting, error) {
	rawResponse, err := c.client.ddp.Call("public-settings/get")
	if err != nil {
		return nil, err
	}

	document, _ := gabs.Consume(rawResponse)

	sett, _ := document.Children()

	var settings []Setting

	for _, rawSetting := range sett {
		setting := Setting{
			ID:   stringOrZero(rawSetting.Path("_id").Data()),
			Type: stringOrZero(rawSetting.Path("type").Data()),
		}

		switch setting.Type {
		case "boolean":
			setting.ValueBool = rawSetting.Path("value").Data().(bool)
		case "string":
			setting.Value = stringOrZero(rawSetting.Path("value").Data())
		case "code":
			setting.Value = stringOrZero(rawSetting.Path("value").Data())
		case "color":
			setting.Value = stringOrZero(rawSetting.Path("value").Data())
		case "int":
			setting.ValueInt = rawSetting.Path("value").Data().(float64)
		case "asset":
			setting.ValueAsset = Asset{
				DefaultURL: stringOrZero(rawSetting.Path("value").Data().(map[string]interface{})["defaultUrl"]),
			}

		default:
			log.Println(setting.Type, rawSetting.Path("value").Data())
		}

		settings = append(settings, setting)
	}

	return settings, nil
}
