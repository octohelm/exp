package xchan

import (
	"context"
	"math/rand"
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func FuzzNotifiableObserver(f *testing.F) {
	for range 10 {
		f.Add(rand.Intn(10000))
	}

	f.Fuzz(func(t *testing.T, n int) {
		x := NewNotifiableObserver[int]()

		for range 10 {
			go func() {
				<-x.Done()
			}()
		}

		go func() {
			for i := range n {
				x.Send(i)
			}

			x.CancelCause(nil)
		}()

		values := slices.Collect(Values(context.Background(), x))

		testingx.Expect(t, len(values), testingx.Be(n))

		_, doneChOk := <-x.Done()
		testingx.Expect(t, false, testingx.Be(doneChOk))

		_, valueChOk := <-x.Value()
		testingx.Expect(t, false, testingx.Be(valueChOk))
	})
}
