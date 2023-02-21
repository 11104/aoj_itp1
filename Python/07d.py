n,m,l=(int(x) for x in input().split())
a=[[0]*(m) for j in range(n)]
for i in range(n):
    tmp=list(map(int,input().split()))
    #if map assign to "a" directly, number of cow be changed!! 
    for j in range(m):
        a[i][j]=tmp[j]
b=[[0]*(l) for j in range(m)]
for i in range(m):
    tmp=list(map(int,input().split()))
    #if map assign to "a" directly, number of cow be changed!! 
    for j in range(l):
        b[i][j]=tmp[j]
c=[[0]*(l) for j in range(n)]
for i in range(n):
	for j in range(l):
			for k in range(m):
				c[i][j] += a[i][k] * b[k][j]
for i in range(n):
	for j in range(l):
            if j!=l-1:
                print(c[i][j],end=' ')
            else:
                print(c[i][j])
