package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"time"
)

func ListScopeCookie(processorId uint, execUuid string) (cookies []domain.ExecCookie) {
	scopeHierarchy := GetScopeHierarchy(execUuid)

	allValidIds := scopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *scopeHierarchy[processorId] {
			cookies = append(cookies, GetScopedCookies(execUuid)[id]...)
		}
	}

	return
}

func GetCookie(processorId uint, cookieName, domain string, execUuid string) (cookie domain.ExecCookie) {
	scopeHierarchy := GetScopeHierarchy(execUuid)
	scopedCookies := GetScopedCookies(execUuid)

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

func SetCookie(processorId uint, cookieName string, cookieValue interface{}, domainName string, expireTime *time.Time, execUuid string) (err error) {
	scopedCookies := GetScopedCookies(execUuid)

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

	SetScopedCookies(execUuid, scopedCookies)

	return
}

func ClearCookie(processorId uint, cookieName string, execUuid string) (err error) {
	scopedCookies := GetScopedCookies(execUuid)

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

		SetScopedCookies(execUuid, scopedCookies)
	}

	return
}
