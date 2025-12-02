commands = []

with open('input.txt', 'r') as f:
    for line in f:
        direction = 1
        if line[0] == 'L':
            direction = -1
        increment = int(line[1:])

        commands.append((direction, increment))

def part1(commands):
    dial, result = 50, 0
    for direction, increment in commands:
        dial += direction * increment

        while dial > 99:
            dial -= 100
        while dial < 0:
            dial += 100

        if dial == 0:
            result += 1

    return result

'''
this function performs a single rotation

if rotations remaining is zero, return
<do rotation>
if dial is now at zero add to counter
decrement rotations
recurse
'''
def rotate(dial: int, direction: int, increments: int, result: int = 0) -> int:
    if increments == 0:
        return dial

    # TODO: do rotation
    # when passing zero set to 1 or 99 depending on direction

    # dial += direction
    # if dial > 99 ...
    # if dial < 0 ...

    if dial == 0:
        result += 1

    increments -= 1

    rotate(dial, direction, increments, result)

def part2(commands):
    dial, result = 50, 0

    for direction, increment in commands:
        result += rotate(dial, direction, increment)

    return result

print('Part 1:', part1(commands))
print('Part 2:', part2(commands))

