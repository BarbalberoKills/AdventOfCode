import re
reg = re.compile(r'(\d*)\s')

l = ["71 73 74 76 78 80 77"]
print(l[0])

m=re.match(reg, l[0])

print(m.group(2))