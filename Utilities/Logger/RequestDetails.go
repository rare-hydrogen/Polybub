package Logger

type RequestDetails struct {
	UserId      int32
	UserName    string
	UserGroupId int32
	Verb        string
	Endpoint    string
	QueryString string
	Body        string
}
