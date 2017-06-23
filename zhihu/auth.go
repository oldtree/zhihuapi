package zhihu

import (
	"encoding/json"
	"errors"
	"fmt"
	"fork/tools/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/juju/persistent-cookiejar"
)

type Result struct {
	R         int         `json:"r"`
	Msg       string      `json:"msg"`
	ErrorCode int         `json:"errcode"`
	Data      interface{} `json:"data"`
}

const limitnumber = 10

var (
	limit         int64
	MainUrl       = "https://www.zhihu.com"
	EmailLoginUrl = "https://www.zhihu.com/login/email"
	LogoutUrl     = "https://www.zhihu.com/logout"
	CheckUrl      = "https://www.zhihu.com/settings/profile"
	userAgent     = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36"
)

var client = http.Client{
	Timeout: time.Second * time.Duration(30),
}

func newHTTPHeaders(isXhr bool) http.Header {
	headers := make(http.Header)
	headers.Set("Accept", "*/*")
	headers.Set("Connection", "keep-alive")
	headers.Set("Host", "www.zhihu.com")
	headers.Set("Origin", "http://www.zhihu.com")
	headers.Set("Pragma", "no-cache")
	headers.Set("User-Agent", userAgent)
	headers.Set("Accept-Charset", "utf-8")
	if isXhr {
		headers.Set("X-Requested-With", "XMLHttpRequest")
	}
	return headers
}

type Auth struct {
	XsrfCode string
}

func (l *Auth) GetConfig(filepath string) (bool, error) {
	if filepath == "" {
		return false, errors.New("filepath params is empty")
	}

	err := Config.ReadConfigFile(filepath)
	if err != nil {
		return false, err
	}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return false, err
	}
	/////here limit the request rate///////
	go l.Bucket()
	return true, nil
}

func (l *Auth) GetXsrfCode() (string, error) {
	resp, err := client.Get(MainUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var _xsrfCode string
	for _, value := range resp.Cookies() {
		if value.Name == "_xsrf" {
			_xsrfCode = value.Value
			break
		} else {
			continue
		}
	}
	return _xsrfCode, nil
}

func (l *Auth) GetCaptchaAndInput() string {
	url := fmt.Sprintf("https://www.zhihu.com/captcha.gif?r=%d&type=login", 1000*time.Now().Unix())
	resp, err := client.Get(url)
	if err != nil {
		log.Error("获取验证码失败", err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error("获取验证码失败", resp.StatusCode)
		return ""
	}

	fileExt := strings.Split(resp.Header.Get("Content-Type"), "/")[1]
	captchaPath, _ := os.Getwd()
	verifyImg := filepath.Join(captchaPath, "verify."+fileExt)
	fd, err := os.OpenFile(verifyImg, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic("打开验证码文件失败：" + err.Error())
	}
	defer fd.Close()

	io.Copy(fd, resp.Body) // 保存验证码文件

	err = exec.Command("open", verifyImg).Run()
	if err != nil {
		log.Error("open captcha file error :", err)
	}
	var captcha string
	fmt.Print("请输入验证码：")
	fmt.Scanln(&captcha)
	return captcha
}

func (l *Auth) LogoutAction() error {
	logoutReq, _ := http.NewRequest("GET", LogoutUrl, nil)
	headers := newHTTPHeaders(true)
	logoutReq.Header = headers
	resp, err := client.Do(logoutReq)
	if err != nil {
		log.Error("logout error:", err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

//login operation
func (l *Auth) LoginAction() (bool, error) {
	isLogin, err := l.islogin()
	if isLogin {
		log.Info("user has been login")
		return true, nil
	}

	_xsrfCode, err := l.GetXsrfCode()
	if err != nil {
		return false, err
	}
	l.XsrfCode = _xsrfCode
	postformData := map[string]string{
		"_xsrf":       l.XsrfCode,
		"password":    Config.ZhihuPassword,
		"email":       Config.ZhihuAccount,
		"remember_me": "true",
		"captcha":     l.GetCaptchaAndInput(),
	}
	urlvalues := url.Values{}
	for key, value := range postformData {
		urlvalues.Set(key, value)
	}
	formData := urlvalues.Encode()
	loginReq, err := http.NewRequest("POST", EmailLoginUrl, strings.NewReader(formData))
	if err != nil {
		return false, err
	}
	headers := newHTTPHeaders(true)
	headers.Set("Content-Length", strconv.Itoa(len(formData)))
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Referer", MainUrl)
	loginReq.Header = headers

	resp, err := client.Do(loginReq)
	if resp == nil || err != nil {
		return false, err
	}
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	re := new(Result)
	json.Unmarshal(respData, re)
	log.Info(re)
	if re.R == 1 {
		log.Error(re.Msg)
		return false, errors.New(re.Msg)
	}
	client.Jar.(*cookiejar.Jar).Save()
	log.Info(resp.Status, ":", resp.StatusCode)
	return true, nil
}

func (l *Auth) Bucket() {
	limit = limitnumber
	tick := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-tick.C:
			atomic.StoreInt64(&limit, limitnumber)
		}
	}
}

func (l *Auth) GetData(targeturl string) ([]byte, error) {
	/*for {
		if limit <= 0 {
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}
	atomic.AddInt64(&limit, -1)*/
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	targetReq, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	headers := newHTTPHeaders(true)
	targetReq.Header = headers
	resp, err := client.Do(targetReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !strings.Contains(resp.Header.Get("Content-Type"), "json") {
		return nil, errors.New("page is not json format")
	}
	targetPageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	client.Jar.(*cookiejar.Jar).Save()
	return targetPageData, nil
}

func (l *Auth) islogin() (bool, error) {
	settingURL := MainUrl + "/settings/profile"
	settingReq, _ := http.NewRequest("GET", settingURL, nil)
	headers := newHTTPHeaders(true)
	settingReq.Header = headers
	resp, err := client.Do(settingReq)
	if err != nil {
		log.Error("访问 profile 页面出错: " + err.Error())
		return false, err
	}
	defer resp.Body.Close()
	lastURL := resp.Request.URL.String()
	log.Info(lastURL)
	if lastURL == settingURL {
		return true, nil
	}
	return false, errors.New("not login yet")
}
