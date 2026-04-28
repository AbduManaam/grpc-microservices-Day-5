package handlers

import (
    "api-gateway/clients"
    "api-gateway/utils"

    "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    userID := c.Query("user_id")
    res, err := clients.CreateOrder(userID)
    if err != nil {
        utils.Error(c, err.Error())
        return
    }
    utils.Success(c, res)
}