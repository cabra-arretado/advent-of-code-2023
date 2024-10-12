from collections import deque

raw_lines = """
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""

# Existing code to parse the input and create hash_table
lines = raw_lines.strip().split('\n')
f_lines = list(lines)

class Card:
    card_number: str
    number_of_matches: int
    lotery_numbers: list[int]
    winning_numbers: list[int]

    def __init__(self, card_number, number_of_matches, lotery_numbers, winning_numbers) -> None:
        self.card_number = card_number
        self.number_of_matches = number_of_matches
        self.lotery_numbers = lotery_numbers
        self.winning_numbers = winning_numbers

def does_number_match(winning_number, lotery_numbers):
    return winning_number in lotery_numbers

def get_number_of_matches(winning_numbers, lotery_numbers):
    number_of_matches = 0
    for ea in winning_numbers:
        if does_number_match(ea, lotery_numbers):
            number_of_matches += 1
    return number_of_matches

def process_line(line):
    """Process the line returning a card"""
    divided = line.strip().split(':')
    card_number = divided[0].strip().split(' ')[1]
    divided = divided[1].strip().split('|')
    lotery_numbers = divided[0].strip().split(' ')
    lotery_numbers = [int(ea) for ea in lotery_numbers if ea]
    winning_numbers = divided[1].strip().split(' ')
    winning_numbers = [int(ea) for ea in winning_numbers if ea]
    number_of_matches = get_number_of_matches(winning_numbers, lotery_numbers)
    card = Card(card_number, number_of_matches, lotery_numbers, winning_numbers)
    return card

# Parse the input and fill the hash_table
hash_table = {}
for ea in f_lines:
    card = process_line(ea)
    hash_table[card.card_number] = card

# Initialize the queue and instances dictionary
queue = deque()
for card_number in hash_table.keys():
    queue.append(card_number)

instances = {}

# Simulate the scratchcard processing
while queue:
    current_card_number = queue.popleft()
    
    # Increment the count of instances for this card
    if current_card_number in instances:
        instances[current_card_number] += 1
    else:
        instances[current_card_number] = 1

    # Get the corresponding Card object
    card = hash_table[current_card_number]

    # Get the number of matches
    num_matches = card.number_of_matches

    if num_matches > 0:
        for i in range(1, num_matches + 1):
            next_card_number = str(int(current_card_number) + i)
            if next_card_number in hash_table:
                queue.append(next_card_number)
            else:
                # No more cards to process
                break

# Calculate the total number of scratchcards
total_scratchcards = sum(instances.values())

print(f"Total scratchcards: {total_scratchcards}")

