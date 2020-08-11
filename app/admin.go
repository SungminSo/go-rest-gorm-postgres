package app

func (app *ProjectApp) Register(adminID, password string) (string, error) {
	_, err := app.admins.FindByID(adminID)
}