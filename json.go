package log_viewer

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const indent = "    "

// JSON prettifies a JSON object passed in as a string.
// It does that by indenting each key and optionally indenting each line extra
func JSON(s string, indentEachLine ...bool) ([]byte, error) {
	var prefix string
	buf := bytes.NewBuffer(make([]byte, len(s))) // size not entirely accurate, but should help a little with the allocations

	if len(indentEachLine) == 1 && indentEachLine[0] {
		prefix = indent
		buf.WriteString(prefix)
	}

	if err := json.Indent(buf, []byte(s), prefix, indent); err != nil {
		return nil, fmt.Errorf("failed to indent json string: %w", err)
	}

	return buf.Bytes(), nil
}
