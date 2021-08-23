package proxy

import (
	entity2 "TTD-golang-gin-test/entity"
)

type IPubSubProxy interface {
	PubShop(shop entity2.Shop) error
}