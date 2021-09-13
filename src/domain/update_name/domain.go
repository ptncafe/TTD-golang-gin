package update_shop_info

import (
	"TTD-golang-gin-test/common/constant"
	"TTD-golang-gin-test/dto"
	"TTD-golang-gin-test/interface/iproxy"
	"TTD-golang-gin-test/interface/irepository"
	"context"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

type domain struct {
	Log                       *logrus.Logger
	IUpdateShopInfoRepository irepository.IUpdateNameRepository
	IGetShop                  irepository.IGetShopRepository
	IPubSubProxy              proxy.IPubSubProxy
}

func NewDomain(log *logrus.Logger,
	updateShopInfoRepository irepository.IUpdateNameRepository,
	getShop irepository.IGetShopRepository,
	pubSubProxy  proxy.IPubSubProxy,
	) *domain {
	return &domain{
		Log: log,
		IUpdateShopInfoRepository: updateShopInfoRepository,
		IGetShop: getShop,
		IPubSubProxy: pubSubProxy,
	}
}

func(d domain) UpdateName (ctx context.Context,request dto.UpdateNameRequest)  (err error){
	if err = d.validation(ctx, request); err != nil {
		return err
	}

	shopEntity, err:= d.IGetShop.GetById(ctx,request.Id)
	if err != nil {
		return err
	}

	if shopEntity == nil {
		return errors.NotFoundf(constant.UpdateName_Error_Message_Not_Found)
	}

	dupShop, err:= d.IGetShop.GetByName(ctx,request.Name)
	if err != nil {
		return err
	}
	if dupShop != nil {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Dup_Name)
	}

	shopEntity.Code = request.Code
	shopEntity.Name = request.Name
	shopEntity.UpdatedUser = request.UpdatedUser
	shopEntity.UpdatedDate = request.UpdatedDate


	err = d.IUpdateShopInfoRepository.UpdateName(ctx,*shopEntity)
	if err != nil {
		return err
	}

	err = d.IPubSubProxy.PubShop(*shopEntity)
	if err != nil {
		return err
	}
	return nil
}

func (d domain) validation(ctx context.Context,request dto.UpdateNameRequest) error {
	if request.Id <= 0 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Id)
	}
	if len(request.Name) <= 5 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Name_Short)
	}
	if len(request.Name) > 50 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_Name_Long)
	}
	if len(request.UpdatedUser) <= 0 {
		return errors.BadRequestf(constant.UpdateName_Error_Message_UpdatedUser)
	}
	return nil
}