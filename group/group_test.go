package group

import (
	"strings"
	"testing"
)

const examplegroup = `# test comment
root:x:0:
ryan:x:501:
ryan2:x:502:
sudo:x:555:ryan,ryan2`

func TestGroup(t *testing.T) {
	r := strings.NewReader(examplegroup)
	data, err := ParseReader(r)
	if err != nil {
		t.Error(err)
	}
	if data["root"].GID != "0" {
		t.Error("data parsing failed")
	}
	if data["sudo"].GID != "555" {
		t.Error("data parsing failed")
	}
	if len(data["sudo"].Members) != 2 {
		t.Error("data parsing failed")
	}
	if data["sudo"].Members[0] != "ryan" {
		t.Error("data parsing failed")
	}
	if data["sudo"].Members[1] != "ryan2" {
		t.Error("data parsing failed")
	}
}
