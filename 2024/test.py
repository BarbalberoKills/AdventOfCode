import re
reg = re.compile(r'(\d*)\s')

report = [7, 620, 4, 2, 1]
print(report[:1])
print(report[1+1:])
# print(report[0])

# m=re.match(reg, report[0])

for i in range(len(report)):
    a = report[:i] + report[i+1:]
    print(a)