package parsing

import (
	"testing"
)

func TestRemoveComments_test(t *testing.T) {
	input := `0
#comment
##start
1 3 5
1 1 0
1 1 8
3 1 5
##end
`

	expected := `0
##start
1 3 5
1 1 0
1 1 8
3 1 5
##end
`

	result := RemoveComments(input)

	if result != expected {
		t.Errorf("\nGot:\n%s\nExpected:\n%s", result, expected)
	}

}
