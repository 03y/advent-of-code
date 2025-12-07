import numpy as np

# open file, read each line into a row, then put into a matrix and rotate 90deg
rotated = np.rot90(np.array([l.split() for l in open('input.txt')]), k=-1)
part1 = 0

for row in rotated:
    row_value, operator = 0, ''
    for i in range(len(row)):
        if i == 0:
            operator = row[i]
        elif i == 1:
            row_value = int(row[i])
        else:
            if operator == '*':
                row_value *= int(row[i])
            elif operator == '+':
                row_value += int(row[i])
    part1 += row_value

print('Part 1:', part1)

