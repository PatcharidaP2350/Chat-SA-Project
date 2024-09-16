package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetMessages(c *gin.Context) {
    roomID := c.Param("room_id")

    var messages []entity.Message
    db := config.DB()

    result := db.Where("room_chat_id = ?", roomID).Find(&messages)

    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, messages)
}

func CreateMessage(c *gin.Context) {
    var request struct {
        RoomChatID uint   `json:"room_chat_id"`
        Content    string `json:"content"`
        SenderID   uint   `json:"sender_id"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    newMessage := entity.Message{
        RoomChatID: request.RoomChatID,
        Content:    request.Content,
        SenderID:   request.SenderID,
    }

    if err := db.Create(&newMessage).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Message saved successfully",
        "data":    newMessage,
    })
}
