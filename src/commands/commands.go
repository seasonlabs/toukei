package commands

import (
	"os/exec"
    	"strconv"
    	"strings"
        "errors"
)

func sanitize(out []byte) (count int) {
        count, err := strconv.Atoi(strings.TrimSpace(string(out)))
        if err != nil {
                count = 0
        }
	return
}

func CountFiles(path string) (int, error) {
        gls := exec.Command("git", "ls-files")
        gls.Dir = path

    	out, _ := gls.Output()

    	if (string(out) == "") {
    		return 0, errors.New("Not a git repository")
    	}

        ss := strings.Split(string(out), "\n")
        
    	return len(ss), nil
}

func CountCommits(path string) (int, error) {
        gls := exec.Command("git", "log", "--pretty=oneline")
        gls.Dir = path
    	
        out, _ := gls.Output()
        if (string(out) == "") {
    		return 0, errors.New("Not a git repository")
    	}
	
	ss := strings.Split(string(out), "\n")

	return len(ss), nil
}
