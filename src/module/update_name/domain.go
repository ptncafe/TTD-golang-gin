package update_shop_info

import (
	"TTD-golang-gin-test/interface/imodule"
	"TTD-golang-gin-test/model/dto"
	"context"
	"github.com/sirupsen/logrus"
)

type domain struct {
	Log                     *logrus.Logger
	IUpdateShopInfoRepository imodule.IUpdateShopInfoRepository
}

func NewDomain(log *logrus.Logger,
	iUpdateShopInfoRepository imodule.IUpdateShopInfoRepository,
	) *domain {
	return &domain{
		Log: log,
		IUpdateShopInfoRepository: iUpdateShopInfoRepository,
	}
}

func(d domain) UpdateName (ctx context.Context,request dto.UpdateNameRequest)  error{
	return nil
}