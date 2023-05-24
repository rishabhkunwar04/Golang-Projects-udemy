package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"111", "Rishabh", "Pune", "424232", "4-10-2000", "1"},
		{"121", "Itachi", "Konoha", "424232", "4-10-2000", "1"},
	}
	return CustomerRepositoryStub{customers}
}
