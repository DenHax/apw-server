package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUpload(c *gin.Context) {
	uploads, err := h.Services.Upload.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uploads)
}
func (h *Handler) createUpload(c *gin.Context) {

}
func (h *Handler) deleteUpload(c *gin.Context) {

}
