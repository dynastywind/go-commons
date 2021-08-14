package either_test

import (
	"testing"

	"github.com/dynastywind/go-commons/either"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestStream(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Either Tests")
}

var _ = ginkgo.Describe("Either function tests", func() {
	ginkgo.Context("General either test", func() {
		ginkgo.When("Executing operations on left", func() {
			ginkgo.It("Should not be nil", func() {
				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("should panic when assigning a nil left to Either")
					}
				}()
				either.OfLeft(nil)
			})
			ginkgo.It("Should succeed", func() {
				e := either.OfLeft(1)
				gomega.Expect(e.HasLeft()).To(gomega.BeTrue())
				gomega.Expect(e.HasRight()).To(gomega.BeFalse())
				gomega.Expect(e.GetLeft()).To(gomega.Equal(1))
				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("Should panic when getting right")
					}
				}()
				e.GetRight()
				gomega.Expect(e.MapLeft(func(data interface{}) interface{} { return 2 }).Get()).To(gomega.Equal(2))
				gomega.Expect(e.MapRight(func(data interface{}) interface{} { return 2 }).IsPresent()).To(gomega.BeFalse())
				var a *int
				e.IfLeftPresent(func(data interface{}) {
					d := data.(*int)
					*d = *d + 1
				})
				gomega.Expect(*a).To(gomega.Equal(2))
				e.IfRightPresent(func(data interface{}) {
					d := data.(*int)
					*d = *d + 1
				})
				gomega.Expect(*a).To(gomega.Equal(2))
			})
		})
		ginkgo.When("Executing operations on right", func() {
			ginkgo.It("Should not be nil", func() {
				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("should panic when assigning a nil left to Either")
					}
				}()
				either.OfRight(nil)
			})
			ginkgo.It("Should succeed", func() {
				e := either.OfRight(1)
				gomega.Expect(e.HasLeft()).To(gomega.BeFalse())
				gomega.Expect(e.HasRight()).To(gomega.BeTrue())
				gomega.Expect(e.GetRight()).To(gomega.Equal(1))
				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("Should panic when getting right")
					}
				}()
				e.GetLeft()
				gomega.Expect(e.MapRight(func(data interface{}) interface{} { return 2 }).Get()).To(gomega.Equal(2))
				gomega.Expect(e.MapLeft(func(data interface{}) interface{} { return 2 }).IsPresent()).To(gomega.BeFalse())
				var a *int
				e.IfRightPresent(func(data interface{}) {
					d := data.(*int)
					*d = *d + 1
				})
				gomega.Expect(*a).To(gomega.Equal(2))
				e.IfLeftPresent(func(data interface{}) {
					d := data.(*int)
					*d = *d + 1
				})
				gomega.Expect(*a).To(gomega.Equal(2))
			})
		})
	})
})
