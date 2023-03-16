# gmiindex

Generate a text/gemini index of the given paths.

Normally, the file (base) name is used as a link name, but if a file has a
.gmi or .gemini extension and starts with a heading, the text of the heading
is used as the link name.  If that file (base) name also starts with what
looks like an YYYY-MM-DD date, that is prepended to the link name.

Links with dates are sorted first, in descending order. Other files are
sorted in dictionary order below any links with dates

## Usage

    gmiindex [FILE]...

* The index is printed to stdout

## Example output

Given the example directory included in the repository:

    $ gmiindex example/*
    => example/2023-03-16.gmi 2023-03-16 - Test title 1
    => example/2023-02-03.gmi 2023-02-03 - Test title 2
    => example/2023-02-02.gmi 2023-02-02.gmi
    => example/2020-01-01.txt 2020-01-01.txt
    => example/nongmi.txt nongmi.txt
    => example/space%20title.txt space title.txt
