package idomain

import (
	dto2 "TTD-golang-gin-test/dto"
	"context"
)

type IUpdateNameDomain interface {
	UpdateName (ctx context.Context,request dto2.UpdateNameRequest)  error
}

