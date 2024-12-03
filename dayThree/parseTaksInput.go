package main

import (
	"bufio"
	"os"
	"regexp"
)



func getTaskInput() string {
    file, err := os.Open("input.txt");

    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file);
    
    result := "";

    for scanner.Scan() {
        result += scanner.Text();
    }
    
    return result;
}



func parseValidExpressions() []string {
    var result []string = []string {};

    input := getTaskInput(); 

    r, _ := regexp.Compile(`mul\(\d*,\d*\)`);

    result = r.FindAllString(input, -1);

    return result;
}


func parseValidExpressionsWithInstructions() []string {
    var result []string = []string {};

    input := getTaskInput();

    r, _ := regexp.Compile(`mul\(\d*,\d*\)`);
    doR, _ := regexp.Compile(`do\(\)`);
    dontR, _ := regexp.Compile(`don't\(\)`);

    
    index := 0;
    enabled := true;

    for index < len(input) {
        s := input[index:]

        funcIndex := r.FindStringIndex(s);
               
        if len(funcIndex) == 0 { break; }

        dontIndex := dontR.FindStringIndex(s);

        if len(dontIndex) != 0 && dontIndex[1] <= funcIndex[0] {
            index += dontIndex[1];
            enabled = false;
            continue;
        }


        doIndex := doR.FindStringIndex(s);

        if len(doIndex) != 0 && doIndex[1] <= funcIndex[0] {
            enabled = true;
        }

        if(enabled) {
            result = append(result, r.FindString(s));
        }

        index += funcIndex[1];
    }

    return result;
}
