while True:
    h,w=(int (x) for x in input().split())
    if h==0 and w==0:
        break
    for i in range(h):
        for j in range(w):
            print("#",end='')
        print()
    print()
