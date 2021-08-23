package controller

import (
	"TTD-golang-gin-test/domain/update_name"
	dto2 "TTD-golang-gin-test/dto"
	"TTD-golang-gin-test/provider/mongo_provider"
	"TTD-golang-gin-test/proxy/pub_sub_proxy"
	"TTD-golang-gin-test/repository/get_shop_repo"
	"TTD-golang-gin-test/repository/update_name_shop_repo"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)
func UpdateNameController(c *gin.Context){
	log := logrus.New()
	requestDto:= dto2.UpdateNameRequest{}
	if err := c.ShouldBindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mongodb := mongo_provider.NewMongoProvider("", log)
	repo:= update_name_shop_repo.NewRepository(log,mongodb)
	getShopRepo:= get_shop_repo.NewRepository(log,mongodb)
	pubSubProxy := pub_sub_proxy.NewProxy(log, "", resty.New())

	domain:= update_shop_info.NewDomain(log, repo, getShopRepo, pubSubProxy)

	if err := domain.UpdateName(c.Request.Context(),requestDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}