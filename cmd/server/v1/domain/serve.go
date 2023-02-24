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
	ProjectId   int64  `json:"projectId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Status      int64  `json:"status"`
	ID          int64  `json:"id"`
	CreateUser  string `json:"createUser"`
	Description string `json:"description"`
}

type ServeVersionReq struct {
	ServeId     int64  `json:"serveId"`
	Value       string `json:"value"`
	ID          uint   `json:"id"`
	CreateUser  string `json:"createUser"`
	Description string `json:"description"`
}

type ServeVersionPaginate struct {
	_domain.PaginateReq
	ServeId    int64  `json:"serveId"`
	Version    string `json:"version"`
	CreateUser string `json:"createUser"`
}

type ServeSchemaReq struct {
	ID       uint   `json:"id"`
	ServeId  int64  `json:"serveId"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Content  string `json:"content"`
	Examples string `json:"examples"`
	Type     string `json:"type"`
	Tags     string `json:"tags"`
}

type ServeServerReq struct {
	ID          uint   `json:"id"`
	ServeId     int64  `json:"serveId"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type ServeSchemaPaginate struct {
	_domain.PaginateReq
	ServeId int64  `json:"serveId"`
	Tag     string `json:"tag"`
}

type JsonContent struct {
	Data string `json:"data"`
}

type SchemaContent struct {
	Data string
}
