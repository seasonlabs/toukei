package commands

import (
	"os/exec"
    	"strconv"
    	"strings"
        "regexp"
        "errors"
)

func sanitize(out []byte) (count int) {
        count, err := strconv.Atoi(strings.TrimSpace(string(out)))
        if err != nil {
                count = 0
        }
	return
}

func CountLines(path string) (int, error) {
        gls := exec.Command("git", "log", "--shortstat", "--pretty=oneline")
        gls.Dir = path

    	out, _ := gls.Output()

    	if (string(out) == "") {
    		return 0, errors.New("Not a git repository")
    	}

        ss := strings.Split(string(out), "\n")

        totalFiles := 0
        totalInsertions := 0
        totalDeletions := 0
        totalLines := 0
        for _, line := range ss {
        	if line == "" || line[0] !=  ' ' {
        		continue
        	} else {
        		re := regexp.MustCompile(" (\\d+) files changed, (\\d+) insertions\\(\\+\\), (\\d+) deletions\\(-\\)")
        		data := re.FindStringSubmatch(line)
        		//fmt.Println("k")
        		
        		files, insertions, deletions := sanitize([]byte(data[1])), sanitize([]byte(data[2])), sanitize([]byte(data[3]))
        		
        		totalFiles += files
        		totalInsertions += insertions
        		totalDeletions += deletions

        		totalLines = totalInsertions - totalDeletions
        	}
        }
        
    	return totalLines, nil
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
