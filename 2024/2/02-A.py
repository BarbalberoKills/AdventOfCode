import pathlib


script_run = str(pathlib.Path(__file__).parent.resolve())

    
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
    if report == sorted(report) or report == sorted(report, reverse=True):
        acc = 0
        for i in range(0,len(report)-1):
            if distance(report[i], report[i+1]) == 1:
                acc += 1
        if acc == len(report)-1:
            safe_report += 1
print(safe_report)
