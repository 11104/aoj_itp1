s=input()
n=int(input())
for i in range(n):
    op=list(input().split())
    a=int(op[1])
    b=int(op[2])
    if op[0]=="replace":
        q=op[3]
        s=s[:a]+q+s[b+1:]
    elif op[0]=="reverse":
        l=s[:a]
        m=s[a:b+1]
        r=s[b+1:]
        m=m[::-1]#-1 -> reverse
        s=l+m+r
    else:
        print(s[a:b+1])