package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"time"
)

func ListScopeCookie(processorId uint) (cookies []domain.ExecCookie) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		if ScopeHierarchy[processorId] != nil {
			for _, id := range *ScopeHierarchy[processorId] {
				cookies = append(cookies, ScopedCookies[id]...)
			}
		}
	}

	return
}

func GetCookie(processorId uint, cookieName, domain string) (cookie domain.ExecCookie) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		if ScopeHierarchy[processorId] != nil {
			for _, id := range *ScopeHierarchy[processorId] {
				for _, item := range ScopedCookies[id] {
					if item.Name == cookieName && (item.Domain == "" || domain == "" || item.Domain == domain) &&
						(item.ExpireTime == nil || item.ExpireTime.Unix() > time.Now().Unix()) {
						cookie = item

						goto LABEL
					}
				}
			}
		}
	}

LABEL:

	return
}

func SetCookie(processorId uint, cookieName string, cookieValue string, domainName string, expireTime *time.Time) (err error) {
	found := false

	newCookie := domain.ExecCookie{
		Name:  cookieName,
		Value: cookieValue,

		Domain:     domainName,
		Path:       "/",
		ExpireTime: expireTime,
	}

	for i := 0; i < len(ScopedCookies[processorId]); i++ {
		if ScopedCookies[processorId][i].Name == cookieName {
			ScopedCookies[processorId][i] = newCookie

			found = true
			break
		}
	}

	if !found {
		ScopedCookies[processorId] = append(ScopedCookies[processorId], newCookie)
	}

	return
}

func ClearCookie(processorId uint, cookieName string) (err error) {
	deleteIndex := -1
	for index, item := range ScopedCookies[processorId] {
		if item.Name == cookieName {
			deleteIndex = index
			break
		}
	}

	if deleteIndex > -1 {
		ScopedCookies[processorId] = append(
			ScopedCookies[processorId][:deleteIndex], ScopedCookies[processorId][(deleteIndex+1):]...)
	}

	return
}
