package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMinimumRecolors(t *testing.T) {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}
