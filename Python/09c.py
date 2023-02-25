n=int(input())
pa=0
pb=0
for i in range(n):
    a,b=str(input()).split()
    if a==b:
        pa+=1
        pb+=1
    elif len(a)>=len(b):
        for j in range(len(b)):
            if ord(a[j])<ord(b[j]):
                pb+=3
                break
            elif ord(a[j])>ord(b[j]):
                pa+=3
                break
            elif ord(a[j])==ord(b[j]) and j==len(b)-1:
                pa+=3
                break
    else:
        for j in range(len(a)):
            if ord(a[j])<ord(b[j]):
                pb+=3
                break
            elif ord(a[j])>ord(b[j]):
                pa+=3
                break
            elif ord(a[j])==ord(b[j]) and j==len(a)-1:
                pb+=3
                break
print(pa,pb)

#easy
'''#According to Py specification,py can compare str by ascii number
N=int(input())
x=0
y=0
for i in range(N):
    a,b=input().split()
    if a<b:
        y+=3
    elif a==b:
        x+=1
        y+=1
    else:
        x+=3
print(x,y)
'''