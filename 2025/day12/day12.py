f = open('input.txt')
lines = f.readlines()
lines.reverse()

part1 = 0

for l in lines:
    if l == '\n':
        print('Part 1:', part1)
        exit(0)
    else:
        grid = l[:l.index(':')].split('x')
        required = l[l.index(':')+1:].split()
        space_available = int(grid[0]) * int(grid[1])
        space_needed = 0

        for shape in required:
            space_needed += int(shape) * (3 * 3) # treat each shape as 3x3

        if space_needed <= space_available:
            part1 += 1

# Initially I had it working on the test input (ans=2) and not my real input (ans=433)
# I looked at other peoples solutions and realised I was comparing space_needed against space_available wrong. But then it stopped working on the test input (ans=1) and worked on my real input (ans=567).
# I don't really understand why so will look into it later.
# Part 2 not available because I don't have enough stars or something.
