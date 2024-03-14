#! /usr/bin/env python3
# TO BE RUN IN THE ROOT FOLDER

## COMMON FOR BOTH PARTS

file = "../inputs/day3_input.txt"
with open(file) as f:
    lines = f.read().strip().split('\n')
    f_lines = list(lines)

rows = len(f_lines)
cols = len(f_lines[0])
ARROUND_FILDS = [ (-1, -1), (-1, 0), (-1, 1), (0, -1), (1, -1), (1, 0), (1, 1), (0, 1) ]

########### Part 1 ###########
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

########### Part 2 ###########

def find_number(row, col, matrix):
    og_col = col
    new_list = []
    # format row-col
    first_index = f'{row}-{col}'

    # <<<<
    while col >= 0 and matrix[row][col].isdigit():
        new_list.insert(0, matrix[row][col])
        first_index = f'{row}-{col}'
        col -= 1
    # >>>>
    col = og_col + 1
    while col < cols and matrix[row][col].isdigit():
        new_list.append(matrix[row][col])
        col += 1

    if not new_list:
        return 1
    return first_index, int(''.join(new_list))


def gear_ratio_processor(row, col, matrix):
    # return the number of times the * touches the numbers
    rows = len(matrix)
    cols = len(matrix[0])
    gears = dict()
    for i, j in ARROUND_FILDS:
        if row + i < 0 or col + j < 0:
            continue
        if row + i >= rows or col + j >= cols:
            continue
        if matrix[row + i][col + j].isdigit():
            first_index, full_number = find_number(row + i, col + j, matrix)
            gears.update({first_index: full_number})

    if len(gears) < 2:
        return 0
    from functools import reduce
    return reduce(lambda x, y: x * y, gears.values())

total = 0

row = 0
while row < rows:
    col = 0
    while col < cols:
        if f_lines[row][col] == '*':
            total += gear_ratio_processor(row, col, f_lines)
        col += 1
    row += 1

print(f"part 2: {total}")
