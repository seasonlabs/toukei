package commands

import (
	   "os/exec"
    	"bytes"
    	"strconv"
    	"strings"
)

func sanitize(out []byte) (count int) {
        count, err := strconv.Atoi(strings.TrimSpace(string(out)))
        if err != nil {
                count = 0
        }
	return
}

func CountLines(path string) (int, error) {
	gls := exec.Command("git", "ls-files")
        gls.Dir = path
    	cat := exec.Command("xargs", "cat")
        cat.Dir = path
    	wc := exec.Command("wc", "-l")
        wc.Dir = path

    	out, _, _ := Pipeline(gls, cat, wc)

    	count := sanitize(out)
        
    	return count, nil
}

func CountCommits(path string) (int, error) {
	gls := exec.Command("git", "log", "--pretty=oneline")
        gls.Dir = path
    	wc := exec.Command("wc", "-l")
        wc.Dir = path

    	out, _, _ := Pipeline(gls, wc)

    	count := sanitize(out)
    	
    	return count, nil
}

// To provide input to the pipeline, assign an io.Reader to the first's Stdin.
func Pipeline(cmds ...*exec.Cmd) (pipeLineOutput, collectedStandardError []byte, pipeLineError error) {
        // Require at least one command
        if len(cmds) < 1 { 
                return nil, nil, nil
        }

        // Collect the output from the command(s)
        var output bytes.Buffer
        var stderr bytes.Buffer

        last := len(cmds) - 1
        for i, cmd := range cmds[:last] {
                var err error
                // Connect each command's stdin to the previous command's stdout
                if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
                        return nil, nil, err
                }
                // Connect each command's stderr to a buffer
                cmd.Stderr = &stderr
        }

        // Connect the output and error for the last command
        cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

        // Start each command
        for _, cmd := range cmds {
                if err := cmd.Start(); err != nil {
                        return output.Bytes(), stderr.Bytes(), err
                }
        }

        // Wait for each command to complete
        for _, cmd := range cmds {
                if err := cmd.Wait(); err != nil {
                        return output.Bytes(), stderr.Bytes(), err
                }
        }

        // Return the pipeline output and the collected standard error
        return output.Bytes(), stderr.Bytes(), nil
}
