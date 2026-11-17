# Go-Reloaded: Overview

## Go-Reloaded is a text processing program written in Go.

    The program reads an input file, applies specific text transformations, and writes the modified content to an output file.

_This project demonstrates:_

- File handling in Go

- String manipulation

- Tokenization

- Parsing logic

- Text formatting rules

**_Requirements_**

-   Go (standard library only)

-   No external packages allowed

-   Code must follow good practices

-   The program must be written in Go

_Usage_

    Run the program
    go run . input.txt output.txt

Where:

    input.txt → file containing the original text

    output.txt → file where the modified text will be written

Supported Transformations

1. Number Conversions

    (hex) → Convert the previous hexadecimal word to decimal.

    (bin) → Convert the previous binary word to decimal.

2. Text Case Modifiers

    (up) → Convert the previous word to uppercase.

    (low) → Convert the previous word to lowercase.

    (cap) → Capitalize the previous word.

**_With word count:_**

    (up, n) → Convert the previous n words to uppercase.

    (low, n) → Convert the previous n words to lowercase.

    (cap, n) → Capitalize the previous n words.

3. Punctuation Rules

    The following punctuation marks:

    . , ! ? : ;

**_Rules:_**

    Must be attached to the previous word.

    Must have a space after them (unless at the end of the text).

    Multiple punctuation marks must be handled correctly.

**_Special groups:_**

    ... must remain grouped.

    !? must remain grouped.

4. Apostrophe Rules

    Text between two apostrophes (') must have:

    No spaces immediately inside the quotes.

    Apostrophes must be placed directly around the enclosed text.

Example:

' awesome ' → 'awesome'
5. Article Rule

    The word a must become an when the next word begins with:

    a vowel (a, e, i, o, u)

    or the letter h

Example:

    a apple → an apple
    Project Structure

**_Recommended structure:_**

    go-reloaded/
    ├── go.mod
    ├── go-reloaded.go
    ├── reloaded-unit-test.go
    ├── sample.txt
    ├── result.txt
    └── README.md
Testing

I should:

    Create a test files.

    Verify all transformation rules.

    Compare output with expected results.

    Test edge cases.

**_Learning Objectives_**

## This project helps me understand:

    Go file system API

    Input/output handling

    String and number manipulation

    Program structure

    Clean coding practices
