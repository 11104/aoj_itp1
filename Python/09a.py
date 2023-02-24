c=str(input()).lower()
s=[]
while True:
    tmp=input()
    if tmp=="END_OF_TEXT":
        break
    s+=tmp.lower().split()
count=0
for tmp in s:
    if tmp == c:
        count += 1
print(count)

#Using count()
'''
W = input().lower()
ans = 0
while 1:
    T = input()
    if T == "END_OF_TEXT":
        break
    ans += T.lower().split().count(W)
print(ans)
'''