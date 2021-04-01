package service

import (
	"go-clean/entity"
	"go-clean/model"
	"go-clean/repository"
	"go-clean/validation"
)

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (p *productServiceImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.ValidateProduct(request)
	product := entity.Product{
		Id:        	request.Id,
		Code:      	request.Code,
		Name:    	request.Name,
		Price:     	request.Price,
		Avaliable: 	request.Avaliable,
		Stock:     	request.Stock,
	}

	p.ProductRepository.Insert(product)
	response = model.CreateProductResponse{
		Id:        product.Id,
		Code:      product.Code,
		Name:      product.Name,
		Price:     product.Price,
		Avaliable: product.Avaliable,
		Stock:     product.Stock,
	}
	return response
}

func (p *productServiceImpl) List() (responses []model.GetProductResponse) {
	products := p.ProductRepository.FindAll()
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:        product.Id,
			Code:      product.Code,
			Name:      product.Name,
			Price:     product.Price,
			Avaliable: product.Avaliable,
			Stock:     product.Stock,
		})
	}
	return responses
}