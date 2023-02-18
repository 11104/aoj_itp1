n=int(input())
min,max,sum=1000001,-1000001,0
tmp=input().split()
for i in range(n):
    tmp[i]=int(tmp[i])
    if tmp[i]<min:
        min=tmp[i]
    if max<tmp[i]:
        max=tmp[i]
    sum+=tmp[i]
print(min,max,sum)

#map
"""
n = int(input())
data = list(map(int, input().split()))   #map()->some data in one line store into list
vmin = min(data)
vmax = max(data)
vsum = sum(data)
print(vmin, vmax, vsum)
"""