# Genome Schema

## Structure

Genome is encoded in a uint256, having a leading magic number. The genome of a Ducky is composed of 32 gene of 8 bits each. the traits are starting from the LSB, and the header is starting from the MSB.

### Example:

`[0xD1]-[01]-[00]-[00]-[00]-[00]-[...]-[29]-[04]-[12]-[01]`

## Magic Number

The first genotype base genome which is used the first implementation of the game holds the magic number `209`, which is the letter &ETH; and Hexadecimal `0xD1`

Following Genome Schema must increment this number for example the Mythic DNA has `0xD2` magic number.

The Magic number is on the **most significant byte** (MSB)

## Base Schema 0xD1

| Pos  | Attribute      | Distribution  | Actual Qty of Traits |
| ---- | -------------- | ------------- | -------------------- |
| 1    | Collection     | static        | 3                    |
| 2    | Rarity         | static        | 4                    |
| 3    | Color          | even          | 4                    |
| 4    | Family         | even          | 5                    |
| 5    | Body           | uneven        | 11                   |
| 6    | Head           | uneven        | 26                   |
| 7    | Eyes           | uneven        | 30                   |
| 8    | Beak           | uneven        | 14                   |
| 9    | Wings          | uneven        | 10                   |
| 10   | Forename       | even          | 36                   |
| 11   | Temper         | uneven        | 16                   |
| 12   | Skill          | uneven        | 12                   |
| 13   | Habitat        | even          | 5                    |
| 14   | Breed          | uneven        | 28                   |
| 15   | Empty          |               |                      |
| 16   | Empty          |               |                      |
| 17   | Empty          |               |                      |
| 18   | Empty          |               |                      |
| 19   | Empty          |               |                      |
| 20   | Empty          |               |                      |
| 21   | Empty          |               |                      |
| 22   | Empty          |               |                      |
| 23   | Empty          |               |                      |
| 24   | Empty          |               |                      |
| 25   | Empty          |               |                      |
| 26   | Empty          |               |                      |
| 27   | Empty          |               |                      |
| 28   | Empty          |               |                      |
| 29   | Empty          |               |                      |
| 30   | Empty          |               |                      |
| 31   | Reserved Flags | Store options |                      |
| 32   | Magic Number   | 209           | 0xD1                 |

## Mythic Schema 0xD2

| Pos  | Attribute      | Distribution | Actual Qty of Traits     |
| ---- | -------------- | ------------ | ------------------------ |
| 1    | Collection     | static       | 3                        |
| 2    | UniqId         | static       | 60                       |
| 3    | Temper         | uneven       | 16                       |
| 4    | Skill          | uneven       | 12                       |
| 5    | Habitat        | even         | 5                        |
| 6    | Breed          | uneven       | 20                       |
| 7    | Birthplace     | even         | 5 (Same as habitat)      |
| 8    | Quirk          | uneven       | 10                       |
| 9    | Favorite Food  | uneven       | 8                        |
| 10   | Favorite Color | even         | 4 (Same as basic colors) |
| 11   | Empty          |              |                          |
| 12   | Empty          |              |                          |
| 13   | Empty          |              |                          |
| 14   | Empty          |              |                          |
| 15   | Empty          |              |                          |
| 16   | Empty          |              |                          |
| 17   | Empty          |              |                          |
| 18   | Empty          |              |                          |
| 19   | Empty          |              |                          |
| 20   | Empty          |              |                          |
| 21   | Empty          |              |                          |
| 22   | Empty          |              |                          |
| 23   | Empty          |              |                          |
| 24   | Empty          |              |                          |
| 25   | Empty          |              |                          |
| 26   | Empty          |              |                          |
| 27   | Empty          |              |                          |
| 28   | Empty          |              |                          |
| 29   | Empty          |              |                          |
| 30   | Empty          |              |                          |
| 31   | Reserved Flags |              |                          |
| 32   | Magic Number   | 210          | 0xD2                     |
