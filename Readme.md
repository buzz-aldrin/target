####Target
Product microservice implementation<br/>

####Instructions to install
1. Follow `https://golang.org/doc/install` to install latest version of golang<br/>
2. Put target folder inside `$GOPATH/src/`<br/>

####Overview
Product microservice provides CRUD api's over product<br/>

####Design
Identified objects<br/>
1. CurrentPrice have 2 properties<br/>
    a. Value: type *Value<br/>
    b. CurrencyCode: type CurrencyCode<br/>
2. Product have 2 properties<br/>
    a. ID: type ProductID<br/>
   	b. CurrentPrice: type CurrentPrice<br/>
3. ProductDesc have 2 properties<br/>
    a. ID: type ProductID<br/>
   	b. Desc: type string<br/>
4. productResp have 2 properties<br/>
   	a. *models.Product: embedded Product type pointer<br/>
   	b. Desc:type string<br/>
    
####Supported Endpoints
1. GET /product/{id}: accepts product id as part of url. Return json representation of productResp.<br/>
2. PUT /product/{id}: expects Product as request body. Creates a new object if product with passed id<br/>
    is not present else updates existing product document. Return json representation of updated product<br/>
3. DELETE /product/{id}: deletes product identified by passed id. Return json representation of deleted product<br/>

####Setup: 
cd to project directory `$GOPATH/src/onefc`<br/>
run `$ ./setup`<br/>
Setup script<br/>
1. Checks if required version of golang(1.10.3) is installed<br/>
2. Runs unit tests<br/>
    a. displays results on command line<br/>
    b. stores the unit test case results in *_coverage.out<br/>
    c. run `$ go tool cover -html=*_coverage.out` to see html view of unit test cases<br/>
3. Builds target binary and stores in `$GOPATH/bin`<br/>

####Run
`$ $GOPATH/bin/target`
