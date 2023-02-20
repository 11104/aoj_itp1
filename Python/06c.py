n=int(input())
map1 = [[0 for i in range(10)] for j in range(3)]
map2 = [[0 for i in range(10)] for j in range(3)]
map3 = [[0 for i in range(10)] for j in range(3)]
map4 = [[0 for i in range(10)] for j in range(3)]
for i in range(n):
    a,b,c,d=(int (x) for x in input().split())
    if a==1:
        map1[b-1][c-1]+=d
    elif a==2:
        map2[b-1][c-1]+=d
    elif a==3:
        map3[b-1][c-1]+=d
    else:
        map4[b-1][c-1]+=d

for i in range(4):
    if i==3:
        for j in range(20):
            print("#",end='')
        print()
    else:
        for j in range(10):
            print("",map1[i][j],end='')
        print()
for i in range(4):
    if i==3:
        for j in range(20):
            print("#",end='')
        print()
    else:
        for j in range(10):
            print("",map2[i][j],end='')
        print()
for i in range(4):
    if i==3:
        for j in range(20):
            print("#",end='')
        print()
    else:
        for j in range(10):
            print("",map3[i][j],end='')
        print()
for i in range(3):
    for j in range(10):
        print("",map4[i][j],end='')
    print()