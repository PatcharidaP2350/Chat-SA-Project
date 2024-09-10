// ใน message.go
package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MemberID uint
	SellerID uint
	Content  string

	// ความสัมพันธ์ many-to-one กับ RoomChat
	RoomChatID uint
	RoomChat   RoomChat `gorm:"foreignKey:RoomChatID"`
}
