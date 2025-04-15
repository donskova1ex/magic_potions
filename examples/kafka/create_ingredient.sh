#echo '{"name":"testing_recipe1","brew_time_seconds":1,"ingredients":[{"name":"ingredient1","quantity":4},{"name":"ingredient2","quantity":5}]}' | kcat -P -b 127.0.0.1:19092 -t create_recipes.v1
for (( i = 1; i < 100; i++ )); do
    n=$((i * 100))
    json='{"name":"testing_recipe'"$i"'","brew_time_seconds":'"$i"',"ingredients":[{"name":"ingredient'"$i"'","quantity":4},{"name":"ingredient'"$n"'","quantity":5}]}'
    echo "$json" | kcat -P -b 127.0.0.1:19092 -t create_recipes.v1
done