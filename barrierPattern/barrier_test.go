package barrierPattern

import "testing"

func TestBarrier(t *testing.T) {
	t.Run("Correct end-points", func(t *testing.T) {
		endPoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
	})

	t.Run("One end-point incorrect", func(t *testing.T) {
		endPoints := []string{"http://Malformem/ss/headers", "http://httpbin.org/User-Agent"}
	})

	t.Run("Very short timeout end-points", func(t *testing.T) {
		endPoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
	})
}

func barrier(endpoints ...string) {

}