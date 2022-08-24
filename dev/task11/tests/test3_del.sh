#deletes first 3 events

curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 0,"user_id": "2", "date": "2019-09-11"}' http://localhost:8080/delete_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 1,"user_id": "2", "date": "2019-09-11"}' http://localhost:8080/delete_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 2,"user_id": "2", "date": "2019-09-11"}' http://localhost:8080/delete_event