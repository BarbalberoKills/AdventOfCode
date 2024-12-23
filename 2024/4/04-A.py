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
    for k, element in enumerate(line):
        mat[c].append(element)


directions = ["right", "bottom-right", "bottom", "bottom-left", "left", "up-left", "up", "up-right"]
word_to_find = "XMAS"
xmas_counter = 0
for actual_row in range(len(mat)):
    for actual_col in range(len(mat[actual_row])):
        for dir in directions:
            next_pos = move([actual_row, actual_col], dir, len(word_to_find))
            if (0 <= next_pos[0] <= len(mat[actual_row])) and (0 <= next_pos[1] <= len(mat)):
                word = ""
                for position_in_word in range(len(word_to_find)):
                    next_pos = move([actual_row, actual_col], dir, position_in_word)
                    word += str(mat[next_pos[0]][next_pos[1]])
                if word == "XMAS":
                    xmas_counter += 1
print(xmas_counter)
