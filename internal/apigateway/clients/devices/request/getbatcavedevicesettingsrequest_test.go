// +build unit,!integration

package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBatCaveSettingRequest_Should_Return_Error_When_ID_Field_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            GetBatCaveDeviceSettingRequest
		id             string
		expectedErrors map[string]interface{}
	}{
		{
			name: "ID_Length_Is_Greater_Then_36_Characters_Long",
			req:  GetBatCaveDeviceSettingRequest{},
			id:   "00000000-1111-2222-3333-5555555555555",
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"id": []string{
						"The id field needs to be a valid!",
					},
				},
			},
		},
		{
			name: "ID_Length_Is_Less_Then_36_Characters_Long",
			req:  GetBatCaveDeviceSettingRequest{},
			id:   "00000000-1111-2222-3333-55555555555",
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"id": []string{
						"The id field needs to be a valid!",
					},
				},
			},
		},
		{
			name: "ID_Contains_An_Invalid_ID_Charater",
			id:   "00000000-1111-2222-3333-55555555555F",
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"id": []string{
						"The id field needs to be a valid!",
					},
				},
			},
		},
		{
			name: "ID_Is_Empty",
			req:  GetBatCaveDeviceSettingRequest{},
			id:   "",
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"id": []string{
						"The id field needs to be a valid!",
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validationErrors := testCase.req.Validate(testCase.id)
			errs := map[string]interface{}{"validationError": validationErrors}
			errorMessage := fmt.Sprintf("should have errors: %s", testCase.expectedErrors)
			assert.Equal(t, testCase.expectedErrors, errs, errorMessage)
		})
	}
}
