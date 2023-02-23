# standard
s=input()
p=input()
for i in range(len(s)):
    for j in range(len(p)):
        if s[(i+j)%len(s)] != p[j]:
            break
        elif j == len(p)-1:
            print("Yes")
            exit()
print("No")

# shorter
'''
s=input()*2
p=input()
if p in s: print("Yes")
else: print("No")
'''