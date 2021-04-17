echo "Calculate average";
echo -n "Enter number of scores : "
read TOTAL

sum=0

for (( i=1; i<=$TOTAL; i++ )); do
  echo -n "Enter number $i : "
  read num
  sum=$((sum + num))
done  
  
avg=$((sum / TOTAL))
echo "Average value $avg"