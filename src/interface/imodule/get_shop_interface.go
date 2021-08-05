package imodule

import (
	"TTD-golang-gin-test/model/entity"
	"context"
)


type IGetShopRepository interface {
	GetById (ctx context.Context, id int)  (*entity.Shop,error)
}