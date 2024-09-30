package commService

import (
	"errors"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/go-ldap/ldap"
)

type LdapService struct {
}

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
	ErrDial               = errors.New("LDAP链接错误")
	ErrSearch             = errors.New("LDAP用户信息查询失败")
	ErrSearchNil          = errors.New("LDAP用户信息查询为空")
)

func (s *LdapService) Login(req v1.LoginReq) (err error) {
	l, err := ldap.Dial("tcp", "ldap-admin.nancalcloud.com:389")
	if err != nil {
		logUtils.Errorf("LDAP dial error: %s", err)
		err = ErrDial
		return
	}
	defer l.Close()
	bindDN := fmt.Sprintf("uid= %s,ou=people,dc=nancalcloud,dc=com", req.Username)
	err = l.Bind(bindDN, req.Password)
	if err != nil {
		logUtils.Errorf("LDAP bind error: %s,username: %s", err, req.Username)
		err = ErrUserNameOrPassword
		return
	}
	return
}

func (s *LdapService) LdapUserInfo(req v1.LoginReq) (userBase v1.UserBase, err error) {
	l, err := ldap.Dial("tcp", "ldap-admin.nancalcloud.com:389")
	if err != nil {
		logUtils.Errorf("LDAP dial error: %s", err)
		err = ErrDial
		return
	}
	defer l.Close()
	bindDN := fmt.Sprintf("uid= %s,ou=people,dc=nancalcloud,dc=com", req.Username)
	err = l.Bind(bindDN, req.Password)
	if err != nil {
		logUtils.Errorf("LDAP bind error: %s,username: %s", err, req.Username)
		err = ErrUserNameOrPassword
		return
	}

	baseDN := "uid=" + req.Username + ",ou=people,dc=nancalcloud,dc=com"
	scope := ldap.ScopeWholeSubtree
	filter := "(&(objectClass=inetOrgPerson)(uid=" + req.Username + "))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base DN to search
		scope,  // The scope of the search
		ldap.NeverDerefAliases,
		0, // Return all attributes
		0, // No time limit
		false,
		filter,
		nil,
		nil,
	)

	searchResult, err := l.Search(searchRequest)
	if err != nil {
		logUtils.Errorf("LDAP Search error: %s,username: %s", err, req.Username)
		err = ErrSearch
	}

	if len(searchResult.Entries) == 0 {
		logUtils.Errorf("No results found by username: %s", err, req.Username)
		err = ErrSearchNil
	} else {
		// 搜索到正确的用户,可以解析用户或者使用其他信息
		for _, entry := range searchResult.Entries {
			userBase.Name = entry.GetAttributeValue("displayName")
			userBase.Username = entry.GetAttributeValue("uid")
			userBase.Email = entry.GetAttributeValue("mail")
			userBase.ImAccount = entry.GetAttributeValue("title") //企业微信账号
			for _, attribute := range entry.Attributes {
				fmt.Printf("%s: %s\n", attribute.Name, attribute.Values)
			}
		}
	}
	return
}
