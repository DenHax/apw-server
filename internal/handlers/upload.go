package handler

import (
	"apw/internal/domain/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getAllUpload(c *gin.Context) {
	uploads, err := h.Services.Upload.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uploads)
}

func (h *Handler) getReport(c *gin.Context) {
	fdateStr := c.Query("fdate")
	sdateStr := c.Query("sdate")
	layout := "2006-01-02"

	fdate, err := time.Parse(layout, fdateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	sdate, err := time.Parse(layout, sdateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	uploads, err := h.Services.Upload.GetReport(fdate, sdate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uploads)
}

func (h *Handler) createUpload(c *gin.Context) {
	var input models.Upload
	var err error

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("create upload")
	fmt.Println(input)

	// layout := "2006-01-02 15:04:05.000"
	// now := time.Now()
	// input.LoadDate, err = time.Parse(layout, now.Format(layout))
	id, err := h.Services.Upload.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Errorf("Error to create upload: %s", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) deleteUpload(c *gin.Context) {
	idStr := c.Param("id")
	id, err := time.Parse(time.RFC3339, idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Errorf("Error convert load date: %s", err)
		return
	}
	err = h.Services.Upload.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Errorf("Error to delete load date: %s", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена", "id": id})

}
