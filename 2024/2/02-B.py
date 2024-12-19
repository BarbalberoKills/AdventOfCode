import pathlib


script_run = str(pathlib.Path(__file__).parent.resolve())


def compare(a, b):
    if a > b:
        return 1
    elif a < b:
        return 0
    else:
        return None
    
def distance(a, b):
    if abs(a-b) > 0 and abs(a-b) <= 3:
        return 1
    else:
        return 0


report_data = []
with open(script_run + "/02-A_inputlist.txt", "r") as the_file:
    for line in the_file:
        report_data.append([int(item) for item in line.strip().split(" ")])
# for i in report_data:
#     print(i)


safe_report = 0
for report in report_data:
    variation = []
    com = 0
    dist = 0
    tollerete = 1
    print(report)
    for i in range(0, len(report)-1):
        #print(report[i], report[i+1])
        variation.append(compare(report[i], report[i+1]))
        if distance(report[i], report[i+1]) == 1:
            dist +=1
        elif distance(report[i], report[i+1]) == 0 and tollerete == 1:
            dist +=1
            tollerete = 0
    print(variation, dist, tollerete)
    if (variation.count(0) == len(report)-1 or variation.count(1) == len(report)-1) and dist == len(report)-1:
        print("safe1", report)
        safe_report += 1
    elif (variation.count(0) == len(report)-2 or variation.count(1) == len(report)-2) and (dist == len(report)-1) and tollerete == 1:
        print("safe2", report)
        safe_report += 1
    # if (com == len(report)-1 or com == 0) and dist == len(report)-1:
    #     print("safe", report)
    #     safe_report += 1

print(safe_report)
