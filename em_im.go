package gim

import (
	"errors"
	"fmt"
	"log"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * EM Chat Client
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * emChatClient数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type emChatClient struct {
	org OrgOption
	app AppOption
}

const applicationBaseUrl string = "https://a1.easemob.com/easemob-playground/test1"

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化emChatClient
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewEmChatClient(orgName, appName, clientId, clientSecret string) IChatClient {
	client := &emChatClient{
		org: OrgOption{
			OrgName: orgName,
			AppName: appName,
		},
		app: AppOption{
			ClientId:     clientId,
			ClientSecret: clientSecret,
		},
	}

	return client
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断用户是否在线
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) IsOnline(username string) (bool, error) {

	return false, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取AccessToken
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) GetAccessToken() (*TokenResponse, error) {
	var err error
	applicationTokenUrl := applicationBaseUrl + "/token"

	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	//设置json数据
	tokenRequest := new(TokenRequest)
	tokenRequest.GrantType = "client_credentials"
	tokenRequest.ClientId = c.app.ClientId
	tokenRequest.ClientSecret = c.app.ClientSecret
	request.SetJson(tokenRequest)

	//发送请求
	httpResponse, err := request.Post(applicationTokenUrl)
	if err != nil {
		return nil, err
	}

	//解析数据
	data := string(httpResponse.GetData())
	log.Printf("GetAccessToken raw data: %s", data)

	var tokenResponse *TokenResponse
	glib.FromJson(data, &tokenResponse)

	//错误处理
	if tokenResponse == nil {
		var responseError *ResponseError
		glib.FromJson(data, &responseError)
		if responseError != nil {
			err = errors.New(responseError.Error)
		}
	}

	return tokenResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取用户数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) GetUser(token, username string) (*GetUserResponse, error) {
	var err error
	userUrl := fmt.Sprintf("%s/%s/%s", applicationBaseUrl, "users", username)
	authorization := fmt.Sprintf("Bearer %s", token)

	//请求头
	headers := map[string]string{
		"Authorization": authorization,
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	//发送请求
	httpResponse, err := request.Get(userUrl)
	if err != nil {
		return nil, err
	}

	//解析数据
	data := string(httpResponse.GetData())
	log.Printf("GetUser raw data: %s", data)

	var getUserResponse *GetUserResponse
	glib.FromJson(data, &getUserResponse)

	//错误处理
	if getUserResponse == nil {
		var responseError *ResponseError
		glib.FromJson(data, &responseError)
		if responseError != nil {
			err = errors.New(responseError.Error)
		}
	}

	return getUserResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) CreateUser(token, username, password string) (*CreateUserResponse, error) {
	users := make([]*CreateUserRequest, 0)
	user := &CreateUserRequest{
		Username: username,
		Password: password,
	}
	users = append(users, user)

	return c.CreateUsers(token, users)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) CreateUsers(token string, requestData []*CreateUserRequest) (*CreateUserResponse, error) {
	var err error

	createUserUrl := fmt.Sprintf("%s/%s", applicationBaseUrl, "users")
	authorization := fmt.Sprintf("Bearer %s", token)

	headers := map[string]string{
		"Authorization": authorization,
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	//设置json数据
	request.SetJson(requestData)

	//发送请求
	httpResponse, err := request.Post(createUserUrl)
	if err != nil {
		return nil, err
	}

	//解析数据
	data := string(httpResponse.GetData())
	log.Printf("CreateUsers raw data: %s", data)

	var createUserResponse *CreateUserResponse
	glib.FromJson(data, &createUserResponse)

	//错误处理
	if createUserResponse == nil {
		log.Print("CreateUsers createUserResponse nil")
		var responseError *ResponseError
		glib.FromJson(data, &responseError)
		log.Printf("CreateUsers responseError: %v", responseError)
		if responseError != nil {
			err = errors.New(responseError.Error)
		}
	} else {
		log.Printf("CreateUsers createUserResponse not nil: %v", createUserResponse)
	}

	return createUserResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 重置用户密码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) ResetPassword(username string) (*ResetPasswordResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 发送文本消息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) SendTextMessage(requestData *TextMessageRequest) (*TextMessageResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 发送图片消息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) SendImageMessage(requestData *ImageMessageRequest) (*ImageMessageResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取用户的离线消息数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) GetOfflineMessageCount(username string) (*OfflineMessageCountResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 下线用户
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) Offline(username string) (*OfflineResponse, error) {

	return nil, nil
}
