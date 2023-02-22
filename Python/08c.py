#ord()...trasfer to Unicode
cnt = [0 for i in range(26)]
while True:
      try:
            s = input()#multi len input
      except:
            break
      for i in range(len(s)):
            m = ord(s[i])-ord('A')
            if(m >= 0 and m < 26):
                  cnt[m]+=1
            m = ord(s[i])-ord('a')
            if(m >= 0 and m < 26):
                  cnt[m]+=1
for i in range(26):
      print(chr(i+ord('a')),":",cnt[i])
