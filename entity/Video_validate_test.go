package entity

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestVideo(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		video := Video{
			Name: "",
			Url:  "https://web.facebook.com/groups/577615270797036",
		}
		ok, err := govalidator.ValidateStruct(video)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))

	})

	t.Run("check name not blank", func(t *testing.T) {
		video := Video{
			Name: "Test",
			Url:  "https://web.facebook.com/groups/577615270797036",
		}
		ok, err := govalidator.ValidateStruct(video)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		// g.Expect(err.Error()).To(Equal("Name cannot be blank"))

	})

	t.Run("check name not blank", func(t *testing.T) {
		video := Video{
			Name: "Tweast",
			Url:  "hookco97036",
		}
		ok, err := govalidator.ValidateStruct(video)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Url: %s does not validate as url", video.Url)))

	})

}
