n,m=(int(x) for x in input().split())
a=[[] for i in range(n)]
for i in range(n):
    a[i]=list(map(int,input().split()))
b=[[0]*1 for i in range(m)]
for i in range(m):
    b[i]=int(input())

for i in range(n):
    tmp=0
    for j in range(m):
        tmp+=a[i][j]*b[j]
    print(tmp)