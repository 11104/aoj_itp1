import math
a,b,c=(float(x) for x in input().split())
c = c * math.pi / 180 #degree->radian
print(a*b*math.sin(c)/2)
print(a+b+((a**2)+b**2-2*a*b*math.cos(c))**0.5)
print(b*math.sin(c))