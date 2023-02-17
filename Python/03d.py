count=0
a,b,c = (int(x) for x in input().split())
flag=[0]*10001
for i in range(1,c+1):#<-search for 1 to c
    if c%i==0:
        flag[i]=1
if a==b:
    if c%b==0:
        count=1
else:
    for i in range(a,b+1):
        if flag[i]==1:
            count+=1
            #print(i)#debug
print(count)