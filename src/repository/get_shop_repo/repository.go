package get_shop_repo

import (
	 "TTD-golang-gin-test/entity"
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

func (r repo) GetByName (ctx context.Context, name string)  (*entity.Shop,error){
	return nil,nil
}
func GetByCode (ctx context.Context, code string)  (*entity.Shop,error){
	return nil,nil
}