package controller
import (
	"TTD-golang-gin-test/model/dto"
	update_shop_info "TTD-golang-gin-test/module/update_name"
	"TTD-golang-gin-test/provider/mongo_provider"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)
func UpdateNameController(c *gin.Context){
	log := logrus.New()
	requestDto:=dto.UpdateNameRequest{}
	if err := c.ShouldBindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mongodb := mongo_provider.NewMongoProvider("", log)
	repo:= update_shop_info.NewRepository(log,mongodb)
	domain:= update_shop_info.NewDomain(log, repo)

	if err := domain.Update(c.Request.Context(),requestDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}