while True:
    s=input()
    l=len(s)
    if l==1:
        s=int(s)
        if s==0:
            break
        print(s)
    else:
        sum=0
        for i in range(l):
            tmp=int(s[i])
            sum+=tmp
        print(sum)

#mod
'''
while True:
    x=int(input())
    if x==0:break
    ans=0
    while x>0:
        ans+=x%10
        x//=10
    print(ans)
'''

#trans int change
'''
while True:
    N = input()
    if len(N) == 1 and int(N) == 0:
        break
    sum = 0
    for i in range(len(N)):
        sum += int(N[i])

    print(sum)
'''