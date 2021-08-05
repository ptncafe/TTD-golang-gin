package imodule

import (
	"TTD-golang-gin-test/model/dto"
	"context"
)

type IUpdateShopInfoDomain interface {
	UpdateName (ctx context.Context,request dto.UpdateNameRequest)  error
}

type IUpdateShopInfoRepository interface {
	UpdateName (ctx context.Context,request dto.UpdateNameRequest)  error
}