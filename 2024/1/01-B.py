import re

exp = re.compile(r'(\d*)\s*(\d*)')

listA =[]
listB =[]
with open("01-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        m = re.match(exp, line.strip())
        if m:
            listA.append(int(m.group(1)))
            listB.append(int(m.group(2)))


simScore = 0
for i in listA:
    count = 0
    for n in listB:
        if i == n:
            count += 1
    simScore += i*count
    print(simScore)