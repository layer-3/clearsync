# Package precision

A Go library to handle numbers with a specified precision
by providing functionality to round numbers to a given number
of significant digits while also considering a maximum number of decimals.

## How it works

Significant figures, also known as the precision of a number in positional notation, are digits in the number that are reliable and necessary to indicate the quantity of something.

The precision level of all trading prices is calculated based on significant figures.

Some examples of five significant digits are 1.0234, 10.234, 120.34, 1234.5, 0.012345, and 0.00012340.

This is similar to how traditional global markets handle the precision of small, medium, and large values.
The reasoning behind this is that the number of decimals becomes less important as the amount increases.
The opposite is true for minimal amounts, where greater precision is more valuable.

If number of decimals is larger than maximum allowed decimals parameter, the decimals are truncated to fit into maximum allowed.
