

mat = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [8, 10, 11, 12],
    [13, 14, 15, 16],
    [17, 18, 19, 20]
]


center = [1, 1]

# print(mat[center[0]][center[1]])

def move(pos, dir, increment):
    match(dir):
        case("right"):
            pos[1] += abs(increment)
            return pos
        case("bottom-right"):
            pos[0] -= abs(increment)
            pos[1] += abs(increment)
            return pos
        case("bottom"):
            pos[0] -= abs(increment)
            return pos
        case("bottom-left"):
            pos[0] -= abs(increment)
            pos[1] -= abs(increment)
            return pos
        case("left"):
            pos[1] -= abs(increment)
            return pos
        case("up-left"):
            pos[0] += abs(increment)
            pos[1] -= abs(increment)
            return pos
        case("up"):
            pos[0] += abs(increment)
            return pos
        case("up-right"):
            pos[0] -= abs(increment)
            pos[1] += abs(increment)
            return pos


class Point():

    def __init__(self, x, y):
        '''Defines x and y variables'''
        self.X = x
        self.Y = y

for actual_row in range(len(mat)):
    for actual_col in range(len(mat[actual_row])):
        p = mat[actual_row][actual_col]
        print("RIGHT", p)
        for lenght_word in range(4):
            if actual_col + 4 <= len(mat[actual_row]):
                next_pos = move([actual_row, actual_col], "right", lenght_word)
                print(next_pos)
        print("------------------------------------")
        print("BOTTOM-RIGHT", p)
        for lenght_word in range(4):
            if (actual_col + 4 <= len(mat[actual_row])) and (actual_row + 4 <= len(mat[actual_row])):
                next_pos = move([actual_row, actual_col], "bottom-right", lenght_word)
                print(next_pos)
        print("------------------------------------")
        print("BOTTOM", p)
        for lenght_word in range(4):
            if actual_row + 4 <= len(mat[actual_row]):
                next_pos = move([actual_row, actual_col], "bottom", lenght_word)
                print(next_pos)
        print("------------------------------------")
        print("BOTTOM-LEFT", p)
        for lenght_word in range(4):
            if (actual_col - 3 >= 0) and (actual_row + 4 <= len(mat[actual_row])):
                next_pos = move([actual_row, actual_col], "bottom-left", lenght_word)
                print(next_pos)
        print("------------------------------------")
        print("LEFT", p)
        for lenght_word in range(4):
            if actual_col - 3 >= 0:
                next_pos = move([actual_row, actual_col], "left", lenght_word)
                print(next_pos)
        print("------------------------------------")


        print("UP", p)
        for lenght_word in range(4):
            if actual_row - 3 >= 0:
                next_pos = move([actual_row, actual_col], "up", lenght_word)
                print(next_pos)
        print("------------------------------------")
        print("UP-RIGHT", p)
        for lenght_word in range(4):
            if (actual_col + 4 <= len(mat[actual_row])) and (actual_row - 3 >= 0):
                next_pos = move([actual_row, actual_col], "up-right", lenght_word)
                print(next_pos)
        print("------------------------------------")

