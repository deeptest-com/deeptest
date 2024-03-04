package serverConsts

const (
	AdminUserName     = "admin"
	AdminUserPassword = "P2ssw0rd"
	AdminRoleName     = "admin"
)

var (
	SortMap = map[string]string{
		"ascend":  "ASC",
		"descend": "DESC",
		"":        "ASC",
	}
)

const (
	Designing  = 1
	Developing = 2
	Published  = 3
	Abandoned  = 4
)

const DefaultSever = "http://localhost"

const IsAdminRole = "YES"
const IsNotAdminRole = "NO"
