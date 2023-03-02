# `Repository` layer

The repository will store any Database handler. Querying, or Creating/ Inserting into any database will be stored here. This layer will act for CRUD to the database only. No business process happens here. Only plain function to Database.

This layer also has the responsibility to choose what DB will use in the Application. Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decide here.

Sometimes, yoy may need to calling other microservices. In this case it can be handled here (but it isn't required, because the way to implement this in Delivery is quite good too, when you present Delivery as communication/transport layer. So, it's holywar question). Create HTTP Requests for other services, and sanitize the data. This layer must fully act as a repository. Handle all data input-output no specific logic happen.

This Repository layer will depend on Connected DB, or other microservices if exists.