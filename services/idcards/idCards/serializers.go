package idCards

func ToUsersDTO(model UserModel) UserDTO {
	return UserDTO{
		model.Name,
		model.Surname,
		model.Slug,
		model.Email,
		model.NCard}
}

func ToUserLoggedDTO(model UserModel, token string) UserLoggedDTO {
	return UserLoggedDTO{
		Name: model.Name,
		Slug:  model.Slug,
		NCard: model.NCard,
		Token: token,
	}
}