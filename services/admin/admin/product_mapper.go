package admin

//ToAdmin : mapper single objectDTO to gorm struct
func ToAdmin(adminDTO AdminDTO) Admin {
	return Admin{Name: adminDTO.Name, Pass: adminDTO.Pass}
}

/*ToAdminDTO : mapper to gorm struct to single object DTO */
func ToAdminDTO(admin Admin) AdminDTO {
	return AdminDTO{ID: admin.ID, Name: admin.Name, Pass: admin.Pass}
}

func ToLoggedAdminDTO(admin Admin, token string) LoggedDTO {
	return LoggedDTO{Name: admin.Name, Token: token}
}

