package devopsutil

import "testing" 

func TestDevOpsDemo(t *testing.T) { 
	err := ValidateFile(nil, "../../InputFile/Test.txt")
	
	if err != nil {
		panic(err)
	}
 } 
