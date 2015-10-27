package devopsutil

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"strconv"
	"errors"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func Print(w http.ResponseWriter, msg string) {
	if w != nil {
		fmt.Fprintln(w, msg)
	} else {
		fmt.Println(msg)
	}
}

func ValidateFile(w http.ResponseWriter, fileName string) (err error) {
	var row int = 0
	var trailerCount int = 0
    var hdrDate string
	
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()
	
    fHandle, err := os.Open(fileName)
	Check(err)
	
	defer fHandle.Close()
	
	s := bufio.NewScanner(fHandle)
	
	for s.Scan() {
		line := s.Text()
		field := strings.Split(line, ",")

		if field[0] == "H" {
			hdrDate = field[1]
		} else if field[0] == "D" {
			row = row + 1
		} else if field[0] == "T" {
			trailerCount, err = strconv.Atoi(field[1]);
			Check(err)
		} else {
			err := errors.New("Invalid Record Type")
			Check(err)
		}
		
		Print(w, line)
	}
	Check(s.Err())
	
	Print(w, "=======================================================")
	Print(w, fmt.Sprint("File Received date = ", hdrDate))
	Print(w, fmt.Sprint("Total dtl rec      = ", row)) 
	Print(w, fmt.Sprint("Trailer rec cnt    = ", trailerCount))
	
	if row != trailerCount {
		err := errors.New("Trailer count mismatch")
		Check(err)
	}
	
	return
}

func OScommand (w http.ResponseWriter, cmd []string) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()
	
    out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	
	Check(err)
	
	Print(w, string(out[:]))	
	
	return
}