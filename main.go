package main

// https://freshman.tech/golang-guess/

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    min, max := 1, 100

    // init random number seed
    rand.Seed( time.Now().UnixNano())

    // Intn() range is 0 to exclusive_max, so Intn(10) is in [0..10)
    secretNumber := rand.Intn(max - min + 1) + min
    // fmt.Println( "The secret number is", secretNumber )


    fmt.Println( "Guess a number between 1 and 100" )
    fmt.Println( "Please input your guess" )

    numAttempts := 0
    for {
        numAttempts++
        reader := bufio.NewReader( os.Stdin )
        input, err := reader.ReadString( '\n' )
        if err != nil {
            fmt.Println( "An error occurred while reading input; please try again", err )
            continue
        }

        input = strings.TrimSuffix( input, "\n" )

        guess, err := strconv.Atoi( input )
        if err != nil {
            fmt.Println( "Invalid input; please enter an INTEGER value" )
            continue
        }

        fmt.Println( "Your guess is", guess )

        if guess > secretNumber {
            fmt.Println( "Too HIGH; try again" )
        } else if guess < secretNumber {
            fmt.Println( "Too LOW; try again" )
        } else {
            guesses := "guesses"
            if numAttempts == 1 {
                guesses = "guess"
            }
            fmt.Println( "Correct -- you win! And it took only", numAttempts, guesses )
            break
        }
    }
}
