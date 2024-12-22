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
    
def test(status, s):
    match (s, status):
        case ("do()", True):
            return False
        case ("don't()", False):
            return False
        case ("do()", False):
            return True
        case ("don't()", True):
            return True
        

print(test(False, "do()"))