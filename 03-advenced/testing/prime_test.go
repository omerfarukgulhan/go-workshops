package prime

import "testing"

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for _, prime := range primes {
		if !IsPrime(prime) {
			t.Errorf("%d should be prime, but it's not", prime)
		}
	}
}
func TestIsNonPrime(t *testing.T) {
	nonPrimes := []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16}
	for _, nonPrime := range nonPrimes {
		if IsPrime(nonPrime) {
			t.Errorf("%d should not be prime, but it is", nonPrime)
		}
	}
}
