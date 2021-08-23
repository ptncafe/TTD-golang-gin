package update_name_shop_repo

import (
	dto2 "TTD-golang-gin-test/dto"
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

func (r repo) UpdateName (ctx context.Context,request dto2.UpdateNameRequest)  error{
	return nil
}