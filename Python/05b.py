while True:
    h,w=(int (x) for x in input().split())
    if h==0 and w==0:
        break
    for i in range(h):
        for j in range(w):
            if i==0 or i==h-1 or j==0 or j==w-1:
                print("#",end='')
            else:
                print(".",end='')
        print()
    print()
