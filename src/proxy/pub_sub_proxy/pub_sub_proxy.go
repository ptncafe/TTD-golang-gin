package pub_sub_proxy

import (
	entity2 "TTD-golang-gin-test/entity"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

type proxy struct {
	Log    *logrus.Logger
	Domain string
	Client *resty.Client
}

func NewProxy(
	log *logrus.Logger,
	domain string,
	client *resty.Client,
) *proxy {
	return &proxy{
		Log: log,
		Domain: domain,
		Client: client,
	}
}

func(p *proxy) PubShop(shop entity2.Shop) error{
	url :=fmt.Sprintf("%s/pub/store.update.info", p.Domain)
	resp, err := p.Client.R().SetBody(shop).
		Post(url)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		return nil
	}
	return errors.New("Pub thông tin shop không thành công")
}