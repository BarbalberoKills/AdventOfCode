you: aaa, bbb
aaa: ccc, ddd
ccc: out
ddd: hhh
hhh: out
bbb: ddd


device.name = you
you == out ?? NO
cache[you] esiste ?? NO
child of you: [aaa, bbb]
        |
        |
device.name = aaa
aaa == out ?? NO
cache[aaa] esiste ?? NO
child of you: [ccc, ddd]
        |
        |
        |
        |
        |
        |
        |
device.name = ccc
ccc == out ?? NO
cache[ccc] esiste ?? NO
child of you: [out]         1 cache[ccc] = 1
        |                 /
        |               /
        |             /
        |           /
        |         /
device.name = out
out == out ?? SI
ritorno 1