import re
reg = re.compile(r'(\d*)\s')

def switch(status, s):
    if s == "do()" and status == True:
        return False
    elif s == "don't()" and status == False:
        return False
    elif s == "do()" and status == False:
        return True
    elif s == "don't()" and status == True:
        return True
    

print(switch(False, "do()"))