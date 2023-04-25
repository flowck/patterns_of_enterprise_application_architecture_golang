package optimistic_offline_lock

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Email     string

	version int
}

func (c *Customer) Edit(firstName, lastName, email string) {
	c.Email = email
	c.FirstName = firstName
	c.LastName = lastName
}

func (c *Customer) Version() int {
	return c.version
}
