package irepository

import (
	 "TTD-golang-gin-test/entity"
	"context"
)


type IGetShopRepository interface {
	GetById (ctx context.Context, id int)  (*entity.Shop,error)
	GetByName (ctx context.Context, name string)  (*entity.Shop,error)
	GetByCode (ctx context.Context, code string)  (*entity.Shop,error)

}