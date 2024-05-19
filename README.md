# GO-RELOADED

Go reloaded is a tool that recieves as arguments the name that contaians a text which needs some modification(input) and the name of the file the modified text should be placed in the ouput.


Some of the modifications include:

### Hex 
For every instance of (hex) should be replace the previous word from hexadecimal to decimal i.e 

   input: 1E (hex) files were added.

   output: 30 files were added.

### Bin
For every instance of (bin) should replace the previous word from binary to decimal i.e 

    input: It has been 10 (bin) years
    output: It has been 2 years

### Up
For every instance of (up) the previous word(words) should be replace to uppercase i.e

    input: Ready, set, go (up) !
    output: Ready, set, GO !

Here is another instance where the same can apply.
    input: Ready, set, go (up, 2) !
    output: Ready, SET, GO !

Here the number that follows dictates the number of words to act upon with.

### Low
For every instance of (low) the previous word(words) should be replace to lowercase i.e

    input: Ready, set, GO (low) !
    output: Ready, set, go !

Here is another instance where the same can apply.
    input: Ready, SET, GO (low, 2) !
    output: Ready, set, go !

 

### Cap
For every instance of (cap) the first letter of the previous word(words) should be replace to uppercase i.e
```text
    input: Ready, set, go (cap) !
    output: Ready, set, Go !
```
Here is another instance where the same can apply.

    input: Ready, set, go (cap, 2) !
    output: Ready, Set, Go !


### Punctuations
Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one i.e 

    input: I was sitting over there ,and then BAMM !!
    output: I was sitting over there, and then BAMM!!

Except for groups of punctuations which should be formated in the following way

    input: I was thinking ... You were right
    output: I was thinking... You were right

For the punctuation mark ' it should be placed on the right and left of the word in the middle i.e 

    input: I am exactly how they describe me: ' awesome '
    output: I am exactly how they describe me: 'awesome'

If it is a set of words it should be placed at corresponding places.

### Artcles
Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h example

    input: There it was. A amazing rock!
    output: There it was. An amazing rock!



## USAGE

First you need to download the program

```bash

cd ~
git clone https://learn.zone01kisumu.ke/git/wnjuguna/go-reloaded.git

```

After cloning the program, go to it's directory. Here you can edit the sample file with you text than you wish to be changed and save it.

```bash
cd go-reloaded
code sample.txt
```

After making the changes save and then run the command as follows 

```go

go run . sample.txt result.txt

```
Note: Here I'm using sample.txt and result.txt as an example but you are free to choose the name of you file as long as it is a .txt file.

## COLLABORATORS
This is a solo project but I'm open to any contributors who feel can improve and make my code to be more efficient.