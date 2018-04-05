# graph-theory-year3

## Intro
My name is Ray Mannion and I am a 3rd student at the Galway/Mayo Institute of Technology studying Software Development

In this repository you will find a project I was asked to complete for a graph theory module. The project was completed using the Go programming language, which I am currently learning.

In order for you to run these programs you will need to install the [GO](https://golang.org/doc/install)
compiler on your computer and run the git clone command to pull the files from github.


## Run the code

Once you have installed the compiler, and have the files on your computer, open a command line of your choice. Next, navigate your way to the folder. Then run the command "go build" followed by the name of the go file(including the file extension) you wish to run.

This will then create an executeable file of the same name, but with an ".exe" ententsion. Then simpily run the .exe file and the GO program will run.

![alt tag](https://github.com/rayman51/graph-theory-year3/blob/master/images/Capture.PNG?raw=true)

## Project Overview

You must write a program in the Go programming language [2] that can build a non-deterministic finite automaton (NFA) from a regular expression,and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library. A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means any number of 1’s. These special characters must be used in your submission.

 Other special characters you might consider allowing as input are brackets “()” which can be used for grouping, “+” which means “at least one of”, and “?” which means “zero or one of”. You might also decide to remove the concatenation character, so that 1.0 becomes 10, with the concatenation implicit. You may initially restrict the non-special characters your program works with to 0 and 1, if you wish. However, you should at least attempt to expand these to all of the digits, and the characters a to z, and A to Z. You are expected to be able to break this project into a number of smaller tasks that are easier to solve, and to plug these together after they have been completed. You might do that for this project as follows:

1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.

## Helpful Websites
[Shunting Yard Algorithm](https://brilliant.org/wiki/shunting-yard-algorithm/) 

[Deterministic Finite Automata Example](https://www.tutorialspoint.com/automata_theory/deterministic_finite_automaton.htm)

[Non-Deterministic Finite Automata NFA Example](https://www.tutorialspoint.com/automata_theory/non_deterministic_finite_automaton.htm)

[Thompson's Regular Expression Explaination](https://swtch.com/~rsc/regexp/regexp1.html)

[Go Lanuage](https://golang.org/)