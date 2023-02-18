import string

while True:
    a,op,b=input().split()
    if op=="?":
        break
    a,b=int(a),int(b)
    #Able to typedef after input!
    if op=='+':
        print(a+b)
    elif op=='-':
        print(a-b)
    elif op=='*':
        print(a*b)
    else:
        print(a//b)