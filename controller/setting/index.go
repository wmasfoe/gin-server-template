package setting

import "github.com/gin-gonic/gin"

type SettingsController struct {
}

func (s SettingsController) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}
