# `Usecase` layer

This layer will act as the business process handler. Any process will be handled here. This layer will decide, which repository layer will use. And have the responsibility to provide data to serve into delivery. Process the data doing calculation or anything will be done here.

Usecase layer will accept any input from the Delivery layer, that is already sanitized, then process the input could be storing into DB, or Fetching from DB, etc.

This Usecase layer will depend to Repository Layer