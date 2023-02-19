n=int(input())
for i in range(1,n+1):
	if i%3 == 0 or i%10 == 3 or (i//10)%10 == 3 or (i//100)%10 == 3 or (i//1000)%10 == 3:
		print("",i,end='')
print()