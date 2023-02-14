package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		user := User{
			Name:        "",
			Email:       "Aon@mail.com",
			StudentID:   "B6310264",
			PhoneNumber: "0842298645",
			Password:    "123456789",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})

	t.Run("check Email", func(t *testing.T) {
		user := User{
			Name:        "Aon",
			Email:       "Aoncom",
			StudentID:   "B6310264",
			PhoneNumber: "0842298645",
			Password:    "123456789",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Email invalid format"))
	})

	t.Run("check StudentID", func(t *testing.T) {

		sid := []string{
			"XA0000000",
			"BA0000000",
			"B123456",
			"B12345678",
		}
		for _, val := range sid {

			user := User{
				Name:        "Aon",
				Email:       "Aon@mail.com",
				StudentID:   val,
				PhoneNumber: "0842298645",
				Password:    "123456789",
			}

			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("StudentID invalid"))

		}

	})

	t.Run("check Phone Nunber", func(t *testing.T) {

		phone := []string{
			"1234567890",
			"01234567890", // เกินมา 1
			"+55123456789",
			"+66123",
			// "+66123456789", Truee
		}
		for _, val := range phone {

			user := User{
				Name:        "Aon",
				Email:       "Aon@mail.com",
				StudentID:   "B6310264",
				PhoneNumber: val,
				Password:    "123456789",
			}

			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("PhoneNumber invalid"))

		}

	})

	t.Run("check Password", func(t *testing.T) {

		pass := []string{
			"1234567",               // less than 8 characters
			"123456789012345678901", // greater than 20 characters
		}
		for _, val := range pass {

			user := User{
				Name:        "Aon",
				Email:       "Aon@mail.com",
				StudentID:   "B6310264",
				PhoneNumber: "0842298645",
				Password:    val,
			}

			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("password must be between 8 and 20 characters"))

		}

	})

}
