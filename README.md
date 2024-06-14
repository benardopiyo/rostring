# rostring

* This program takes a single string as an argument and displays the string after rotating it one word to the left. 
* This means that the first word in the string becomes the last, and all other words remain in the same order.


## Requirements

* A word is defined as a sequence of alphanumerical characters.
* Words will be separated by only one space in the output.
* If the number of arguments is different from 1, the program displays a newline (\n).

## Usage

* To run the program, use the following commands in your terminal:

```bash
$ go run . "abc   " | cat -e
abc$
```
```bash
$ go run . "Let there 	be light"
there be light Let
```
```bash
$ go run . " 	AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
```
```bash
$ go run . | cat -e
$
```
