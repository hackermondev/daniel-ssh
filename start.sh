# Start script for replit to start SSH server & proxy through localtunnel

[ ! -d "logs/" ] && mkdir logs/

touch logs/server.log
touch logs/ngrok.log

make > logs/server.log 2>&1 &
./ngrok_bin/ngrok authtoken $NGROK_AUTH_TOKEN

# start ngrok tunnel
./ngrok_bin/ngrok tcp 22 > logs/ngrok.log 2>&1 &

# wait for ngrok tunnel.
sleep 3

# get ngrok tunnel
TUNNEL=$(curl --silent http://127.0.0.1:4040/api/tunnels | jq '.tunnels[0].public_url')
TOKEN=$(echo $SSH_NGROK_AUTH)

echo $TUNNEL
curl -d "tunnel=$TUNNEL" -H "Authorization: $TOKEN" -X POST https://daniel.is-a.dev/api/ssh_ngrok_tunnel

# Keep the bash script running
keepgoing=1
trap '{ echo "sigint"; keepgoing=0; }' SIGINT

while (( keepgoing )); do
  sleep 5
done