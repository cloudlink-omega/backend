package structs

// Email configuration parameters
type MailConfig struct {
	Port     int
	Server   string
	Username string
	Password string
}

// Arguments needed to send an email
type EmailArgs struct {
	Template string // Name of template to use (/pkg/data/email_templates/*.html)
	Subject  string // Subject of email
	To       string // Email address of recipient (to user)
	Nickname string // Nickname of sender (i.e. this server)
}

// Template data to fill in for specified email to send
type TemplateData struct {
	Name                 string
	ServerName           string
	DeveloperOwner       string
	DeveloperName        string
	DeveloperDescription string
}
