package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := sub(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestMul(t *testing.T) {
	result := mul(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result := div(6, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestMod(t *testing.T) {
	result := mod(7, 3)
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestInc(t *testing.T) {
	result := inc(5)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestDec(t *testing.T) {
	result := dec(5)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestIsEven(t *testing.T) {
	result := isEven(4)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	result = isEven(5)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestIsOdd(t *testing.T) {
	result := isOdd(5)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	result = isOdd(4)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestIsPrime(t *testing.T) {
	primeNumbers := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	}

	for _, num := range primeNumbers {
		if !isPrime(num) {
			t.Errorf("%d is expected to be prime, but it is not", num)
		}
	}

	nonPrimeNumbers := []int{
		1, 4, 6, 8, 9, 10, 12, 14, 15, 16,
		18, 20, 21, 22, 25, 26, 27, 28, 30, 32,
		33, 34, 35, 38, 39, 40, 44, 45, 46, 48,
		49, 50, 51, 52, 55, 56, 57, 58, 62, 63,
		64, 65, 66, 68, 69, 70, 74, 75, 76, 77,
		78, 80, 81, 82, 85, 86, 87, 88, 91, 92,
		93, 94, 95, 96, 98, 99, 100, 102, 104, 105,
		106, 110, 111, 112, 114, 115, 116, 118, 119, 120,
		121, 122, 123, 124, 125, 126, 128, 129, 130, 132,
		133, 134, 135, 136, 138, 140, 141, 142, 143, 144,
		145, 146, 147, 148, 150, 152, 153, 154, 155, 156,
		158, 159, 160, 161, 162, 164, 165, 166, 168, 169,
		170, 171, 172, 174, 175, 176, 177, 178, 180, 182,
		183, 184, 185, 186, 187, 188, 189, 190, 192, 194,
		195, 196, 198, 200, 201, 202, 203, 204, 205, 206,
		207, 208, 209, 210, 212, 213, 214, 215, 216, 217,
		218, 219, 220, 221, 222, 224, 225, 226, 228, 230}

	for _, num := range nonPrimeNumbers {
		if isPrime(num) {
			t.Errorf("%d is expected to be non-prime, but it is not", num)
		}
	}
}
