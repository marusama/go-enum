//go:generate go-enum -f=$GOFILE --json --names

package example

// ListingType x ENUM(
// Sale
// Rent
// )
type ListingType int32
