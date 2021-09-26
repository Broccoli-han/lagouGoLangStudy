package ch0525

import (
	"errors"
	"regexp"
)

//
type Detail struct {
	Username string
	Email    string
}

//
type GeyPersonDetailRsp struct {
	RetCode int32   `json:"retCode"`
	Result  *Detail `json:"result"`
	RetMsg  string  `json:"retMsg"`
}

//
func checkUsername(username string) bool {
	const pattern = `^[a-z0-9_-]{3,16}`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

//
func checkEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9z_-]+(\.[a-zA-Z0-9_-]+)+$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//
func getPersonDetailHttp(username string) (*Detail, error) {
	//ctx := trpc.BackgroundContext()
	//
	//proxy := thttp.NewClientProxy("trpc.person.svr.cgi")
	//
	rsp := &GeyPersonDetailRsp{}
	//err := proxy.Get(ctx, fmt.Sprintf("/getPersonDetail?username=%v", username), rsp)
	//if err != nil {
	//	return nil, err
	//}
	return rsp.Result, nil
}

//
func GetPersonDetail(username string) (*Detail, error) {
	if ok := checkUsername(username); !ok {
		return nil, errors.New("invalid username")
	}

	// 从http接口获取信息
	detail, err := getPersonDetailHttp(username)

	if err != nil {
		return nil, err
	}

	//校验
	if ok := checkEmail(detail.Email); !ok {
		return nil, errors.New("invalid email")
	}

	return detail, nil
}
