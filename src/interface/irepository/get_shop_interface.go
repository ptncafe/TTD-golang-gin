package irepository

import (
	entity2 "TTD-golang-gin-test/entity"
	"context"
)


type IGetShopRepository interface {
	GetById (ctx context.Context, id int)  (*entity2.Shop,error)
}