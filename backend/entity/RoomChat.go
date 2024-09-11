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
