#! /usr/bin/env python3

# Part 1

with open('python/input.txt') as f:
    lines = f.read().strip().split('\n')
    f_lines = list(lines)

# Helper functions
rows = len(f_lines)
cols = len(f_lines[0])
ARROUND_FILDS = [ (-1, -1), (-1, 0), (-1, 1), (0, -1), (1, -1), (1, 0), (1, 1), (0, 1) ]

def is_symbol(c) -> bool:
# return true if c is not a number or a '.'
	return not c.isdigit() and c != '.'

def touches(row, col, matrix):
# return True if the position (row, col) touches a symbol
	rows = len(matrix)
	cols = len(matrix[0])
	for i, j in ARROUND_FILDS:
		if row + i < 0 or col + j < 0:
			continue
		if row + i >= rows or col + j >= cols:
			continue
		if is_symbol(matrix[row + i][col + j]):
			return True
	return False

touched = False
acc = '0'
total = 0

row = 0
while row < rows:
    col = 0
    acc = ''
    while col < cols:
        if f_lines[row][col].isdigit():
            if not touched:
                touched = touches(row, col, f_lines)
            acc += f_lines[row][col]

            if col == cols - 1 and touched:
                # This is just in case it is the last number in the row and is digit
                total += int(acc)   
        else:
            if touched and acc:
                total += int(acc)
            acc = ''
            touched = False
        col += 1
    row += 1

print(f"part 1: {total}")


