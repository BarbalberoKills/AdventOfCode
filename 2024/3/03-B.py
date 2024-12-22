import re
import pathlib


def mul(a, b):
    return a*b


#check do/don't passed instruction against the actual running instruction
def switch(switch, status):
    match(switch):
        case("do()"):
            return True
        case("don't()"):
            return False
        case(""):
            return status


script_run = str(pathlib.Path(__file__).parent.resolve())
puzzle_input = ""
with open(script_run + "/03-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        puzzle_input += line.strip()

exp = re.compile(r"(don't\(\)|do\(\)|mul\(\d+,\d+\))")
numbers_exp = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')

#parse puzzle input and store useful instructions into list
mul_instructions = re.findall(exp, puzzle_input)

add_up = 0
doing = True

for i in mul_instructions:
    # print(i)
    if i == "don't()" or i == "do()":
        doing = switch(i, doing)
    else:
        numbers_to_multiply = re.match(numbers_exp, i)
        # print(doing)
        if doing == True:
            add_up += mul(int(numbers_to_multiply.group(1)), int(numbers_to_multiply.group(2)))

print(add_up)