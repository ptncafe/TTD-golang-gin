package update_shop_info

import (
	"TTD-golang-gin-test/common/constant"
	"TTD-golang-gin-test/interface/imodule"
	mocks "TTD-golang-gin-test/mocks/interface/imodule"
	"TTD-golang-gin-test/model/dto"
	"context"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)
//https://github.com/stretchr/testify
//https://medium.com/@thegalang/testing-in-go-mocking-mvc-using-testify-and-mockery-c25344a88691
type UpdateNameTestSuite struct {
	suite.Suite
	ctx context.Context
	log *logrus.Logger
	updateNameRepo imodule.IUpdateNameRepository
	getShopRepository imodule.IGetShopRepository
}
func TestTestUpdateName(t *testing.T) {
	suite.Run(t, new(UpdateNameTestSuite))
}


func (suite *UpdateNameTestSuite) SetupTest() {
	suite.ctx=context.Background()
	suite.log = logrus.New()
	suite.updateNameRepo =  new(mocks.IUpdateNameRepository)
	suite.getShopRepository =  new(mocks.IGetShopRepository)
}


func (suite *UpdateNameTestSuite)  TestUpdateName_Return_BadRequest_When_Id_Zero() {
	assert := assert.New(suite.T())
	request:= dto.UpdateNameRequest{Id: 0}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
	)

	err := dom.UpdateName(suite.ctx, request)
	assert.NotNil(err)
	assert.Error(err)
	if errors.IsBadRequest(err) == false{
		assert.Fail("TestUpdateName_Return_BadRequest_When_Id_Zero IsBadRequest(err) == false")
	}
	assert.Errorf(err,constant.UpdateName_Error_Message_Id)

}

func (suite *UpdateNameTestSuite)  TestUpdateName_Return_BadRequest_When_Name_Is_Short() {
	assert := assert.New(suite.T())
	request:= dto.UpdateNameRequest{Id: 1, Name: "A"}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
	)

	err := dom.UpdateName(suite.ctx, request)
	assert.NotNil(err)
	assert.Error(err)
	if errors.IsBadRequest(err) == false{
		assert.Fail("TestUpdateName_Return_BadRequest_When_Id_Zero IsBadRequest(err) == false")
	}
	assert.Errorf(err,constant.UpdateName_Error_Message_Name_Short)

}

func (suite *UpdateNameTestSuite)  TestUpdateName_Return_BadRequest_When_Name_Is_Long() {
	assert := assert.New(suite.T())
	request:= dto.UpdateNameRequest{Id: 1, Name: "TestUpdateName_Return_BadRequest_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long"}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
	)

	err := dom.UpdateName(suite.ctx, request)
	assert.NotNil(err)
	assert.Error(err)
	if errors.IsBadRequest(err) == false{
		assert.Fail("TestUpdateName_Return_BadRequest_When_Id_Zero IsBadRequest(err) == false")
	}
	assert.Errorf(err,constant.UpdateName_Error_Message_Name_Long)
}

func (suite *UpdateNameTestSuite)  TestUpdateName_Return_BadRequest_When_Not_Found_Shop() {
	assert := assert.New(suite.T())
	request:= dto.UpdateNameRequest{Id: 1, Name: "TestUpdateName_Return_BadRequest_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long_When_Name_Is_Long_TestUpdateName_Return_BadRequest_When_Name_Is_Long"}
	suite.getShopRepository.GetById(suite.ctx, request.Id)
	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
	)
	err := dom.UpdateName(suite.ctx, request)
	assert.NotNil(err)
	assert.Error(err)
	if errors.IsBadRequest(err) == false{
		assert.Fail("TestUpdateName_Return_BadRequest_When_Id_Zero IsBadRequest(err) == false")
	}
	assert.Errorf(err,constant.UpdateName_Error_Message_Name_Long)
}