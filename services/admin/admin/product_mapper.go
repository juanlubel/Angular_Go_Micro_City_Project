package admin

//ToAdmin : mapper single objectDTO to gorm struct
func ToAdmin(adminDTO AdminDTO) Admin {
	return Admin{Name: adminDTO.Name, Pass: adminDTO.Pass, Token: adminDTO.Token}
}

/*ToAdminDTO : mapper to gorm struct to single object DTO */
func ToAdminDTO(admin Admin) AdminDTO {
	return AdminDTO{ID: admin.ID, Name: admin.Name, Pass: admin.Pass}
}

func ToLoggedAdminDTO(admin Admin) LoggedDTO {
	return LoggedDTO{Name: admin.Name, Token: admin.Token}
}

