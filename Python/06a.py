#range
n1=int(input())
tmp=input().split()
for i in range(n1-1,-1,-1):
    tmp[i]=int(tmp[i])
    print(tmp[i],end='')
    if i!=0:
        print(" ",end='')
print()

#reversed
'''
n2=int(input())
tmp=input().split()
l=len(tmp)
count=0
for j in reversed(tmp):
    count+=1
    print(j,end='')
    if count!=l:
        print(" ",end='')
print()

'''