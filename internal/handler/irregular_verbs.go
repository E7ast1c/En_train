package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIrregularVerbsList(c *gin.Context) {
	iv, err := h.repos.GetAll()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNoContent,"")
		}
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, iv)
}

func (h *Handler) GetIrregularVerbById(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	iv, err := h.repos.GetIrregularVerbById(int64(intId))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNoContent,"")
		}
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, iv)
}

func (h *Handler) GetRandomVerb(c *gin.Context) {
	iv, err := h.repos.GetRandomVerb()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNoContent,"")
		}
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, iv)
}
