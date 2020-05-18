package tools

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bag", func() {
	It("NewBag should return an instance of Bag", func() {
		Expect(NewBag()).NotTo(BeNil())
	})
})

var _ = Describe("factoryBag", func() {
	It("Get should return a value", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_key"] = "my_value"
		Expect(b.Get("my_key")).To(Equal("my_value"))
	})

	It("Get should return a default value if key does not exist", func() {
		b := &factoryBag{make(map[string]interface{})}
		v := b.GetOrDefault("my_key", "my_value")
		Expect(v).To(Equal("my_value"))
	})

	It("Has should return a boolean", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_key"] = "my_value"
		Expect(b.Has("my_key")).To(BeTrue())
		Expect(b.Has("my_another_key")).To(BeFalse())
	})

	It("Set should allow to set value", func() {
		b := &factoryBag{make(map[string]interface{})}
		Expect(b.Has("my_key")).To(BeFalse())
		b.Set("my_key", "my_value")
		Expect(b.Has("my_key")).To(BeTrue())
	})

	It("Remove should allow to remove a key", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_key"] = "my_value"
		Expect(b.Has("my_key")).To(BeTrue())
		b.Remove("my_key")
		Expect(b.Has("my_key")).To(BeFalse())
	})

	It("All should return all items", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_key"] = "my_value"
		b.data["my_another_key"] = 1
		Expect(len(b.All())).To(Equal(2))
	})

	It("GetInt should return int64 value", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_int64"] = 10
		b.data["my_float64"] = float64(12)
		Expect(b.GetInt("my_int64")).To(Equal(int64(10)))
		Expect(b.GetInt("my_another_int64")).To(Equal(int64(0)))
		Expect(b.GetInt("my_float64")).To(Equal(int64(12)))
	})

	It("GetFloat should return float64 value", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_float64"] = 10.01
		Expect(b.GetFloat("my_float64")).To(Equal(float64(10.01)))
		Expect(b.GetFloat("my_another_float64")).To(Equal(float64(0.0)))
	})

	It("GetString should return string value", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_string"] = "10.01"
		Expect(b.GetString("my_string")).To(Equal("10.01"))
		Expect(b.GetString("my_another_string")).To(BeEmpty())

		b.data["my_another_string"] = 12
		Expect(b.GetString("my_another_string")).To(Equal("12"))
	})

	It("GetBool should return boolean value", func() {
		b := &factoryBag{make(map[string]interface{})}
		b.data["my_bool"] = true
		Expect(b.GetBool("my_bool")).To(BeTrue())
		Expect(b.GetBool("my_another_bool")).To(BeFalse())

		b.data["my_another_bool"] = "true"
		Expect(b.GetBool("my_another_bool")).To(BeTrue())

		b.data["my_another_bool"] = "1"
		Expect(b.GetBool("my_another_bool")).To(BeTrue())

		b.data["my_bool"] = "false"
		Expect(b.GetBool("my_bool")).To(BeFalse())

		b.data["my_bool"] = "0"
		Expect(b.GetBool("my_bool")).To(BeFalse())
	})
})
