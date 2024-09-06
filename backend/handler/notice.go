package handler

import (
	"net/http"
	"student/model"
	service "student/service"

	"github.com/gin-gonic/gin"
)

func NoticeCreate(c *gin.Context) {
	// Validate input
	var notice model.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	noticeService := service.NoticeRepository{}
	response, err := noticeService.NoticeCreate(&notice)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func NoticeList(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	noticeService := service.NoticeRepository{}
	response, err := noticeService.NoticeList(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteNotices(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	noticeService := service.NoticeRepository{}
	response, err := noticeService.DeleteNotices(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteNotice(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	noticeService := service.NoticeRepository{}
	response, err := noticeService.DeleteNotice(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func UpdateNotice(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	// Validate input
	var notice model.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	noticeService := service.NoticeRepository{}
	response, err := noticeService.UpdateNotice(&notice, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
