r,c=(int(x) for x in input().split())
a=[[0]*(c+1) for j in range(r+1)]
for i in range(r):
    tmp=list(map(int,input().split()))
    #if map assign to "a" directly, number of cow be changed!! 
    for j in range(c):
        a[i][j]=tmp[j]
#sum row
for i in range(r):
    for j in range(c):
        a[i][c]+=a[i][j]
#sum cow
for i in range(c+1):
    for j in range(r):
        a[r][i]+=a[j][i]
for i in range(r+1):
    for j in range(c+1):
        if j!=c:
            print(a[i][j],end=' ')
        else:
            print(a[i][j])