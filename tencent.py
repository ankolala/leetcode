str = 'abcefeddf**adk'
leng = len(str)
item = list(str)
flag = 1
while flag == 1 and  leng != 0:
	flag = 0
	for i in range(leng-3):
		if item[i] != item[i+1] and item[i] == item[i+2]:
			del item[i]
			flag = 1
			leng = leng - 1
for each in item:
	print(each, end="")