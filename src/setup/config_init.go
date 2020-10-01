package setup

// Init setup
func Init() {
	LoadEnvDevelopment()
	db, err := InitDatabase()
	if err != nil {
		println(err.Error())
	}
	RouterNoSetMode(true, db)
}
