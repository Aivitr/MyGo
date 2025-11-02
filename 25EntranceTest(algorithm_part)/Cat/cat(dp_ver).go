package main


import "fmt"


func main() {
    
    var n, E, maxE, r, cnt int

    fmt.Scanln(&n, &E, &r)

    dp := make([][]int, n)
    a := make([]int, n)

    var str string
    
    fmt.Scanf("%s", &str)

    for i := 0; i < n; i++ {

        if str[i] == '0' {a[i] = 0}

        if str[i] == '-' {a[i] = 1}

        if str[i] == '+' {a[i] = 2;cnt++}

    }

    maxE = cnt*r + E


    for i := 0; i < n; i++ {

        dp[i] = make([]int, maxE+1)

        for j := 1; j <= maxE; j++ {dp[i][j] = 2147483647}

    }


    dp[0][E] = 0


    for i := 0; i < n; i++ {
        for j := 1; j <= maxE; j++ {
            if dp[i][j] != 2147483647 {
                for k := 1; k <= j && i+k < n; k++ {

                    if a[i+k] == 0 {dp[i+k][j] = min(dp[i+k][j], dp[i][j]+1)}

                    if a[i+k] == 2 {dp[i+k][j+r] = min(dp[i+k][j+r], dp[i][j]+1)}

                }
            }
        }
    }


    minJump := n

    for i := 1; i <= maxE; i++ {minJump = min(minJump, dp[n-1][i])}

    fmt.Println(minJump)
    
}