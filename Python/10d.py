n=int(input())
a=list(input().split())
b=list(input().split())
d1=0.0
for i in range(n):
    d1+=abs(float(a[i])-float(b[i]))
print(d1)
d2=0.0
for i in range(n):
    d2+=abs(float(a[i])-float(b[i]))**2
print(d2**0.5)
d3=0.0
for i in range(n):
    d3+=abs(float(a[i])-float(b[i]))**3
print(d3**(1.0/3.0))
max=-0.1
for i in range(n):
    if max<abs(float(a[i])-float(b[i])):
        max=abs(float(a[i])-float(b[i]))
print(max)