package numberParser

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func FuzzAppendNonEmpty(f *testing.F) {
	f.Add("hello")
	f.Add("something")
	f.Add("")
	f.Fuzz(func(t *testing.T, s string) {
		var testSlice = appendNonEmpty(make([]string, 0), s)
		if s == "" {
			if len(testSlice) != 0 {
				t.Fatal("Empty string appended to slice!")
			}
		} else if testSlice[0] != s {
			t.Fatal("Element not appended!")
		}
	})
}

func TestPrepend(t *testing.T) {
	var testSlice = make([]int, 0, 1)
	for i := 0; i < 10; i++ {
		testSlice = prepend(testSlice, i)
		if testSlice[0] != i {
			t.Fatalf("Fail to prepend element %d. Its not in first position!", i)
		}
	}
}

func FuzzGroupHundreds(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(5))
	f.Add(uint64(10))
	f.Add(uint64(105))
	f.Add(uint64(1623))
	f.Add(uint64(123123123123))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u uint64) {
		var groups = groupHundreds(u)
		var expected = strconv.FormatUint(u, 10)
		var returned = strings.Join(groups, "")
		if returned != expected {
			t.Errorf("Numbers do not match!\nExpected: %s. Received: %s.", expected, returned)
		}
		if expectedLength, receivedLength := len(expected)%3, len(groups[0]); expectedLength != receivedLength && !(expectedLength == 0 && receivedLength == 3) {
			t.Errorf(
				"The first group doesn't have the expected length!\nExpected lenght: %d, Received: %d. Group: %s should be %s.",
				expectedLength, receivedLength, groups[0], expected[:expectedLength],
			)
		}
		if len(expected) > 3 && len(groups[1]) != 3 {
			t.Fatalf("Every group except the first should have length 3!\nLength: %d, Group: %s.", len(groups[1]), groups[1])
		}
	})
}

func TestTranslateGroup(t *testing.T) {
	for zeroes := "0"; len(zeroes) <= 3; zeroes += "0" {
		if result := translateGroup(zeroes); result != "" {
			t.Errorf("Zero numbers should return an empty string, but returned %q instead!", result)
		}
	}
	var testcases = [...]struct {
		in, expected string
	}{
		{"5", "cinco"},
		{"10", "dez"},
		{"15", "quinze"},
		{"20", "vinte"},
		{"100", "cem"},
		{"105", "cem e cinco"},
		{"215", "duzentos e quinze"},
		{"525", "quinhentos e vinte e cinco"},
	}
	for _, test := range testcases {
		var result = translateGroup(test.in)
		if result != test.expected {
			t.Errorf(
				"Group not translated as expected!\nGroup: %s, expected: %q, returned: %q",
				test.in, test.expected, result,
			)
		}
	}
}

func FuzzGetNumberInFull(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(10))
	f.Add(uint64(138))
	f.Add(uint64(1627))
	f.Add(uint64(3877672))
	f.Add(uint64(6565236473))
	f.Add(uint64(7364762377839))
	f.Add(uint64(9312676762736284))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u uint64) {
		var returned = GetNumberInFull(u)
		if u == 0 {
			if returned != ZERO {
				t.Fatalf("Number 0 should be %q, but received %q instead!", ZERO, returned)
			}
			return
		}
		var index = 1
		for period := uint64(1000); period <= u; period *= 1000 {
			if getPeriod(u, period) > 0 && !strings.Contains(returned, groupNames[index]) {
				if !strings.Contains(returned, plural(groupNames[index])) {
					t.Errorf(
						"Number %d should contain %q, but does not!\nReceived: %q",
						u, groupNames[index], returned,
					)
				}
			}
			if period == uint64(1000000000000000000) {
				break
			}
			index++
		}
	})
}

func getPeriod(number, period uint64) int {
	return int(number/period) % 1000
}
