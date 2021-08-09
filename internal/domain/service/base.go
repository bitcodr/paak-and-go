package service

import "context"

type Service interface {
	list()
	show()
	store()
}


func New(_ context.Context, repository  service Service) Service{

}