// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger Potion maker - OpenAPI 3.0
 *
 * Сервис изготовления зелий
 *
 * API version: 1.0.0
 */

package openapi

type Witch struct {
	Uuid string `json:"uuid"`

	Name string `json:"name"`
}

// AssertWitchRequired checks if the required fields are not zero-ed
func AssertWitchRequired(obj Witch) error {
	elements := map[string]interface{}{
		"uuid": obj.Uuid,
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertWitchConstraints checks if the values respects the defined constraints
func AssertWitchConstraints(obj Witch) error {
	return nil
}
