// +build unit,!integration

package request

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateBatCaveSettingRequest_Validate_Should_Return_Error_When_ID_Field_Is_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveDeviceSettingRequest
		id             string
		expectedErrors map[string]interface{}
	}{
		{
			name: "ID_Length_Is_Greater_Then_36_Characters_Long",
			req: UpdateBatCaveDeviceSettingRequest{
				DeepSleepDelay: 10,
			},
			id: "00000000-1111-2222-3333-5555555555555",
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
			req: UpdateBatCaveDeviceSettingRequest{
				DeepSleepDelay: 10,
			},
			id: "00000000-1111-2222-3333-55555555555",
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
			req: UpdateBatCaveDeviceSettingRequest{
				DeepSleepDelay: 15,
			},
			id: "00000000-1111-2222-3333-55555555555F",
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
			req: UpdateBatCaveDeviceSettingRequest{
				DeepSleepDelay: 20,
			},
			id: "",
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

func Test_UpdateBatCaveSettingRequest_Validate_Should_Return_Error_When_DeepSleepDelayIs_Not_Valid(t *testing.T) {
	testCases := []struct {
		name           string
		req            UpdateBatCaveDeviceSettingRequest
		id             string
		expectedErrors map[string]interface{}
	}{
		{
			name: "Deep_Sleep_Delay_Equals 0",
			req: UpdateBatCaveDeviceSettingRequest{
				DeepSleepDelay: 0,
			},
			id: "00000000-1111-2222-3333-555555555555",
			expectedErrors: map[string]interface{}{
				"validationError": url.Values{
					"deepSleepDelay": []string{
						"The deepSleepDelay field should be a positive non-zero value!",
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
