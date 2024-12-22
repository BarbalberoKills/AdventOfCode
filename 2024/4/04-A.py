import pathlib


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



total_words = []
for actual_row in range(len(mat)):
    for actual_col in range(len(mat[actual_row])):
        p = mat[actual_row][actual_col]
        size_word = 4
        word = ""
        #print("RIGHT", p)
        for lenght_word in range(size_word):
            if actual_col + size_word <= len(mat[actual_row]):
                next_pos = move([actual_row, actual_col], "right", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("BOTTOM-RIGHT", p)
        for lenght_word in range(size_word):
            if (actual_col + size_word <= len(mat[actual_row])) and (actual_row + size_word <= len(mat)):
                next_pos = move([actual_row, actual_col], "bottom-right", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("BOTTOM", p)
        for lenght_word in range(size_word):
            if actual_row + size_word <= len(mat):
                next_pos = move([actual_row, actual_col], "bottom", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("BOTTOM-LEFT", p)
        for lenght_word in range(size_word):
            if (actual_col - 3 >= 0) and (actual_row + size_word <= len(mat)):
                next_pos = move([actual_row, actual_col], "bottom-left", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("LEFT", p)
        for lenght_word in range(size_word):
            if actual_col - 3 >= 0:
                next_pos = move([actual_row, actual_col], "left", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("UP-LEFT", p)
        for lenght_word in range(size_word):
            if (actual_col - 3 >= 0) and (actual_row - size_word >= 0):
                next_pos = move([actual_row, actual_col], "up-left", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("UP", p)
        for lenght_word in range(size_word):
            if actual_row - size_word >= 0:
                next_pos = move([actual_row, actual_col], "up", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
        #print("UP-RIGHT", p)
        for lenght_word in range(size_word):
            if (actual_col + size_word <= len(mat[actual_row])) and (actual_row - 4 >= 0):
                next_pos = move([actual_row, actual_col], "up-right", lenght_word)
                #print(next_pos)
                word += str(mat[next_pos[0]][next_pos[1]])
        if word != "": total_words.append(word)
        word = ""
        #print("------------------------------------")
#print(total_words)


xmas_counter = 0
for word in total_words:
    if (''.join(word) == "XMAS"):
        xmas_counter +=1

print(xmas_counter)