package handlers

import (
	"context"

	"github.com/0xatanda/shantytown/templates"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	ctx := context.Background()
	err := templates.Home().Render(ctx, c.Writer)
	if err != nil {
		c.String(500, "Internal Server Error")
	}
}

func CommunityMembers(c *gin.Context) {
	members := []string{
		"https://picsum.photos/100?random=1",
		"https://picsum.photos/100?random=2",
		"https://picsum.photos/100?random=3",
		// Add more random image URLs as needed
	}

	html := "<div class='profile-grid'>"
	for _, member := range members {
		html += "<div class='profile-circle'><img src='" + member + "' alt='Community Member'></div>"
	}
	html += "</div>"

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func Join(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Joined successfully"})
}
