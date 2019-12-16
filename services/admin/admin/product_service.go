package admin

//ProductService : Is in charge of communicate the api with the repository
type Service struct {
	Repository Repository
}

func ProvideAdminService(p Repository) Service {
	return Service{Repository: p}
}

func (p *Service) FindAll() []Admin {
	return p.Repository.FindAll()
}

func (p *Service) FindByID(id uint) Admin {
	return p.Repository.FindByID(id)
}

func (p *Service) FindByName(condition interface{}) Admin {
	return p.Repository.FindByName(condition)
}

func (p *Service) Save(admin Admin) Admin {
	p.Repository.Save(admin)

	return admin
}

func (p *Service) Delete(admin Admin) {
	p.Repository.Delete(admin)
}
