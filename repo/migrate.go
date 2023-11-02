package repo

func (c *Conn) AutoMigrate() error {
	err := c.db.Migrator().AutoMigrate(&User{}, &Company{}, &Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
