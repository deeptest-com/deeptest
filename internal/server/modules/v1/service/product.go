package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ProductService struct {
	ProductRepo *repo.ProductRepo `inject:""`
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) Paginate(req serverDomain.ProductReqPaginate) (ret domain.PageData, err error) {
	items := make([]*model.Product, 0)

	items, ret, err = s.ProductRepo.Paginate(req)
	if err != nil {
		return
	}

	// deal with nested items
	products := make([]*model.Product, 0)
	mp := map[uint]*model.Product{}
	for _, po := range items {
		mp[po.ID] = po
		if po.ParentId == 0 { // root
			products = append(products, po)
		}
	}

	for _, po := range items {
		if po.ParentId != 0 && mp[po.ParentId] != nil {
			mp[po.ParentId].Children = append(mp[po.ParentId].Children, po)
		}
	}

	ret.Result = products
	return
}

func (s *ProductService) FindById(id uint) (serverDomain.ProductResponse, error) {
	return s.ProductRepo.FindById(id)
}

func (s *ProductService) Create(req serverDomain.ProductRequest) (uint, error) {
	return s.ProductRepo.Create(req)
}

func (s *ProductService) Update(id uint, req serverDomain.ProductRequest) error {
	return s.ProductRepo.Update(id, req)
}

func (s *ProductService) DeleteById(id uint) error {
	return s.ProductRepo.BatchDelete(id)
}
