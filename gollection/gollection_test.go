package gollection_test

import (
	g "gollection/gollection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gollection", func() {
	It("should do stuff", func(){
		count := 0
		g.New([]string{"a", "b"}).
			Filter(func(s string) bool {return s == "b"}).
			Each(func(s string){
				count++
			})
		Expect(count).To(Equal(1))
	})
})
