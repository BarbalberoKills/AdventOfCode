import re
import pathlib


def mul(a, b):
    return a*b


#check do/don't passed instruction against the actual running instruction
def switch(status, s):
    if s == "do()" and status == True:
        return True
    elif s == "don't()" and status == False:
        return False
    elif s == "do()" and status == False:
        return True
    elif s == "don't()" and status == True:
        return False
    elif s == "":
        return status


script_run = str(pathlib.Path(__file__).parent.resolve())
puzzle_input = ""
with open(script_run + "/03-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        puzzle_input += line.strip()

exp = re.compile(r"(don't\(\)|do\(\))?.?(mul\(\d+,\d+\))")
numbers_exp = re.compile(r'(\d*),(\d*)')

#parse puzzle input and store useful instructions into list
mul_instructions = re.findall(exp, puzzle_input)

add_up = 0
doing = True
for i in mul_instructions:
    numbers_to_multiply = re.findall(numbers_exp, i[1])
    doing = switch(doing, i[0])
    print(doing, i)
    #print(doing)
    if doing == True : add_up += mul(int(numbers_to_multiply[0][0]), int(numbers_to_multiply[0][1])) 
    #print("-------------------------------------")
print(add_up)