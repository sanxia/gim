package gim

/* ================================================================================
 * Chat Client
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 选项
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type OrgOption struct {
	OrgName string //机构名
	AppName string //应用名
}

type AppOption struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Token请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TokenRequest struct {
	GrantType string `json:"grant_type"`
	AppOption
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Token响应
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Application string `json:"application"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户请求数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type CreateUserResponse struct {
	*BaseResult
	Entities []*CreateUserResult `json:"entities"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 基础结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type BaseResult struct {
	Action          string      `json:"action"`
	Application     string      `json:"application"`
	Params          interface{} `json:"params"`
	Path            string      `json:"path"`
	Uri             string      `json:"uri"`
	Timestamp       int         `json:"timestamp"`
	Duration        int         `json:"duration"`
	Organization    string      `json:"organization"`
	ApplicationName string      `json:"applicationName"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type CreateUserResult struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建用户响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type GetUserResponse struct {
	*BaseResult
	Entities []*GetUserResult `json:"entities"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用户结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type GetUserResult struct {
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	Activated bool   `json:"activated"`
	Type      string `json:"type"`
	Created   string `json:"created"`
	Modified  string `json:"modified"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 重置密码响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ResetPasswordResponse struct {
	Action    string `json:"action"`
	Duration  int    `json:"duration"`
	Timestamp int    `json:"timestamp"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 文本消息请求数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TextMessageRequest struct {
	MessageTarget
	TextMessage
	From string `json:"from"` //表示消息发送者。无此字段Server会默认设置为"from":"admin"
}

type MessageTarget struct {
	TargetType string   `json:"target_type"` //users 给用户发消息。chatgroups: 给群发消息，chatrooms: 给聊天室发消息
	Target     []string `json:"target"`      //即使只有一个用户也要用数组 ['u1']，给用户发送时数组元素是用户名，给群组发送时数组元素是groupid
}

type TextMessage struct {
	Type string   `json:"type"` //消息类型 txt
	Msg  []string `json:"msg"`  //消息内容
}

type TextMessageResponse struct {
	*BaseResult
	Entities []interface{}     `json:"entities"`
	Data     map[string]string `json:"data"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 图片消息请求数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ImageMessageRequest struct {
	MessageTarget
	ImageMessage
	From string `json:"from"` //表示消息发送者。无此字段Server会默认设置为"from":"admin"
}

type ImageMessage struct {
	Type      string    `json:"type"`     //消息类型 img
	Url       string    `json:"url"`      //成功上传文件返回的UUID
	Filename  string    `json:"filename"` //指定一个文件名
	Secret    string    `json:"secret"`   //成功上传文件后返回的secret
	ImageSize ImageSize `json:"size"`
}

type ImageSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 图片消息响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ImageMessageResponse struct {
	*BaseResult
	Entities []interface{}     `json:"entities"`
	Data     map[string]string `json:"data"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取离线消息数响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type OfflineMessageCountResponse struct {
	Action    string         `json:"action"`
	Entities  []interface{}  `json:"entities"`
	Data      map[string]int `json:"data"`
	Duration  int            `json:"duration"`
	Timestamp int            `json:"timestamp"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用户下线消息数响应数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type OfflineResponse struct {
	Action    string         `json:"action"`
	Data      *OfflineResult `json:"data"`
	Duration  int            `json:"duration"`
	Timestamp int            `json:"timestamp"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 下线结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type OfflineResult struct {
	Result bool `json:"result"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据错误数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ResponseError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Exception        string `json:"exception"`
	Duration         int    `json:"duration"`
	Timestamp        int    `json:"timestamp"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * IChatClient接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type IChatClient interface {
	IsOnline(username string) (bool, error)

	GetAccessToken() (*TokenResponse, error)
	GetUser(token, username string) (*GetUserResponse, error)

	CreateUser(token, username, password string) (*CreateUserResponse, error)
	CreateUsers(token string, requestData []*CreateUserRequest) (*CreateUserResponse, error)
	ResetPassword(username string) (*ResetPasswordResponse, error)

	SendTextMessage(requestData *TextMessageRequest) (*TextMessageResponse, error)
	SendImageMessage(requestData *ImageMessageRequest) (*ImageMessageResponse, error)

	GetOfflineMessageCount(username string) (*OfflineMessageCountResponse, error)
	Offline(username string) (*OfflineResponse, error)
}
