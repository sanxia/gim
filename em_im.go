package gim

import (
//"fmt"
//"log"
//"strings"
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
	const applicationBaseUrl string = "https://a1.easemob.com/easemob-playground/test1"
	const applicationTokenUrl string = applicationBaseUrl + "/token"

	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	request := glib.NewHttpRequest()
	request.SetHeaders(headers)

	tokenRequest := new(TokenRequest)
	tokenRequest.GrantType = "client_credentials"
	tokenRequest.ClientId = c.app.ClientId
	tokenRequest.ClientSecret = c.app.ClientSecret

	request.SetJson(tokenRequest)

	httpResponse, err := request.Post(applicationTokenUrl)
	if err != nil {
		return nil, err
	}

	var tokenResponse *TokenResponse
	glib.FromJson(string(httpResponse.GetData()), &tokenResponse)

	return tokenResponse, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取用户数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) GetUser(username string) (*GetUserResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) CreateUser(request []*CreateUserRequest) (*CreateUserResponse, error) {

	return nil, nil
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
func (c *emChatClient) SendTextMessage(request *TextMessageRequest) (*TextMessageResponse, error) {

	return nil, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 发送图片消息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (c *emChatClient) SendImageMessage(request *ImageMessageRequest) (*ImageMessageResponse, error) {

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
