package http

import (
	"github.com/sirupsen/logrus"
	"testing"
)

const (
	getUrl         = "https://www.curve.fi/raw-stats/apys.json"
	registerUrl    = "http://127.0.0.1:8080/user/register"
	updateUsername = "http://127.0.0.1:8080/user/updateUsername"
)

type (
	CurveYield struct {
		Apy struct {
			Day struct {
				ThreePool float64 `json:"3pool"`
				Busd      float64 `json:"busd"`
				Compound  float64 `json:"compound"`
				Dusd      float64 `json:"dusd"`
				Gusd      float64 `json:"gusd"`
				Hbtc      float64 `json:"hbtc"`
				Husd      float64 `json:"husd"`
				Linkusd   float64 `json:"linkusd"`
				Musd      float64 `json:"musd"`
				Pax       float64 `json:"pax"`
				Ren2      float64 `json:"ren2"`
				Rens      float64 `json:"rens"`
				Rsv       float64 `json:"rsv"`
				Susd      float64 `json:"susd"`
				Tbtc      float64 `json:"tbtc"`
				Usdk      float64 `json:"usdk"`
				Usdn      float64 `json:"usdn"`
				Usdt      float64 `json:"usdt"`
				Y         float64 `json:"y"`
			} `json:"day"`
			Month struct {
				ThreePool float64 `json:"3pool"`
				Busd      float64 `json:"busd"`
				Compound  float64 `json:"compound"`
				Dusd      int64   `json:"dusd"`
				Gusd      float64 `json:"gusd"`
				Hbtc      float64 `json:"hbtc"`
				Husd      float64 `json:"husd"`
				Linkusd   float64 `json:"linkusd"`
				Musd      float64 `json:"musd"`
				Pax       float64 `json:"pax"`
				Ren2      float64 `json:"ren2"`
				Rens      float64 `json:"rens"`
				Rsv       float64 `json:"rsv"`
				Susd      float64 `json:"susd"`
				Tbtc      int64   `json:"tbtc"`
				Usdk      float64 `json:"usdk"`
				Usdn      float64 `json:"usdn"`
				Usdt      float64 `json:"usdt"`
				Y         float64 `json:"y"`
			} `json:"month"`
			Total struct {
				ThreePool float64 `json:"3pool"`
				Busd      float64 `json:"busd"`
				Compound  float64 `json:"compound"`
				Dusd      float64 `json:"dusd"`
				Gusd      float64 `json:"gusd"`
				Hbtc      float64 `json:"hbtc"`
				Husd      float64 `json:"husd"`
				Linkusd   float64 `json:"linkusd"`
				Musd      float64 `json:"musd"`
				Pax       float64 `json:"pax"`
				Ren2      float64 `json:"ren2"`
				Rens      float64 `json:"rens"`
				Rsv       float64 `json:"rsv"`
				Susd      float64 `json:"susd"`
				Tbtc      float64 `json:"tbtc"`
				Usdk      float64 `json:"usdk"`
				Usdn      float64 `json:"usdn"`
				Usdt      float64 `json:"usdt"`
				Y         float64 `json:"y"`
			} `json:"total"`
			Week struct {
				ThreePool float64 `json:"3pool"`
				Busd      float64 `json:"busd"`
				Compound  float64 `json:"compound"`
				Dusd      int64   `json:"dusd"`
				Gusd      float64 `json:"gusd"`
				Hbtc      float64 `json:"hbtc"`
				Husd      float64 `json:"husd"`
				Linkusd   float64 `json:"linkusd"`
				Musd      float64 `json:"musd"`
				Pax       float64 `json:"pax"`
				Ren2      float64 `json:"ren2"`
				Rens      float64 `json:"rens"`
				Rsv       float64 `json:"rsv"`
				Susd      float64 `json:"susd"`
				Tbtc      float64 `json:"tbtc"`
				Usdk      float64 `json:"usdk"`
				Usdn      float64 `json:"usdn"`
				Usdt      float64 `json:"usdt"`
				Y         float64 `json:"y"`
			} `json:"week"`
		} `json:"apy"`
		Volume struct {
			ThreePool float64 `json:"3pool"`
			Busd      float64 `json:"busd"`
			Compound  float64 `json:"compound"`
			Dusd      float64 `json:"dusd"`
			Gusd      float64 `json:"gusd"`
			Hbtc      float64 `json:"hbtc"`
			Husd      float64 `json:"husd"`
			Musd      float64 `json:"musd"`
			Pax       float64 `json:"pax"`
			Ren2      float64 `json:"ren2"`
			Rens      float64 `json:"rens"`
			Rsv       float64 `json:"rsv"`
			Susd      float64 `json:"susd"`
			Tbtc      float64 `json:"tbtc"`
			Usdn      float64 `json:"usdn"`
			Usdt      float64 `json:"usdt"`
			Y         float64 `json:"y"`
		} `json:"volume"`
	}

	UserInfo struct {
		Code string `json:"code"`
		Data struct {
			CreatedAt string `json:"createdAt"`
			Email     string `json:"email"`
			ID        string `json:"id"`
			Username  string `json:"username"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	RegisterInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	LoginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateUsername struct {
		NewUsername string `json:"newUsername"`
	}

	Response struct {
		Code string `json:"code"`
		Data string `json:"data"`
		Msg  string `json:"msg"`
	}
)

func TestGet(t *testing.T) {
	res := new(CurveYield)
	if err := Get(getUrl, res); err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(res.Apy.Day.Busd)
}

func TestGetWithHeader(t *testing.T) {
	m := make(map[string]string)
	userInfo := new(UserInfo)
	m["Authorization"] = "d2b7e6f4-de6b-4180-9dd1-0fe2884d7d2d"
	err := Get("http://127.0.0.1:8080/user/info", userInfo, Param{Header: m})
	//err := Get("http://127.0.0.1:8080/user/info", userInfo) // failed
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Info(userInfo)
}

func TestGetWithQueryParam(t *testing.T) {
	m1 := make(map[string]string)
	m2 := make(map[string]string)
	m1["user"] = "1/% " // test url encoding (escape the special symbols)
	m2["Authorization"] = "d2b7e6f4-de6b-4180-9dd1-0fe2884d7d2d"
	userInfo := new(UserInfo)
	err := Get("http://127.0.0.1:8080/user/info", userInfo, Param{Query: m1, Header: m2})
	//err := Get("http://127.0.0.1:8080/user/info", userInfo) // failed
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Info(userInfo)
}

func TestPost(t *testing.T) {
	payload := &RegisterInfo{Email: "xx@xx.com", Username: "zhangsan", Password: "12345"}
	res := new(Response)
	err := Post(registerUrl, payload, res)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Info(res)

}

func TestPostWithHeader(t *testing.T) {
	payload := &UpdateUsername{NewUsername: "lisi"}
	res := new(Response)
	err := Post(updateUsername, payload, res, map[string]string{"Authorization": "651348e9-8f08-4b9b-aec6-08f49ccc19fd"})
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Info(res)
}
