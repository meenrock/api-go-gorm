# api-go-gorm
### Features

- REST API Server
- GRPC Server & Client
- Database Connector
	- MySQL Connector
	- PostgreSQL Connector

# api-go-gorm



####Install Dependencies

`$ `

####Data and Definition
```

```

                
###Operation Flowchart

```flow
st=>start: Login
op=>operation: Login operation
cond=>condition: Successful Yes or No?
e=>end: To admin

st->op->cond
cond(yes)->e
cond(no)->op
```

###Sequence Diagram
                    
```seq
Server->Client: Says Hello 
Note right of Client: Client thinks\nabout it 
Client-->Server: How are you? 
Client->>Server: I am good thanks!
```

