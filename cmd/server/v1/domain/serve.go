package domain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ServeReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type ServeReq struct {
	ProjectId   int64  `json:"projectId"`
	Name        string `json:"name"`
	Status      int64  `json:"status"`
	ID          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	Description string `json:"description"`
}

type ServeVersionReq struct {
	ServeId     int64  `json:"serveId"`
	value       string `json:"value"`
	ID          uint   `json:"id"`
	UserId      int64  `json:"userId"`
	Description string `json:"description"`
}
