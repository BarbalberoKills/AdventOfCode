import pathlib


script_run = str(pathlib.Path(__file__).parent.resolve())
puzzle_input = ""
with open(script_run + "/4/04-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        puzzle_input += line

grid = []
for c, line in enumerate(puzzle_input.splitlines()):
    grid.append(line)



def check_direction(grid, row, col, word, d_row, d_col):
    for i in range(len(word)):
        new_row = row + i * d_row
        new_col = col + i * d_col

        # Check bounds
        if not (0 <= new_row < len(grid) and 0 <= new_col < len(grid[0])):
            return False

        # Check character match
        if grid[new_row][new_col] != word[i]:
            return False

    return True


def find_word(grid, word):
    count = 0
    directions = [
        (0, 1), (0, -1), (1, 0), (-1, 0), 
        (1, 1), (-1, -1), (1, -1), (-1, 1)
    ]
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            for d_row, d_col in directions:
                if check_direction(grid, row, col, word, d_row, d_col):
                    count += 1
    return count





grid = [list(row) for row in grid]
word = "XMAS"

result = find_word(grid, word)
print(f"Number of occurrences of '{word}': {result}")