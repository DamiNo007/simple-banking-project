package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			"1",
			"Damir",
			"Almaty",
			"0500060",
			"05.07.1999",
			"ACTIVE",
		},
		{
			"2",
			"Ruslan",
			"Almaty",
			"0500060",
			"22.03.2000",
			"ACTIVE",
		},
		{
			"3",
			"Daniyar",
			"Almaty",
			"0500070",
			"05.01.2000",
			"ACTIVE",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
