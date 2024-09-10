package entity


type Member struct {

	MemberID    uint `gorm:"primaryKey"` // ID ต้องมีเพื่อเชื่อมโยงกับ RoomChat
	Username    string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	PicProfile  string
	Seller      Seller `gorm:"foreignKey:MemberID"`
	// RoomChat    RoomChat `gorm:"foreignKey:MemberID"`

	RoomChats   []RoomChat `gorm:"foreignKey:MemberID"`
}
