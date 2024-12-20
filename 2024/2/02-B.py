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


def safe_check(elements_to_check):
    variation = []
    difference = []
    for i in range(len(elements_to_check)-1):
        variation.append(compare(elements_to_check[i], elements_to_check[i+1]))
        difference.append(distance(elements_to_check[i], elements_to_check[i+1]))
    # print(variation, difference)
    # print(all(difference))
    if (all(variation) == any(variation)) and all(difference):
        return True


safe_report = 0
for report in report_data:
    if safe_check(report):
        safe_report += 1
    else:
        for i in range(len(report)):
            modified_report = report[:i] + report[i+1:]
            if safe_check(modified_report):
                safe_report += 1
                break


print(safe_report)
