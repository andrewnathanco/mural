# vars
green_dep_file=mural.stack.yaml
blue_dep_file=mural-blue.stack.yaml
public_endpoint=https://mural.andrewnathan.net
nginx_conf_file="/etc/nginx/nginx.conf"

# get random port
# Picks a random number between 3000 and 3999.
function random-number {
  floor=3000
  range=3999
  number=0
  while [ "$number" -le $floor ]
  do
    number=$RANDOM
    let "number %= $range"
  done
  echo $number
}

echo "Getting a new random port for the green switch"
new_port=$(random-number)
while [[ $(lsof -i -P -n | grep :$new_port) ]]
do
  new_port=$(random-number)
done

# now that we have our port lets sub in that port to the stack
echo "Setting new env variables"
sed -i.bak "s/\$PORT/$new_port/g" $green_dep_file

# okay now we should be able to spin up our green env
echo "Running green switch"
docker compose -f $green_dep_file up -d
green_logs=$(docker compose -f $green_dep_file logs -t)

echo "Checking green deployment"
if [[ $green_logs == *"panic: 1"* ]]; then
    echo "Green deployment failed."
    echo $green_logs
    echo "Spinning down green deployment"
    docker compose -f $green_dep_file down
    exit 1
fi

# get the old port
full_port_mapping=$(yq eval '.services.[].ports[0]' $blue_dep_file)
echo $full_port_mapping
old_port=$(echo "$full_port_mapping" | awk -F ':' '{print $1}')
echo $old_port

# modify the nginx conf
echo "Updating nginx config"
sed -i.bak "s/proxy_pass http:\/\/localhost:${old_port}/proxy_pass http:\/\/localhost:${new_port}/" $nginx_conf_file
title=$(curl -s $public_endpoint | grep -o '<title>.*</title>' | awk -F '<title>' '{print $2}' | awk -F '</title>' '{print $1}')
sudo systemctl restart nginx

echo "Checking public endpoint"
if [ "$title" == "Mural" ]; then
    echo "Website is not being server resetting back."
    rm $nginx_conf_file
    cp $nginx_conf_file.bak $nginx_conf_file

    sudo systemctl restart nginx
    exit 1
fi

echo "Decommisioning blue switch"
docker compose -f $blue_dep_file down