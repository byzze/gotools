package gotool

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSortByVersion(t *testing.T) {
	var list = []string{
		"v1.1.1",
		"v1.1.3",
		"v1.1.4.5",
		"v1.1.4.4",
		"v2.1.4.4",
		"v3.1.4.6",
		"v3.2",
		"v3.20",
		"v3.10",
		"v3.10.1",
		"v1.10",
		"v4.10.1",
	}
	SortByVersion(list, Asc)
	logrus.Println(list)
}
