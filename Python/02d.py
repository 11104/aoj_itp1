w, h, x, y, r =(int(x) for x in input().split())
if (0 <= x-r) and (x+r <= h) and (0 <= y-r) and (y+r <= h):
	print("Yes")
else:
	print("No")
