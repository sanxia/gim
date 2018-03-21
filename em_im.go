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
	org         OrgOption
	app         AppOption
	accessToken string
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
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
	}

	c.accessToken = tokenResponse.AccessToken

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
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
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
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
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
func (c *emChatClient) SendTextMessage(token, fromUsername string, toUsernames []string, content string) (*TextMessageResponse, error) {
	var err error

	messageUrl := fmt.Sprintf("%s/%s", applicationBaseUrl, "messages")
	authorization := fmt.Sprintf("Bearer %s", token)

	headers := map[string]string{
		"Content-Type":  "application/json;charset=utf-8",
		"Authorization": authorization,
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	//设置json数据
	requestData := new(TextMessageRequest)
	requestData.TargetType = "users"
	requestData.Target = toUsernames
	requestData.From = fromUsername
	requestData.Message = TextMessage{
		Type: "txt",
		Msg:  content,
	}

	request.SetJson(requestData)

	//发送请求
	httpResponse, err := request.Post(messageUrl)
	if err != nil {
		return nil, err
	}

	//解析数据
	data := string(httpResponse.GetData())
	log.Printf("TextMessage raw data: %s", data)

	var messageResponse *TextMessageResponse
	glib.FromJson(data, &messageResponse)

	//错误处理
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
	}

	return messageResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 发送图片消息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) SendImageMessage(token, fromUsername string, toUsernames []string, url, secret string, width, height int) (*ImageMessageResponse, error) {
	var err error

	messageUrl := fmt.Sprintf("%s/%s", applicationBaseUrl, "messages")
	authorization := fmt.Sprintf("Bearer %s", token)

	headers := map[string]string{
		"Content-Type":  "application/json;charset=utf-8",
		"Authorization": authorization,
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	//设置json数据
	filename := glib.RandString(16)
	requestData := new(ImageMessageRequest)
	requestData.TargetType = "users"
	requestData.Target = toUsernames
	requestData.From = fromUsername
	requestData.Message = ImageMessage{
		Type:     "img",
		Url:      url,
		Filename: filename,
		Secret:   secret,
		ImageSize: ImageSize{
			Width:  width,
			Height: height,
		},
	}

	request.SetJson(requestData)

	//发送请求
	httpResponse, err := request.Post(messageUrl)
	if err != nil {
		return nil, err
	}

	//解析数据
	data := string(httpResponse.GetData())
	log.Printf("ImageMessage raw data: %s", data)

	var messageResponse *ImageMessageResponse
	glib.FromJson(data, &messageResponse)

	//错误处理
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
	}

	return messageResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取用户的离线消息数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) GetOfflineMessageCount(token, username string) (*OfflineMessageCountResponse, error) {
	var err error
	userUrl := fmt.Sprintf("%s/users/%s/offline_msg_count", applicationBaseUrl, username)
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
	log.Printf("Offline raw data: %s", data)

	var offlineMessageCountResponse *OfflineMessageCountResponse
	glib.FromJson(data, &offlineMessageCountResponse)

	//错误处理
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
	}

	return offlineMessageCountResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 下线用户
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) Offline(token, username string) (*OfflineResponse, error) {
	var err error
	userUrl := fmt.Sprintf("%s/users/%s/disconnect", applicationBaseUrl, username)
	authorization := fmt.Sprintf("Bearer %s", token)

	//请求头
	headers := map[string]string{
		"Content-Type":  "application/json",
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
	log.Printf("Offline raw data: %s", data)

	var offlineResponse *OfflineResponse
	glib.FromJson(data, &offlineResponse)

	//错误处理
	var errorResponse *ErrorResponse
	glib.FromJson(data, &errorResponse)
	if len(errorResponse.Error) > 0 {
		err = errors.New(errorResponse.Error)
	}

	return offlineResponse, err
}
