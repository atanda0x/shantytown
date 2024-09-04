package handlers

import (
	"net/http"

	"github.com/0xatanda/shantytown/templates"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	templates.Home().Render(c.Writer)
}

func CommunityMembers(c *gin.Context) {
	members := []string{"member 1", "Member 2", "Member 3"}
	c.JSON(http.StatusOK, members)
}

func Join(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Joined successfully"})
}
