package sort

import (
	"github.com/alexgunkel/go_algorithms/dataprovider"
	"testing"
)

func TestQuickSort(t *testing.T) {
	dataprovider.SorterTest(t, QuickSort)
}
