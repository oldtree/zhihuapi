package zhihu

import (
	"errors"
	"fork/tools/log"
)

var Client *ZhihuClient

type ZhihuClient struct {
	Authed *Auth `json:"authed,omitempty"`
}

func NewZhihuClient() *ZhihuClient {
	Client = &ZhihuClient{
		Authed: new(Auth),
	}
	return Client
}

func (zhihu *ZhihuClient) Get(url string, params string) (data []byte, err error) {
	defer func() {
		if re := recover(); re != nil {
			data = nil
			err = errors.New("http request failed")
		}
	}()
	data, err = zhihu.Authed.GetData(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (zhihu *ZhihuClient) Post(url string, params map[string]interface{}) (data []byte, err error) {
	defer func() {
		if re := recover(); re != nil {
			data = nil
			err = errors.New("http request failed")
		}
	}()
	return nil, nil
}

func (zhihu *ZhihuClient) Put(url string, paramas map[string]interface{}) (data []byte, err error) {
	defer func() {
		if re := recover(); re != nil {
			data = nil
			err = errors.New("http request failed")
		}
	}()
	return nil, nil
}

func (zhihu *ZhihuClient) Delete(url string, params map[string]interface{}) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New("http request failed")
		}
	}()
	return nil
}

func (zhihu *ZhihuClient) Login() error {
	_, err := zhihu.Authed.LoginAction()
	if err != nil {
		return err
	}
	return nil
}

func (zhihu *ZhihuClient) Logout() error {
	err := zhihu.Authed.LogoutAction()
	if err != nil {
		return err
	}
	return nil
}

func (zhihu *ZhihuClient) Init(cfgfile string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			err = errors.New("http request failed")
		}
	}()
	_, err = zhihu.Authed.GetConfig(cfgfile)
	if err != nil {
		log.Info("load config file error : ", err)
		return err
	}
	return nil
}
