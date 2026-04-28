package handlers

import (
    "api-gateway/clients"
    "api-gateway/utils"

    "github.com/gin-gonic/gin"
)

type AuthRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Register(c *gin.Context) {
    var req AuthRequest
    c.ShouldBindJSON(&req)
    res, err := clients.Register(req.Username, req.Password)
    if err != nil {
        utils.Error(c, err.Error())
        return
    }
    utils.Success(c, res)
}

func Login(c *gin.Context) {
    var req AuthRequest
    c.ShouldBindJSON(&req)
    res, err := clients.Login(req.Username, req.Password)
    if err != nil {
        utils.Error(c, err.Error())
        return
    }
    utils.Success(c, res)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    res, err := clients.GetUser(id)
    if err != nil {
        utils.Error(c, err.Error())
        return
    }
    utils.Success(c, res)
}