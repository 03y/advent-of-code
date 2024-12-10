import re
from datetime import datetime

seed = 0x0000
magic = 0xFFFF
cache = {'first': 0, 'last':-1}
begin_exec = datetime.now()

class Block:
    def __init__(self, size, empty):
        self.size = size
        self.empty = empty

def is_sorted(fs):
    for i in range(len(fs)):
        if fs[i] == chr(magic):
            for j in range(i, len(fs)):
                if fs[j] != chr(magic):
                    return False
    return True

def get_last(fs):
    pos = cache['last']
    while pos > 0:
        if fs[pos] != chr(magic):
            cache['last'] = pos
            return pos
        pos -= 1
    return -1

def get_free(fs):
    for i in range(cache['first'], len(fs)):
        if fs[i] == chr(magic):
            cache['first'] = i
            return i
    return -1

# the example only had 0-9 so i made the dumb assumption that was all the numbers would go to (wrap or something after 9), but no they increase forever so lazy unicode fix lol
def get_unicode_char():
    global seed
    n = seed
    seed += 1
    return chr(n)

with open('input.txt') as f:
    input = f.readline().strip('\n')

    size, empty = 1, False
    blocks = []

    for i in range(len(input)):
        blocks.append(Block(int(input[i]), empty))
        empty = not empty

    id = get_unicode_char()
    used = [id]
    fs = []
    for i in range(len(blocks)):
        if blocks[i].empty == True:
            used += id
            fs += chr(magic) * blocks[i].size
        else:
            fs += str(id) * blocks[i].size
            id = get_unicode_char()

    cache['last'] = len(fs)-1

    sorted = False
    while not sorted:
        if is_sorted(fs):
            sorted = True
        else:
            n = get_last(fs)
            fs[get_free(fs)] = fs[n]
            fs[n] = chr(magic)

    checksum = 0
    n = 0
    for i in range(len(fs)):
        if ord(fs[i]) != magic:
            checksum += i * ord(fs[i])

    print('Part 1:', checksum)
    
    print('execution time:', datetime.now()-begin_exec)

'''
how fast can a brutefoce be?
benchmark: i9-13950HX, 64GB RAM

Part 1: 1928
execution time: 0:00:00.000445

Part 1: 6299243228569
execution time: 0:01:39.892172

implemented start cache:

Part 1: 6299243228569
execution time: 0:01:15.048865

implemented end cache:

Part 1: 6299243228569
execution time: 0:00:54.461692
'''
