#moves first 3 events from 2019-09-09 to 2019-09-10, changes title and changes user in 3-rd even from 3 to 1

curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 0, "user_id": "1", "date": "2019-09-10", "title": "updatedTitle_usr1_msg1"}' http://localhost:8080/update_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 1, "user_id": "2", "date": "2019-09-10", "title": "updatedTitle_usr3_msg1"}' http://localhost:8080/update_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"id": 2, "user_id": "1", "date": "2019-09-10", "title": "updatedTitle_usr3_msg1"}' http://localhost:8080/update_event

