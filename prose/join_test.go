package proseimport

import (
	"fmt"
	"testing"
)

func TestZeroElement(t *testing.T) {
	// Arrange
	//
	//
	list := []string{}
	want := ""
	// Act
	got := JoinWithCommas(list)
	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestOneElement(t *testing.T) {
	// Arrange
	//
	list := []string{"apple"}
	want := "apple"
	// Act
	//
	got := JoinWithCommas(list)
	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	// Arrange
	//
	list := []string{"apple", "orange"}
	want := "apple and orange"
	// Act
	//

	got := JoinWithCommas(list)
	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}
func TestThreeElements(t *testing.T) {
	// Arrange
	//
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	// Act
	//
	got := JoinWithCommas(list)
	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}
func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
