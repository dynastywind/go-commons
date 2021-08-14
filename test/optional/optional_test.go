package optional_test

import (
	"errors"
	"testing"

	"github.com/dynastywind/go-commons/optional"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestStream(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Optional Tests")
}

var _ = ginkgo.Describe("Optional function tests", func() {
	ginkgo.Context("General optional test", func() {
		ginkgo.When("Executing operations on real object", func() {
			ginkgo.It("Should succeed", func() {
				o := optional.Of(1)
				gomega.Expect(o.IsPresent()).To(gomega.BeTrue())
				gomega.Expect(o.Get()).To(gomega.Equal(1))
				gomega.Expect(o.Map(func(data interface{}) interface{} {
					return data.(int) + 1
				}).Get()).To(gomega.Equal(2))
				gomega.Expect(o.FlatMap(func(data interface{}) optional.Optional {
					return optional.Of(data)
				}).Get()).To(gomega.Equal(1))
				gomega.Expect(o.Filter(func(data interface{}) bool {
					return false
				}).IsPresent()).To(gomega.BeFalse())

				m, _ := o.OrElseError(func() error {
					return errors.New("error")
				})
				gomega.Expect(m).To(gomega.Equal(1))

				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("Should panic due to creating optional via Of method with nil value")
					}
				}()
				optional.Of(nil)
				gomega.Expect(optional.OfNillable(1).IsPresent()).To(gomega.BeTrue())
				gomega.Expect(optional.OfNillable(nil).IsPresent()).To(gomega.BeFalse())
			})
		})
		ginkgo.When("Executing operations on empty object", func() {
			ginkgo.It("Should succeed", func() {
				o := optional.OfEmpty()
				gomega.Expect(o.IsPresent()).To(gomega.BeFalse())
				defer func() {
					if r := recover(); r == nil {
						ginkgo.Fail("Should fail due to getting empty result")
					}
				}()
				o.Get()
				gomega.Expect(o.Map(func(data interface{}) interface{} {
					return data.(int) + 1
				}).IsPresent()).To(gomega.BeFalse())
				gomega.Expect(o.FlatMap(func(data interface{}) optional.Optional {
					return optional.Of(data)
				}).IsPresent()).To(gomega.BeFalse())
				gomega.Expect(o.Filter(func(data interface{}) bool {
					return true
				}).IsPresent()).To(gomega.BeFalse())
				gomega.Expect(o.OrElse(1)).To(gomega.Equal(1))
				gomega.Expect(o.OrElseGet(func() interface{} {
					return 1
				})).To(gomega.Equal(1))

				_, e := o.OrElseError(func() error {
					return errors.New("error")
				})
				gomega.Expect(e).NotTo(gomega.BeNil())
			})
		})
	})
})
