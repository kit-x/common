package syncx

import "testing"

func TestWaitGroupWrapper_Wrap(t *testing.T) {
	var (
		wg     WaitGroupWrapper
		target bool
	)

	wg.Wrap(func() {
		target = true
	})
	wg.Wait()

	if !target {
		t.Fatal("should handle callback")
	}
}
