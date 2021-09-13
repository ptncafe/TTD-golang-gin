package irepository

import (
	"TTD-golang-gin-test/entity"
	"context"
)


type IUpdateNameRepository interface {
	UpdateName (ctx context.Context,request entity.Shop)  error
}