package api

import (
	pb "common/proto/catalogue_service"
	"users_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	productPb := &pb.User{
		Id            user.Id.Hex(),
		Name          user.Name,
		Lastname      user.Lastname,
		Email          user.Email,
		PhoneNumber	   user.PhoneNumber,
		Gender		   user.Gender,
		BirthDate	   user.BirthDate,
		Username	   user.Username,
		Biography	   user.Biography
	}
	return productPb
}
