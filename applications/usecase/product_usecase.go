package usecase

import (
	"log"
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/security"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewProductUsecase(productRepository domains.ProductRepository,userRepo domains.UserRepository,) domains.ProductUsecase{
	return &ProductUsecaseImpl{
		productRepo: productRepository,
		userRepo: userRepo,
	}
}

type ProductUsecaseImpl struct {
	productRepo domains.ProductRepository
	userRepo domains.UserRepository	
}

func (productUsecase *ProductUsecaseImpl) CreateProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse) {
	var request model.ProductCreateRequest
	responseCode := make(chan string, 1)
	ctx.BodyParser(&request)

	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	userResult, err := productUsecase.userRepo.GetUserByEmail(email)
	if err != nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}


	product := entities.Product{
		Name: request.Name,
		Stock: request.Stock,
		UserId: userResult.Id,
	}
	err = productUsecase.productRepo.Insert(&product)
	if err != nil {
		responseCode<-"500"
		return <-responseCode,nil,nil
	}
	responseCode <- "200"
	return <-responseCode,nil,nil
}

func (ProductUsecase *ProductUsecaseImpl) FindProductById(ctx *fiber.Ctx,id string) (string,interface{},*model.ProductResponse){
	responseCode := make(chan string,1)
	
	var response = &model.ProductResponse{}
	
	email:=ctx.Locals("email").(string)

	isAdmin:=false
	userResult, err := ProductUsecase.userRepo.GetUserByEmail(email)
	for _, role := range userResult.Roles {
		if role.Name == "admin" {
			isAdmin = true
			break
		}
		isAdmin = false
	}

	result,err:= ProductUsecase.productRepo.FindById(id)
	if result == nil && err != nil {
		responseCode <- "404"
		return <-responseCode,nil,nil
	}
	log.Print(err)

	if !isAdmin {
		if userResult.Id != result.UserId {
			responseCode <- "403"
			return <-responseCode,nil,nil
		}else{
			if err == nil && result != nil {
				response = &model.ProductResponse{
					Id:             result.Id,
					Name:           result.Name,
					Stock: 			response.Stock,
					UserId: result.UserId,
					CreatedAt:      result.CreatedAt,
					UpdatedAt:      result.UpdatedAt,
				}
				responseCode<-"200"
				return <-responseCode,nil,response
			}
			responseCode <- "404"
			response = nil
			return <-responseCode,nil,nil
		}
	}


	if result.UserId != userResult.Id {
		responseCode <- "403"
		return <-responseCode,nil,nil
	}

	if err != nil && result == nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}

	if err == nil && result != nil {
		response = &model.ProductResponse{
			Id:             result.Id,
			Name:           result.Name,
			Stock: 			response.Stock,
			UserId: result.UserId,
			CreatedAt:      result.CreatedAt,
			UpdatedAt:      result.UpdatedAt,
		}
		responseCode<-"200"
		return <-responseCode,nil,response
	}
	responseCode<-"200"
	return <-responseCode,nil,response
}

func (ProductUsecase *ProductUsecaseImpl) EditProduct(ctx *fiber.Ctx)(string,interface{},*model.ProductResponse){
	responseCode := make(chan string,1)
	var request model.ProductCreateRequest
	id:=ctx.Params("id")
	ctx.BodyParser(&request)

	

	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	userResult, err := ProductUsecase.userRepo.GetUserByEmail(email)
	product,err := ProductUsecase.productRepo.FindById(id)

	product.Name = request.Name
	product.Stock = request.Stock

	isAdmin := false
	for _, role := range userResult.Roles {
		if role.Name == "admin" {
			isAdmin = true
			break
		}
		isAdmin = false
	}

	if !isAdmin {
		if product.UserId != userResult.Id {
			responseCode <- "403"
			return <-responseCode,nil,nil
		}
		entities,err := ProductUsecase.productRepo.Update(product)
		if err != nil {
			responseCode <- "500"
			return <-responseCode,nil,nil
		}
	
		response:= model.ProductResponse{
			Id:entities.Id,
			Name: entities.Name,
			Stock: entities.Stock,
			CreatedAt: entities.CreatedAt,
			UpdatedAt: entities.UpdatedAt,
		}	
		responseCode<-"200"
		return <-responseCode,nil,&response
	}
	
	 
	if err != nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}
	
	if product == nil && err == nil {
		responseCode <- "404"
		return <-responseCode,nil,nil
	}

	entities,err := ProductUsecase.productRepo.Update(product)
	
	response:= model.ProductResponse{
		Id:entities.Id,
		Name: entities.Name,
		Stock: entities.Stock,
		CreatedAt: entities.CreatedAt,
		UpdatedAt: entities.UpdatedAt,
	}
	responseCode<-"200"
	return <-responseCode,nil,&response
}

func (ProductUsecase *ProductUsecaseImpl) DeleteProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse){
	responseCode := make(chan string,1)
	id := ctx.Params("id")
	
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	Product,err := ProductUsecase.productRepo.FindById(id)

	userResult, err := ProductUsecase.userRepo.GetUserByEmail(email)

	isAdmin := false
	for _, role := range userResult.Roles {
		if role.Name == "admin" {
			isAdmin = true
			break
		}
		isAdmin = false
	}

	if !isAdmin {
		if Product.UserId != userResult.Id {
			responseCode <- "403"
			return <-responseCode,nil,nil
		}
		err = ProductUsecase.productRepo.Delete(id)
		if err != nil {
			responseCode <- "500"
		}else{
			responseCode<-"200"	
		}	
		return <-responseCode,nil,nil
	}

	if Product == nil && err == nil {
		responseCode <- "404"
		return <-responseCode,nil,nil
	}
	err = ProductUsecase.productRepo.Delete(id)
	if err != nil {
		responseCode <- "500"
	}else{
		responseCode<-"200"	
	}	
	return <-responseCode,nil,nil
}