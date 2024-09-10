// package entity

// import (
//     "gorm.io/gorm"
// )

// type Seller struct {
// 	gorm.Model
// 	StudentId       uint
// 	Year            uint
// 	UserId          uint
// 	Institute_of    string
// 	Major           string
// 	PictureStudent  string

// 	MemberTD   uint  `gorm:"unique"`
// 	// Products []Products  `gorm:"foreignKey:SellerID"`
// }

package entity

import "gorm.io/gorm"


type Seller struct {
	gorm.Model             // ใช้ gorm.Model เพื่อเพิ่มฟิลด์พื้นฐาน (ID, CreatedAt, UpdatedAt, DeletedAt)

	StudentID string
	Year int
	Institute string
	Major string
	PictureStudentID string


	MemberID       uint   `gorm:"unique"`  // One-to-One relationship กับ Member
	
}

