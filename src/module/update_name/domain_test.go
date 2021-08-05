package update_shop_info

import (
	mocks "TTD-golang-gin-test/mocks/interface/imodule"
	"TTD-golang-gin-test/model/dto"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

///https://medium.com/@thegalang/testing-in-go-mocking-mvc-using-testify-and-mockery-c25344a88691
func TestUpdateName (t *testing.T) {
	assert := assert.New(t)
	ctx:=context.Background()
	log := logrus.New()

	request:= dto.UpdateNameRequest{}

	repo := new(mocks.IUpdateShopInfoRepository)
	dom := NewDomain(log, repo)

	err := dom.UpdateName(ctx, request)
	assert.Nil(err)
}