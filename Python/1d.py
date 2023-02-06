d=int(input())
h=int(d/3600)
d%=3600
m=int(d/60)
s=d%60
print(h,m,s,sep=':')