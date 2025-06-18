package models

type CustomerCardCreate struct {
	Surname     *string
	Name        *string
	Patronymic  *string
	PhoneNumber *string
	City        *string
	Street      *string
	ZipCode     *string
	Percent     *int
}

type CustomerCardRetrieve struct {
	CardNumber  *string
	Surname     *string
	Name        *string
	Patronymic  *string
	PhoneNumber *string
	City        *string
	Street      *string
	ZipCode     *string
	Percent     *int
}

type CustomerCardUpdate struct {
	Surname     *string
	Name        *string
	Patronymic  *string
	PhoneNumber *string
	City        *string
	Street      *string
	ZipCode     *string
	Percent     *int
}
