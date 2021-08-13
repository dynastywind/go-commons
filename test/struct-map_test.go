package structs_test

import (
	"github.com/dynastywind/go-commons/structs"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Struct-Map function tests", func() {
	ginkgo.Context("Struct to map test", func() {
		ginkgo.When("Converting int field", func() {
			ginkgo.It("Should return a map with an int", func() {
				type A struct {
					A int
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": 1}))
			})
			ginkgo.It("Should return a map with an int8", func() {
				type A struct {
					A int8
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": int8(1)}))
			})
			ginkgo.It("Should return a map with an int16", func() {
				type A struct {
					A int16
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": int16(1)}))
			})
			ginkgo.It("Should return a map with an int32", func() {
				type A struct {
					A int32
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": int32(1)}))
			})
			ginkgo.It("Should return a map with an int64", func() {
				type A struct {
					A int64
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": int64(1)}))
			})
		})
		ginkgo.When("Converting uint field", func() {
			ginkgo.It("Should return a map with a uint", func() {
				type A struct {
					A uint
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uint(1)}))
			})
			ginkgo.It("Should return a map with a uint8", func() {
				type A struct {
					A uint8
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uint8(1)}))
			})
			ginkgo.It("Should return a map with a uint16", func() {
				type A struct {
					A uint16
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uint16(1)}))
			})
			ginkgo.It("Should return a map with a uint32", func() {
				type A struct {
					A uint32
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uint32(1)}))
			})
			ginkgo.It("Should return a map with a uint64", func() {
				type A struct {
					A uint64
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uint64(1)}))
			})
		})
		ginkgo.When("Converting float field", func() {
			ginkgo.It("Should return a map with a float32", func() {
				type A struct {
					A float32
				}
				m := structs.Map(A{A: 1.0})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": float32(1.0)}))
			})
			ginkgo.It("Should return a map with a float64", func() {
				type A struct {
					A float64
				}
				m := structs.Map(A{A: 1.0})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": 1.0}))
			})
		})
		ginkgo.When("Converting complex field", func() {
			ginkgo.It("Should return a map with a complex64", func() {
				type A struct {
					A complex64
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": complex64(1)}))
			})
			ginkgo.It("Should return a map with a complex128", func() {
				type A struct {
					A complex128
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": complex128(1)}))
			})
		})
		ginkgo.When("Converting bool field", func() {
			ginkgo.It("Should return a map with a bool", func() {
				type A struct {
					A bool
				}
				m := structs.Map(A{A: true})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": true}))
			})
		})
		ginkgo.When("Converting string field", func() {
			ginkgo.It("Should return a map with a string", func() {
				type A struct {
					A string
				}
				m := structs.Map(A{A: "s"})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": "s"}))
			})
		})
		ginkgo.When("Converting array field", func() {
			ginkgo.It("Should return a map with an array", func() {
				type A struct {
					A []int
				}
				m := structs.Map(A{A: []int{1}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": []int{1}}))
			})
			ginkgo.It("Should return a map with an array of complex type", func() {
				type V struct {
					V int
				}
				type A struct {
					A []V
				}
				m := structs.Map(A{A: []V{{V: 1}}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": []V{{V: 1}}}))
			})
		})
		ginkgo.When("Converting map field", func() {
			ginkgo.It("Should return a map with a map", func() {
				type A struct {
					A map[string]int
				}
				m := structs.Map(A{A: map[string]int{"a": 1}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": map[string]int{"a": 1}}))
			})
			ginkgo.It("Should return a map with a map of complex value type", func() {
				type V struct {
					V int
				}
				type A struct {
					A map[string]V
				}
				m := structs.Map(A{A: map[string]V{"a": {V: 1}}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": map[string]V{"a": {V: 1}}}))
			})
		})
		ginkgo.When("Converting ptr field", func() {
			ginkgo.It("Should return a map with a ptr value", func() {
				type A struct {
					A *bool
				}
				b := true
				m := structs.Map(A{A: &b})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": true}))
			})
			ginkgo.It("Should return a map with a double-ptr value", func() {
				type A struct {
					A **bool
				}
				b := true
				c := &b
				m := structs.Map(A{A: &c})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": true}))
			})
		})
		ginkgo.When("Converting uintptr field", func() {
			ginkgo.It("Should return a map with a uintptr", func() {
				type A struct {
					A uintptr
				}
				m := structs.Map(A{A: 1})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"A": uintptr(1)}))
			})
		})
		ginkgo.When("Converting struct field", func() {
			ginkgo.It("Should return a map with a struct", func() {
				type A struct {
					A string
				}
				type B struct {
					B A
				}
				m := structs.Map(B{B: A{A: "a"}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"B": map[string]interface{}{"A": "a"}}))
			})
		})
		ginkgo.When("Converting a tag-named struct field", func() {
			ginkgo.It("Should return a map with a struct of tagged field name", func() {
				type A struct {
					A string `structs:"custom"`
				}
				type B struct {
					B A `structs:"custom"`
				}
				m := structs.Map(B{B: A{A: "a"}})
				gomega.Expect(m).To(gomega.Equal(map[string]interface{}{"custom": map[string]interface{}{"custom": "a"}}))
			})
		})
	})
})
