import pathlib


def move(pos, dir, increment):
    match(dir):
        case("right"):
            pos[1] += abs(increment)
            return pos
        case("bottom-right"):
            pos[0] += abs(increment)
            pos[1] += abs(increment)
            return pos
        case("bottom"):
            pos[0] += abs(increment)
            return pos
        case("bottom-left"):
            pos[0] += abs(increment)
            pos[1] -= abs(increment)
            return pos
        case("left"):
            pos[1] -= abs(increment)
            return pos
        case("up-left"):
            pos[0] -= abs(increment)
            pos[1] -= abs(increment)
            return pos
        case("up"):
            pos[0] -= abs(increment)
            return pos
        case("up-right"):
            pos[0] -= abs(increment)
            pos[1] += abs(increment)
            return pos
        

script_run = str(pathlib.Path(__file__).parent.resolve())
puzzle_input = ""
with open(script_run + "/04-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        puzzle_input += line


mat = []
for c, line in enumerate(puzzle_input.splitlines()):
    mat.append([])
    for element in line:
        mat[c].append(element)


directions = ["right", "bottom-right", "bottom", "bottom-left", "left", "up-left", "up", "up-right"]
word_to_find = "XMAS"
xmas_counter = 0
for actual_row in range(len(mat)):
    for actual_col in range(len(mat[actual_row])):
        if mat[actual_row][actual_col] == "A":
            up_left = move([actual_row, actual_col], "up-left", 1)
            up_right = move([actual_row, actual_col], "up-right", 1)
            bottom_right = move([actual_row, actual_col], "bottom-right", 1)
            bottom_left = move([actual_row, actual_col], "bottom-left", 1)
            if mat[up_left[0]][up_left[1]]+mat[actual_row][actual_col]+mat[bottom_right[0]][bottom_right[1]] == "MAS" and mat[bottom_left[0]][bottom_left[1]]+mat[actual_row][actual_col]+mat[up_right[0]][up_right[1]] == "MAS":
#         for dir in directions:
#             next_pos = move([actual_row, actual_col], dir, (len(word_to_find)-1))
#             if (0 <= next_pos[0] < len(mat)) and (0 <= next_pos[1] < len(mat[actual_row])):
#                 word = ""
#                 for position_in_word in range(len(word_to_find)):
#                     next_pos = move([actual_row, actual_col], dir, position_in_word)
#                     word += str(mat[next_pos[0]][next_pos[1]])
#                 if word == "XMAS":
#                     xmas_counter += 1
# print(xmas_counter)