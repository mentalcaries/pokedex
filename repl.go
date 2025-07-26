package main

import "strings"

func cleanInput(text string) []string {

    lowerCaseString := strings.ToLower(text)
    cleanStrings := strings.Fields(lowerCaseString)
    

    return cleanStrings
}