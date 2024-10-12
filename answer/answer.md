# My Answer

<!--toc:start-->

- [My Answer](#my-answer)
  - [Section Number 1](#section-number-1)
    - [Question Number 1](#question-number-1)
    - [Question Number 2](#question-number-2)
    - [Question Number 3](#question-number-3)
  - [Section Number 2](#section-number-2)
  <!--toc:end-->

## Section Number 1

Disclaimer: I've asked about what "mechanism diagram" mean in email, but haven't got the answer until the time I write this, so I will interpret such as Flowchart.

### Question Number 1

TBA

<!--![alt text](path) -->

### Question Number 2

TBA

<!--![alt text](path) -->

### Question Number 3

as per my limited knowledge, how real-time chat communication works is you have one or more users with a client, whether its a webapp or mobile app, then you have your database, and then you a backend server connecting all of them.

say we have 2 users in a chat room, when User A sends a message, the message will be sent to the backend server, which might also sends UUID of the user, unix timestamp, and maybe which uuid of the room. Then the server will store that message to the database to be linked with previous messages that maybe related if using relational database.

The server then sends the message to the User B, and then the client will sends another message to the server that the message has be received by the recepient which the message will be rendered in the client that Both users used, User A will get a little pop up that the message has been delivered and User B will get a message from user A.

the database schema will be different from project to project but at minimum it might have

- unix time of when the message was sent
- uuid of the user who sent the message
- uuid of the message room
- content of the message
- who reads the message (if any)

---

## Section Number 2
