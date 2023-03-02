# `Delivery` layer

This layer will act as the presenter. Decide how the data will be presented. Could be as REST API, HTML File, or gRPC whatever the delivery type.

The client will call the resource endpoint over the network, and the Delivery layer will get the input or request, and send it to Usecase Layer.

This template's using REST API as the delivery method. As a framework using [`Echo`](https://echo.labstack.com/). This is why you can see dedicated directory for it (http/echo). It's easy to change framework/router with this construction.

