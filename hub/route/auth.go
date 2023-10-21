package route

import (
	"aime-api/app/telegram"
	"aime-api/config"
	"crypto/sha256"
	"time"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Data string `json:"data" binding:"required"`
	Type string `json:"type" binding:"required"`
}

func telegramOauth(cfg *config.General) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()

		var auth Auth

		if err := c.ShouldBindJSON(&auth); err != nil {
			c.JSON(400, gin.H{
				"message": "Bad Request",
				"error":   err.Error(),
			})
			return
		}

		if auth.Data == "" {
			c.JSON(400, gin.H{
				"message": "Bad Request",
			})
			return
		}

		if auth.Type == "oauth" {
			secret := sha256.Sum256([]byte(cfg.Telegram.BOT_TOKEN))
			if err := telegram.Validate(&telegram.ValidateParams{
				InitData: auth.Data,
				Token:    cfg.Telegram.BOT_TOKEN,
				ExpIn:    5 * time.Minute,
				Secret:   secret[:],
			}); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
					"error":   err.Error(),
				})
				return
			}
		} else {
			if err := telegram.Validate(&telegram.ValidateParams{
				InitData: auth.Data,
				Token:    cfg.Telegram.BOT_TOKEN,
				ExpIn:    5 * time.Minute,
			}); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
					"error":   err.Error(),
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"message": "Authorized",
		})
	}
}
