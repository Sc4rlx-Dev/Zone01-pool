    /problem overview : ------------------------------------------------------------------------------
    
- takes 9x9 grid
- parse the input to convert it into d array (9x9)
- prints the solved Sudoku grid || an error msg

##    /parsing input :

- accepte args , an array of 9 strings 
- converts this strings into 9x9 integer grid
- translate '.' (empty cells) into 0 and ensures other
- characters are valid numbers (1 - 9)


##    /solving sudoku :

- it find an empty cell (with 0)
- put the numbers into , 1-9
- check the placement is valid (isValidChecker)
- trying to solve the rest of the grid (recersivity)
- if the nember leads to a solution > return TRUE , Otherwise , it backtracks (remove the number and tries the next)

##    /find Empty Cells (FindEmptyCell) :

- check into the grid row by row to find the first empty cell (0)


##    /checking the validation placements(isValidChecker) : 
- isnt in the same row or col
- isnt in same 3x3

##    /printing the Grid (printGrid) :
- formats and print the solved grid neatly


//backtracking solver
//check row , colum, and block consraints
-----------------------------------------------------------------------------------------------------

