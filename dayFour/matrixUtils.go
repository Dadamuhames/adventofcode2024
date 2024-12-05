package main



func rotateMarix(matrix [][]string) [][]string {
    matrixLen := len(matrix);
        
    if matrixLen == 0 {
        return make([][]string, 0);
    }

    result := make([][]string, len(matrix[0]));

    for l := 0; l < matrixLen; l++ {
        result[l] = make([]string, matrixLen);
    }

    for i := 0; i < matrixLen; i++ {
        rowNum := matrixLen - 1 - i;

        for j := 0; j < len(matrix[i]); j++ {
            result[j][rowNum] = matrix[i][j];
        } 
    }

    return result;
}
