package commands

import(
	"testing"
)

func TestSanitize(t *testing.T) {
	foo := []byte("   12345   ")

	count := sanitize(foo)
	if count == 0 {
		t.Error("Error count")
	}
}

func TestCountCommits(t *testing.T) {
	foo, _ := CountCommits(".")

	if foo == 0 {
		t.Fail()
	} else {
		t.Log(foo)
	}
}

func TestCountLinesGitDir(t *testing.T) {
	foo, err := CountLines(".")

	if err != nil {
		t.Fail()
	}

	if foo == 0 {
		t.Fail()
	} else {
		t.Log(foo)
	}
}

func TestCountLinesNoGitDir(t *testing.T) {
	foo, _ := CountLines("/tmp")

	if foo != 0 {
		t.Fail()
	} else {
		t.Log(foo)
	}
}