package example

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type listingTypeData struct {
	ListingTypeX ListingType `json:"listingType"`
}

type listingTypeData2 struct {
	ListingTypeX *ListingType `json:"listingType"`
}

func TestListingTypeMarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         *listingTypeData
		output        string
		errorExpected bool
		err           error
	}{
		{
			name:          "sale",
			output:        `{"listingType": "Sale"}`,
			input:         &listingTypeData{ListingTypeX: ListingTypeSale},
			errorExpected: false,
			err:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			raw, err := json.Marshal(test.input)
			require.NoError(tt, err, "failed marshalling to json")
			assert.JSONEq(tt, test.output, string(raw))
		})
	}
}

func TestListingTypeMarshal2(t *testing.T) {
	sale := ListingTypeSale
	tests := []struct {
		name          string
		input         *listingTypeData2
		output        string
		errorExpected bool
		err           error
	}{
		{
			name:          "nil",
			output:        `{"listingType": null}`,
			input:         &listingTypeData2{ListingTypeX: nil},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "sale",
			output:        `{"listingType": "Sale"}`,
			input:         &listingTypeData2{ListingTypeX: &sale},
			errorExpected: false,
			err:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			raw, err := json.Marshal(test.input)
			require.NoError(tt, err, "failed marshalling to json")
			assert.JSONEq(tt, test.output, string(raw))
		})
	}
}

func TestListingTypeUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		output        *listingTypeData
		errorExpected bool
		err           error
	}{
		{
			name:          "nil",
			input:         `{"listingType": null}`,
			output:        &listingTypeData{ListingTypeX: 0},
			errorExpected: true,
			err:           errors.New(" is not a valid ListingType, try [Sale, Rent]"),
		},
		{
			name:          "sale",
			input:         `{"listingType": "Sale"}`,
			output:        &listingTypeData{ListingTypeX: ListingTypeSale},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "unknown",
			input:         `{"listingType": "someunknown"}`,
			output:        &listingTypeData{ListingTypeX: 0},
			errorExpected: true,
			err:           errors.New("someunknown is not a valid ListingType, try [Sale, Rent]"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			x := &listingTypeData{}
			err := json.Unmarshal([]byte(test.input), x)
			if !test.errorExpected {
				require.NoError(tt, err, "failed unmarshalling the json.")
				assert.Equal(tt, test.output.ListingTypeX, x.ListingTypeX)
			} else {
				require.Error(tt, err)
				assert.EqualError(tt, err, test.err.Error())
			}
		})
	}
}

func TestListingTypeUnmarshal2(t *testing.T) {
	sale := ListingTypeSale
	tests := []struct {
		name          string
		input         string
		output        *listingTypeData2
		errorExpected bool
		err           error
	}{
		{
			name:          "nil",
			input:         `{"listingType": null}`,
			output:        &listingTypeData2{ListingTypeX: nil},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "sale",
			input:         `{"listingType": "Sale"}`,
			output:        &listingTypeData2{ListingTypeX: &sale},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "unknown",
			input:         `{"listingType": "someunknown"}`,
			output:        &listingTypeData2{ListingTypeX: nil},
			errorExpected: true,
			err:           errors.New("someunknown is not a valid ListingType, try [Sale, Rent]"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			x := &listingTypeData2{}
			err := json.Unmarshal([]byte(test.input), x)
			if !test.errorExpected {
				require.NoError(tt, err, "failed unmarshalling the json.")
				assert.Equal(tt, test.output.ListingTypeX, x.ListingTypeX)
			} else {
				require.Error(tt, err)
				assert.EqualError(tt, err, test.err.Error())
			}
		})
	}
}
