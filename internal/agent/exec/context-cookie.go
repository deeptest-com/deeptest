package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"time"
)

func ListScopeCookie(session *ExecSession, processorId uint) (cookies []domain.ExecCookie) {
	scopeHierarchy := session.ScopeHierarchy
	scopedCookies := session.ScopedCookies

	allValidIds := scopeHierarchy[processorId]
	if allValidIds != nil {
		if scopeHierarchy[processorId] != nil {
			for _, id := range *scopeHierarchy[processorId] {
				cookies = append(cookies, scopedCookies[id]...)
			}
		}
	}
	return
}

func GetCookie(session *ExecSession, processorId uint, cookieName, domain string) (cookie domain.ExecCookie) {
	scopeHierarchy := session.ScopeHierarchy
	scopedCookies := session.ScopedCookies

	allValidIds := scopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *scopeHierarchy[processorId] {
			for _, item := range scopedCookies[id] {
				if item.Name == cookieName && (item.Domain == "" || domain == "" || item.Domain == domain) &&
					(item.ExpireTime == nil || item.ExpireTime.Unix() > time.Now().Unix()) {
					cookie = item

					goto LABEL
				}
			}
		}
	}

LABEL:

	return
}

func SetCookie(session *ExecSession, processorId uint, cookieName string, cookieValue string, domainName string, expireTime *time.Time) (err error) {
	scopedCookies := session.ScopedCookies

	found := false

	newCookie := domain.ExecCookie{
		Name:  cookieName,
		Value: cookieValue,

		Domain:     domainName,
		Path:       "/",
		ExpireTime: expireTime,
	}

	for i := 0; i < len(scopedCookies[processorId]); i++ {
		if scopedCookies[processorId][i].Name == cookieName {
			scopedCookies[processorId][i] = newCookie

			found = true
			break
		}
	}

	if !found {
		scopedCookies[processorId] = append(scopedCookies[processorId], newCookie)
	}

	return
}

func ClearCookie(session *ExecSession, processorId uint, cookieName string) (err error) {
	scopedCookies := session.ScopedCookies

	deleteIndex := -1
	for index, item := range scopedCookies[processorId] {
		if item.Name == cookieName {
			deleteIndex = index
			break
		}
	}

	if deleteIndex > -1 {
		scopedCookies[processorId] = append(
			scopedCookies[processorId][:deleteIndex], scopedCookies[processorId][(deleteIndex+1):]...)
	}

	return
}
