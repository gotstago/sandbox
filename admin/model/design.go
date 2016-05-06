package model

import "github.com/jinzhu/gorm"

type Entity struct {
	gorm.Model
	DisplayName   string
	Description   string
	TableName     string
	NonReportable bool
	//ReportingFolder bool

	// UserID            uint
	// User              User
	// PaymentAmount     float32
	// AbandonedReason   string
	// DiscountValue     uint
	// TrackingNumber    *string
	// ShippedAt         *time.Time
	// CancelledAt       *time.Time
	// ShippingAddressID uint
	// ShippingAddress   Address
	// BillingAddressID  uint
	// BillingAddress    Address
	EntityAttributes []EntityAttribute
	// transition.Transition
	Resource bool
	Enabled  bool
}

type EntityAttribute struct {
	gorm.Model
	EntityID           uint
	Name               string
	DisplayName        string
	AttributeTypeID    uint
	AttributeType      AttributeType
	ObjectDisplayValue bool
	Deprecated         bool
	Private            bool

	// SizeVariationID uint
	// SizeVariation   SizeVariation
	// Quantity        uint
	// Price           float32
	// DiscountRate    uint
	// transition.Transition
}

type DataType struct {
	gorm.Model
	//Code        string
	DisplayName string
}

type AttributeType struct {
	gorm.Model
	Code        string
	DisplayName string
}

// func (entity Entity) Amount() (amount float32) {
// 	for _, entityItem := range entity.EntityAttributes {
// 		amount += entityItem.Amount()
// 	}
// 	return
// }

// func (item EntityAttribute) Amount() float32 {
// 	return item.Price * float32(item.Quantity) * float32(100-item.DiscountRate) / 100
// }

// var (
// 	EntityState = transition.New(&Entity{})
// 	ItemState  = transition.New(&EntityAttribute{})
// )

func init() {
	// Define Entity's States
	// EntityState.Initial("draft")
	// EntityState.State("checkout")
	// EntityState.State("cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	tx.Model(value).UpdateColumn("cancelled_at", time.Now())
	// 	return nil
	// })
	// EntityState.State("paid").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	// freeze stock, change items's state
	// 	return nil
	// })
	// EntityState.State("paid_cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	// do refund, release stock, change items's state
	// 	return nil
	// })
	// EntityState.State("processing")
	// EntityState.State("shipped")
	// EntityState.State("returned")

	// EntityState.Event("checkout").To("checkout").From("draft")
	// EntityState.Event("pay").To("paid").From("checkout")
	// cancelEvent := EntityState.Event("cancel")
	// cancelEvent.To("cancelled").From("draft", "checkout")
	// cancelEvent.To("paid_cacelled").From("paid", "processing", "shipped")
	// EntityState.Event("process").To("processing").From("paid")
	// EntityState.Event("ship").To("shipped").From("processing")
	// EntityState.Event("return").To("returned").From("shipped")

	// // Define ItemItem's States
	// ItemState.Initial("checkout")
	// ItemState.State("cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	// release stock, upate entity state
	// 	return nil
	// })
	// ItemState.State("paid").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	// freeze stock, update entity state
	// 	return nil
	// })
	// ItemState.State("paid_cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
	// 	// do refund, release stock, update entity state
	// 	return nil
	// })
	// ItemState.State("processing")
	// ItemState.State("shipped")
	// ItemState.State("returned")

	// ItemState.Event("checkout").To("checkout").From("draft")
	// ItemState.Event("pay").To("paid").From("checkout")
	// cancelItemEvent := ItemState.Event("cancel")
	// cancelItemEvent.To("cancelled").From("checkout")
	// cancelItemEvent.To("paid_cacelled").From("paid")
	// ItemState.Event("process").To("processing").From("paid")
	// ItemState.Event("ship").To("shipped").From("processing")
	// ItemState.Event("return").To("returned").From("shipped")
}
