# Colorized Elementary Cellular Automata

This program generates space-time diagrams for elementary cellular automata and color-codes cells based on the state configuration from the previous generation that determined the cell state.

For elementary cellular automata, the next generation of a cell depends on the state of three cells: the neighbor cell to the left, the cell itself and the neighbor cell to the right.
This means that there are 2^3 = 8 possible state configurations that determine the state of a cell in the next generation.

By assigning a color to numbers 0-7, we can color-code each cell.
For example, if the left neighbor is 'on' (1), the cell itself 'off' (0) and the right neighbor 'off' (0); this state configuration is encoded as binary 100, which translates to color number 4 in decimal.
