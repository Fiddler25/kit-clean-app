// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kit-clean-app/ent/order"
	"kit-clean-app/ent/product"
	"kit-clean-app/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreated is the schema descriptor for created field.
	orderDescCreated := orderMixinFields0[0].Descriptor()
	// order.DefaultCreated holds the default value on creation for the created field.
	order.DefaultCreated = orderDescCreated.Default.(func() time.Time)
	// orderDescUpdated is the schema descriptor for updated field.
	orderDescUpdated := orderMixinFields0[1].Descriptor()
	// order.DefaultUpdated holds the default value on creation for the updated field.
	order.DefaultUpdated = orderDescUpdated.Default.(func() time.Time)
	// order.UpdateDefaultUpdated holds the default value on update for the updated field.
	order.UpdateDefaultUpdated = orderDescUpdated.UpdateDefault.(func() time.Time)
	// orderDescTotalPrice is the schema descriptor for total_price field.
	orderDescTotalPrice := orderFields[4].Descriptor()
	// order.TotalPriceValidator is a validator for the "total_price" field. It is called by the builders before save.
	order.TotalPriceValidator = orderDescTotalPrice.Validators[0].(func(float64) error)
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreated is the schema descriptor for created field.
	productDescCreated := productMixinFields0[0].Descriptor()
	// product.DefaultCreated holds the default value on creation for the created field.
	product.DefaultCreated = productDescCreated.Default.(func() time.Time)
	// productDescUpdated is the schema descriptor for updated field.
	productDescUpdated := productMixinFields0[1].Descriptor()
	// product.DefaultUpdated holds the default value on creation for the updated field.
	product.DefaultUpdated = productDescUpdated.Default.(func() time.Time)
	// product.UpdateDefaultUpdated holds the default value on update for the updated field.
	product.UpdateDefaultUpdated = productDescUpdated.UpdateDefault.(func() time.Time)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[1].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescDescription is the schema descriptor for description field.
	productDescDescription := productFields[2].Descriptor()
	// product.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	product.DescriptionValidator = productDescDescription.Validators[0].(func(string) error)
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[3].Descriptor()
	// product.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	product.PriceValidator = productDescPrice.Validators[0].(func(float64) error)
	// productDescStock is the schema descriptor for stock field.
	productDescStock := productFields[4].Descriptor()
	// product.StockValidator is a validator for the "stock" field. It is called by the builders before save.
	product.StockValidator = productDescStock.Validators[0].(func(uint8) error)
}
