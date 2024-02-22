// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

import (
        "fmt"
        "os"
)

const (
        usage    = "Usage: [username] [password]"
        errUser  = "Access denied for %q.\n"
        errPwd   = "Invalid password for %q.\n"
        accessOK = "Access granted for %q.\n"

        user     = "jack"
        pwd      = "1888"
)

func main() {
        args := os.Args

        if len(args) != 3 {
                fmt.Println(usage)
                return
        }

        u, p := args[1], args[2]

        if u != user {
                fmt.Printf(errUser, u)
        } else if p != pwd {
                fmt.Printf(errPwd, u)
        } else {
                fmt.Printf(accessOK, u)
        }
}
