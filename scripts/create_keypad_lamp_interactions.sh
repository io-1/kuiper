!#/bin/bash

HOSTNAME=t1

http -v --json POST http://$HOSTNAME:9090/api/v1/interactions name="keypad and lamp interaction test" description="keypad and lamp interaction test" | jq '.id' | remove-quotes | { read INTERACTIONID }

http -v --json POST http://$HOSTNAME:9090/api/v1/keypad/condition mac=84cca8b322e4 buttonID:=1 | jq '.id' | remove-quotes | { read CONDITIONID1 }

http -v --json POST http://$HOSTNAME:9090/api/v1/keypad/condition mac=84cca8b322e4 buttonID:=2 | jq '.id' | remove-quotes | { read CONDITIONID2 }

http -v --json POST http://$HOSTNAME:9090/api/v1/lamp/event/pulse mac=3c6105171f4c red:=170 green:=0 blue:=70 | jq '.id' | remove-quotes | { read EVENTID1 }

http -v --json POST http://$HOSTNAME:9090/api/v1/lamp/event/toggle mac=3c6105171f4c | jq '.id' | remove-quotes | { read EVENTID2 }

http -v --json POST http://$HOSTNAME:9090/api/v1/interact/keypad/lamp interactionID=$INTERACTIONID conditionID=$CONDITIONID1 eventID=$EVENTID1

http -v --json POST http://$HOSTNAME:9090/api/v1/interact/keypad/lamp interactionID=$INTERACTIONID conditionID=$CONDITIONID2 eventID=$EVENTID2

http -v --json GET http://$HOSTNAME:9090/api/v1/interactions/$INTERACTIONID/details
