# GraphTheoryProject2018
Darragh King, Student Id G00326176

Write a program in Go that can build Non-Deterministic Finite Automatas (NFAs) from Regular Expressions, that can use the NFA to check if the regular expression matches a string.

This project uses two algorithms; Thompsons Construction and the Shunting Yard Algorithm. 

Thompson's construction is an algorithm for creating NFAs from regular expressions. It uses two structs (state and nfa), that represent the NFA. As the program reads your regular expression, it updates your NFA with the new parameters. 

Once the NFA is finished, the program tests each character from the input string. This loops through the input string and states of the NFA to check if they match.

The Shunting Yard Algorithm is used to convert regular expressions from infix to postfix notation. This is needed so that the program understands the regular expressions without parentheses.


# How to run
You must install Go if you have not already to run this program on your machine.
Downloaded this repository onto your machine. 
Open the command prompt and navigate to the program folder (navigate to folder first, highlight the address bar then type "cmd" and press enter). 
Type "go run main.go" and press enter.
Enter a regular expression in infix notation. 
Enter a string to compare to your regular expression. 
The program will then tell you if the string matches.

# References

https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
http://www.cs.man.ac.uk/~pjj/cs212/fix.html
https://golang.org/pkg/regexp/


