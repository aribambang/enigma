package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Ari", "Bekasi", "17156", "2000-01-01", "1"},
		{"2", "Bambang", "Bekasi", "17411", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
