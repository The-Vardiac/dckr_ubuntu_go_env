package main

/* A class should have only one reason to change */

/* bad example */
type UserService struct {}

func (u *UserService) Register(data map[string]string) error {
	// validate input
	// save to database
	// send email
}


/* good example */
type UserValidator struct {}
func (v *UserValidator) Validate(data map[string]string) error {
	return nil
}

type UserRepository struct {}
func (r *UserRepository) Save(data map[string]string) error {
	return nil
}

type Mailer struct {}
func (m *Mailer) SendWelcomeEmail() error {
	return nil
}