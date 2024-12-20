import re
import pathlib


script_run = str(pathlib.Path(__file__).parent.resolve())

def mul(a, b):
    return a*b

s = ""
with open(script_run + "/03-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        s += line.strip()

print(s)


exp = re.compile(r'(mul\(\d+,\d+\))')
num_exp = re.compile(r'(?P<first>\d*),(?P<second>\d*)')


#s = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(110,8)mul(8,50))"

mul_instructions = re.findall(exp, s)

add_up = 0
for i in mul_instructions:
    num = re.findall(num_exp, i)
    #print(i)
    add_up += mul(int(num[0][0]), int(num[0][1]))
print(add_up)