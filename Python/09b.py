while 1:
    s=str(input())
    if s=='-':
        break
    n=int(input())
    for i in range(n):
        h=int(input())
        tmp1=s[0:h]
        tmp2=s[h:]
        s=tmp2+tmp1
    print(s)