package services

import "context"

type ProdService struct {

}

func(s *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error){
	return &ProdResponse{ProdStock: 20},nil
}

