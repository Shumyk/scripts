#!/bin/bash
#
# assign a value to a variable
word='script'

# display that value using the variable
echo "$word"

# demonstrate that single quotes cause variables to NOT get expanded.
echo '$word'
echo $word

# combine the variable with hard-coded text
echo "This is a shell $word"

# display the contents of the variable using an alternative syntax
echo "This is a shell ${word}"

# append text to the variable
echo "${word}ing is fun!"

# show how NOT to append text to a variable
echo "$wording is fun!"

# create a new variable
ENDING='ed'

# combine the two variables
echo "This is ${word}${ENDING}."

# change the value stored in the enDING varible (reassignemnt)
ENDING='ing'
echo "${word}${ENDING} is fun!"

# reassign value to ENDING
ENDING='s'
echo "You are going to write many ${word}${ENDING} in this class!"
