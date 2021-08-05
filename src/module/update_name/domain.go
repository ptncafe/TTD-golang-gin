package update_shop_info

import (
	"TTD-golang-gin-test/common/constant"
	"TTD-golang-gin-test/interface/imodule"
	"TTD-golang-gin-test/model/dto"
	"context"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

type domain struct {
	Log                     *logrus.Logger
	IUpdateShopInfoRepository imodule.IUpdateNameRepository
}

func NewDomain(log *logrus.Logger,
	iUpdateShopInfoRepository imodule.IUpdateNameRepository,
	) *domain {
	return &domain{
		Log: log,
		IUpdateShopInfoRepository: iUpdateShopInfoRepository,
	}
}

func(d domain) UpdateName (ctx context.Context,request dto.UpdateNameRequest)  (err error){
	if err = d.validation(ctx, request); err != nil {
		return err
	}
	return err
}

func (d domain) validation(ctx context.Context,request dto.UpdateNameRequest) error {
	if request.Id <= 0 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Id)
	}
	if len(request.Name) <= 5 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Name_Short)
	}
	if len(request.Name) > 30 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Name_Long)
	}
	return nil
}