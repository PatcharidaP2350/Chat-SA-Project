// package entity

// type RoomChat struct {

// 	RoomID   int `gorm:"unique"` // RoomID เป็น string และต้องไม่ซ้ำกันในแต่ละห้อง

// 	MemberID uint
// 	Member   Member `gorm:"foreignKey:MemberID"`

// 	SellerID uint
// 	Seller   Seller `gorm:"foreignKey:SellerID"`  // One-to-One relationship กับ Seller

// Messages []Message `gorm:"foreignKey:RoomChatID"`
// }

/*----------------------------------------------------*/

// package entity

// type RoomChat struct {
// 	RoomID   uint `gorm:"unique"` // RoomID ควรเป็น unique

// 	MemberID uint
// 	Member   Member `gorm:"foreignKey:MemberID"`

// 	SellerID uint
// 	Seller   Seller `gorm:"foreignKey:SellerID"`

// 	// ความสัมพันธ์ one-to-many กับ Message
// 	Messages []Message `gorm:"foreignKey:RoomChatID"`
// }

/*----------------------------------------------------*/

package entity


type RoomChat struct {

	RoomID uint `gorm:"primaryKey;autoIncrement"` // ใช้ uint เป็น Primary Key

	MemberID uint
	Member   Member `gorm:"foreignKey:MemberID"`

	SellerID uint
	Seller   Seller `gorm:"foreignKey:SellerID"`

	// เชื่อมต่อ one-to-many กับ Message
	Messages []Message `gorm:"foreignKey:RoomChatID"`
}
