package util_test

import (
	"github.com/hstreamdb/http-server/pkg/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintResult(t *testing.T) {
	header := []string{"col1", "col2", "col3"}
	datas := [][]string{{"a", "acv", "bb"}, {"dd", "efghs", "b"}}
	expect := `
+------+-------+------+
| col1 | col2  | col3 |
+------+-------+------+
| a    | acv   | bb   |
| dd   | efghs | b    |
+------+-------+------+
`
	result, err := util.RenderTable(header, datas)
	require.NoError(t, err)
	require.Equal(t, expect[1:], result)
}
