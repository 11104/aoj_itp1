import string

n=int(input())
s=[0]*13
h=[0]*13
c=[0]*13
d=[0]*13
for i in range(n):
    op,tmp=input().split()
    tmp=int(tmp)
    if op=='S':
        s[tmp-1]=1
    elif op=='H':
        h[tmp-1]=1
    elif op=='C':
        c[tmp-1]=1
    else:
        d[tmp-1]=1

for i in range(13):
    if s[i]==0:
        print("S",i+1)
for i in range(13):
    if h[i]==0:
        print("H",i+1)
for i in range(13):
    if c[i]==0:
        print("C",i+1)
for i in range(13):
    if d[i]==0:
        print("D",i+1)
