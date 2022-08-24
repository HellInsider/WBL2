# creates 3 users {1,2,3} in day 2019-09-09
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "usr1_msg1"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "usr2_msg1"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "usr2_msg2"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "3", "date": "2019-09-09", "title": "usr3_msg1"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "3", "date": "2019-09-09", "title": "usr3_msg2"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "3", "date": "2019-09-09", "title": "usr3_msg3"}' http://localhost:8080/create_event

# creates 3 users {1,2,3} in days {2019-09-10, 2019-09-11, 2019-09-12} 
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-10", "title": "usr1_msg2"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-11", "title": "usr2_msg3"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "3", "date": "2019-09-12", "title": "usr3_msg4"}' http://localhost:8080/create_event

# creates 3 users {1,2,3} in days {2019-09-01, 2019-09-19, 2019-09-29}
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-01", "title": "usr1_msg3"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-19", "title": "usr2_msg4"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "3", "date": "2019-09-29", "title": "usr3_msg5"}' http://localhost:8080/create_event