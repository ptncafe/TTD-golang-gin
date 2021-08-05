package get_shop_repo

import (
	"TTD-golang-gin-test/model/entity"
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

func (r repo) GetById (ctx context.Context, id int)  (*entity.Shop,error){
	return nil,nil
}