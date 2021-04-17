print("Calculate average")
 
score = []
sum = 0
TOTAL = int(input("Enter number of scores :"))
for j in range(0, TOTAL):
   ele = int(input("Enter : "))
   score.append(ele)
 
for i in range(0, len(score)):
   sum = sum + score[i];   
 
avg=sum/TOTAL
print("Total number of scores", TOTAL)
print("Total Sum =", sum)
print("Average value =", avg)