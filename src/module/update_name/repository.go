package update_shop_info

import (
	"TTD-golang-gin-test/model/dto"
	"TTD-golang-gin-test/provider/mongo_provider"
	"context"
	"github.com/sirupsen/logrus"
)

type repo struct {
	Log                     *logrus.Logger
	IMongoProvider mongo_provider.IMongoProvider
}

func NewRepository(
	log *logrus.Logger,
	mongoProvider mongo_provider.IMongoProvider,
	) *repo {
	return &repo{
		Log: log,
		IMongoProvider: mongoProvider,
	}
}

func (r repo) Update (ctx context.Context,request dto.UpdateNameRequest)  error{
	return nil
}