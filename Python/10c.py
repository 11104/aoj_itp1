while True:
    n=int(input())
    if n==0:
        break
    s=list(input().split())
    m=0.0
    for i in range(n):
        m+=float(s[i])
    m=m/n   #average
    a=0.0
    for i in range(n):
        a+=(float(s[i])-m)**2
    print((a/n)**0.5)

