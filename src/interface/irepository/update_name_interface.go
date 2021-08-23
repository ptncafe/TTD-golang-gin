package irepository

import (
	dto2 "TTD-golang-gin-test/dto"
	"context"
)


type IUpdateNameRepository interface {
	UpdateName (ctx context.Context,request dto2.UpdateNameRequest)  error
}