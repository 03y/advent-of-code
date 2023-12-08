#!/bin/python3

plays = []
winnings = 0
card_priority = {'T': 'A', 'J': 'B', 'Q': 'C', 'K': 'D', 'A': 'E'}

def evaluate(hand):
    counts = [hand.count(card) for card in hand]

    if 5 in counts:
        return 6
    if 4 in counts:
        return 5
    if 3 in counts:
        if 2 in counts:
            return 4
        return 3
    if counts.count(2) == 4:
        return 2
    if 2 in counts:
        return 1
    return 0

def score(hand):
    return (evaluate(hand), [card_priority.get(card, card) for card in hand])

for line in open("input.txt"):
    hand, bid = line.split()
    plays.append((hand, int(bid)))

plays.sort(key=lambda play: score(play[0]))

for rank, (hand, bid) in enumerate(plays, 1):
    winnings += rank*bid

print('Part 1:', winnings)
