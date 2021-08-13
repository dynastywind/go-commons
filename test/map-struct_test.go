package structs_test

import (
	"reflect"

	"github.com/dynastywind/go-commons/structs"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Map-Struct function tests", func() {
	ginkgo.Context("Map to struct test", func() {
		ginkgo.When("Converting int field", func() {
			ginkgo.It("Should return a struct with an int", func() {
				type A struct {
					A int
				}
				s := structs.Struct(map[string]interface{}{"A": 1}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with an int8", func() {
				type A struct {
					A int8
				}
				s := structs.Struct(map[string]interface{}{"A": int8(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with an int16", func() {
				type A struct {
					A int16
				}
				s := structs.Struct(map[string]interface{}{"A": int16(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with an int32", func() {
				type A struct {
					A int32
				}
				s := structs.Struct(map[string]interface{}{"A": int32(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with an int64", func() {
				type A struct {
					A int64
				}
				s := structs.Struct(map[string]interface{}{"A": int64(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
		})
		ginkgo.When("Converting uint field", func() {
			ginkgo.It("Should return a struct with a uint", func() {
				type A struct {
					A uint
				}
				s := structs.Struct(map[string]interface{}{"A": uint(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with a uint8", func() {
				type A struct {
					A uint8
				}
				s := structs.Struct(map[string]interface{}{"A": uint8(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with a uint16", func() {
				type A struct {
					A uint16
				}
				s := structs.Struct(map[string]interface{}{"A": uint16(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with a uint32", func() {
				type A struct {
					A uint32
				}
				s := structs.Struct(map[string]interface{}{"A": uint32(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with a uint64", func() {
				type A struct {
					A uint64
				}
				s := structs.Struct(map[string]interface{}{"A": uint64(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
		})
		ginkgo.When("Converting float field", func() {
			ginkgo.It("Should return a struct with a float32", func() {
				type A struct {
					A float32
				}
				s := structs.Struct(map[string]interface{}{"A": float32(1.0)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1.0}))
			})
			ginkgo.It("Should return a struct with a float64", func() {
				type A struct {
					A float64
				}
				s := structs.Struct(map[string]interface{}{"A": 1.0}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1.0}))
			})
		})
		ginkgo.When("Converting complex field", func() {
			ginkgo.It("Should return a struct with a complex64", func() {
				type A struct {
					A complex64
				}
				s := structs.Struct(map[string]interface{}{"A": complex64(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
			ginkgo.It("Should return a struct with a complex128", func() {
				type A struct {
					A complex128
				}
				s := structs.Struct(map[string]interface{}{"A": complex128(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
		})
		ginkgo.When("Converting bool field", func() {
			ginkgo.It("Should return a struct with a bool", func() {
				type A struct {
					A bool
				}
				s := structs.Struct(map[string]interface{}{"A": true}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: true}))
			})
		})
		ginkgo.When("Converting string field", func() {
			ginkgo.It("Should return a struct with a string", func() {
				type A struct {
					A string
				}
				s := structs.Struct(map[string]interface{}{"A": "s"}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: "s"}))
			})
		})
		ginkgo.When("Converting array field", func() {
			ginkgo.It("Should return a struct with an array", func() {
				type A struct {
					A []int
				}
				s := structs.Struct(map[string]interface{}{"A": []int{1}}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: []int{1}}))
			})
			ginkgo.It("Should return a struct with an array of complex type", func() {
				type V struct {
					V int
				}
				type A struct {
					A []V
				}
				s := structs.Struct(map[string]interface{}{"A": []V{{V: 1}}}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: []V{{V: 1}}}))
			})
		})
		ginkgo.When("Converting map field", func() {
			ginkgo.It("Should return a struct with a map", func() {
				type A struct {
					A map[string]int
				}
				s := structs.Struct(map[string]interface{}{"A": map[string]int{"a": 1}}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: map[string]int{"a": 1}}))
			})
			ginkgo.It("Should return a struct with a map of complex value type", func() {
				type V struct {
					V int
				}
				type A struct {
					A map[string]V
				}
				s := structs.Struct(map[string]interface{}{"A": map[string]V{"a": {V: 1}}}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: map[string]V{"a": {V: 1}}}))
			})
		})
		ginkgo.When("Converting ptr field", func() {
			ginkgo.It("Should return a struct with a ptr value", func() {
				type A struct {
					A *bool
				}
				b := true
				s := structs.Struct(map[string]interface{}{"A": true}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: &b}))
			})
			ginkgo.It("Should return a struct with a double-ptr value", func() {
				type A struct {
					A **bool
				}
				b := true
				c := &b
				s := structs.Struct(map[string]interface{}{"A": true}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: &c}))
			})
		})
		ginkgo.When("Converting uintptr field", func() {
			ginkgo.It("Should return a struct with a uintptr", func() {
				type A struct {
					A uintptr
				}
				s := structs.Struct(map[string]interface{}{"A": uintptr(1)}, reflect.TypeOf(A{}))
				gomega.Expect(s).To(gomega.Equal(A{A: 1}))
			})
		})
		ginkgo.When("Converting struct field", func() {
			ginkgo.It("Should return a struct with a struct", func() {
				type A struct {
					A string
				}
				type B struct {
					B A
				}
				s := structs.Struct(map[string]interface{}{"B": map[string]interface{}{"A": "s"}}, reflect.TypeOf(B{}))
				gomega.Expect(s).To(gomega.Equal(B{B: A{A: "s"}}))
			})
		})
		ginkgo.When("Converting a tag-named field", func() {
			ginkgo.It("Should return a struct with a tagged field name", func() {
				type A struct {
					A string `structs:"custom"`
				}
				type B struct {
					B A `structs:"custom"`
				}
				s := structs.Struct(map[string]interface{}{"custom": map[string]interface{}{"custom": "s"}}, reflect.TypeOf(B{}))
				gomega.Expect(s).To(gomega.Equal(B{B: A{A: "s"}}))
			})
		})
	})
})
