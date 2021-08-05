package imodule

import (
	"TTD-golang-gin-test/model/dto"
	"context"
)

type IUpdateNameDomain interface {
	UpdateName (ctx context.Context,request dto.UpdateNameRequest)  error
}

type IUpdateNameRepository interface {
	UpdateName (ctx context.Context,request dto.UpdateNameRequest)  error
}