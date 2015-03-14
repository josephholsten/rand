rand(1) -- print a random number
================================

## SYNOPSIS

`rand` [max [min [count]]]

## DESCRIPTION

`rand` will print a pseudorandom positive integer between 0 and *max*, or *MAX_INT* if no max is provided.

It is designed as a less frustrating alternative to bash's `$RANDOM`, BSD `jot(1)`, or having to dig into a more powerful scripting language for a random function.


## EXAMPLES

The following are equivalant:

Print an integer between 0 and 512:

    rand 512
    jot -r 1 0 512
    echo $(( $RANDOM % 512 ))
    awk 'BEGIN {srand(); print int(rand() * 512)}'





## AUTHOR

Joseph Anthony Pasquale Holsten <joseph@josephholsten.com>

## SEE ALSO

awk(1),
bash(1)
jot(1)
seq(1)
rand(3)

https://launchpad.net/rand
