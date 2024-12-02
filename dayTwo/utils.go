package main



func isDiffValid(a int, b int, direction int) bool {
    diff := (a - b) * direction;

    return diff >= 1 && diff <= 3;
}


func getDirection(report []int) int {
    var trend int = 0;
    var direction int = 0;

    for i := 1; i < len(report); i++ {
        if report[i-1] > report[i] {
            trend--;
        } else if report[i-1] < report[i] {
            trend++;
        }
    }

    if trend > 0 {
        direction = 1;
    } else if trend < 0 {
        direction = -1;
    }


    return direction;
}


func remove(list []int, index int) []int {
    var sliceCopy []int =  make([]int, len(list));

    copy(sliceCopy, list);

    return append(sliceCopy[:index], sliceCopy[index+1:]...);
}


func isReportSafe(report []int, dampener bool) bool { 
    var direction int = getDirection(report);

    if (direction == 0) { return false };
 

    for i := 1; i < len(report); i++ {
        prev := report[i-1]

        diffValid := isDiffValid(report[i], prev, direction);

        if !diffValid { 
            if dampener {
                return isReportSafe(remove(report, i), false) || isReportSafe(remove(report, i-1), false);
            }

            return false;
        }
    }
    
    return true;
}

