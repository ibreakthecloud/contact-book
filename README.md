# Contact Book

Contact book is an go based contact server that allows to create contact with name and email, where email is unique.

Note:
The default username and password for the basic auth for authorization is:

Username: `admin`
Password: `password` 


- To Create contact:

Method: POST
API: `/contact`
```
curl --request POST 'https://contact-book-as-a-service.herokuapp.com/contact' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "k@rn",
    "name" : "karn"
}'
```

- To Delete contact:

Method: DELETE
API: `/contact`
```
curl --request DELETE 'https://contact-book-as-a-service.herokuapp.com/contact' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "k@rn",
}'
```
- To Edit contact:
Every contact has unique email address, so the only editable thing is name,
sp to edit contact name:

Method: PUT
API: `/contact`
```
curl --request POST 'https://contact-book-as-a-service.herokuapp.com/contact' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "k@rn",
    "name" : "karn H"
}'
```
- To Get all contacts:
Responses are paginated. 
There are 10 items by default per invocation, and default page is first.
In order to change the page, pass the page number in `page` parameter.
example: `?page=2`

Method: GET
API: `/contact`
```
curl --request GET 'localhost:9000/contact' 
```

 
- To Search User
    - With Name:
    *Method*: POST
    *API*: `/contact?name=karn`
     ```
     curl --request GET 'localhost:9000/contact?name=karn'
     ```
  
   - With Email
    *Method*: POST
    *API*: `/contact?name=karn`
     ```
     curl --request GET 'localhost:9000/contact?email=k@rn'
     ```
     
