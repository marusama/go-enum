// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	// ListingTypeSale is a ListingType of type Sale
	ListingTypeSale ListingType = iota
	// ListingTypeRent is a ListingType of type Rent
	ListingTypeRent
)

const _ListingTypeName = "SaleRent"

var _ListingTypeNames = []string{
	_ListingTypeName[0:4],
	_ListingTypeName[4:8],
}

// ListingTypeNames returns a list of possible string values of ListingType.
func ListingTypeNames() []string {
	tmp := make([]string, len(_ListingTypeNames))
	copy(tmp, _ListingTypeNames)
	return tmp
}

// Names returns a list of possible string values of ListingType.
func (x ListingType) Names() []string {
	return ListingTypeNames()
}

var _ListingTypeMap = map[ListingType]string{
	0: _ListingTypeName[0:4],
	1: _ListingTypeName[4:8],
}

// String implements the Stringer interface.
func (x ListingType) String() string {
	if str, ok := _ListingTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ListingType(%d)", x)
}

var _ListingTypeValue = map[string]ListingType{
	_ListingTypeName[0:4]: 0,
	_ListingTypeName[4:8]: 1,
}

// ParseListingType attempts to convert a string to a ListingType
func ParseListingType(name string) (ListingType, error) {
	if x, ok := _ListingTypeValue[name]; ok {
		return x, nil
	}
	return ListingType(0), fmt.Errorf("%s is not a valid ListingType, try [%s]", name, strings.Join(_ListingTypeNames, ", "))
}

// MarshalJSON implements the json marshaller method
func (x ListingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements the json unmarshaller method
func (x *ListingType) UnmarshalJSON(b []byte) error {
	var name string
	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}
	tmp, err := ParseListingType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
