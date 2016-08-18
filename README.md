# Ol - REST API Project

> Ol was created as part of a challenge to create a REST API to serve business data in the JSON format. It is intended to be built on a Linux system.
> This is the first project I've created in Go, and I had a great time learning the language's features
> and strengths while working in it. I look forward to future opportunities where I may work in Go again.

### Installation

Ol requires Go to run.
You will need to follow on of the guides to [install Go for Linux](https://golang.org/doc/install?download=go1.7.linux-amd64.tar.gz)


After you have Go installed, we'll want to make sure you have a few dependencies!
```sh
$ sudo apt-get install git
$ sudo apt-get install sqlite3
$ sudo apt-get install g++
```
Now we'll need to clone our repository into our Go workspace:
```sh
$ mkdir $GOPATH/src/github.com/lpszBuffer
$ cd $GOPATH/src/github.com/lpszBuffer
$ git clone https://github.com/lpszBuffer
```

Finally, we need to get ol's dependencies and install the application
```sh
$ go get github.com/lpszBuffer/ol
$ go install github.com/lpszBuffer/ol
```

### Execution
After installation, to run the ol server, use the following command in your terminal:
```sh
$ $GOPATH/bin/ol
```
You should see the following message, letting you know ol is ready to handle your requests!
```sh  
"Server listening on Port 8080..."
```

### Usage Examples
There are two endpoints available through ol.

#### /businesses
* Accepts pagination arguments page and per_page.
* If no arguments are provided, page 1 is displayed at 50 businesses per page.
* Returns a paginated JSON array of businesses, along with the pagination metadata.

##### Examples
Request:
``` http
http://localhost:8080/businesses
```

Response:
``` json
{
   "_metadata":{
      "page":1,
      "per_page":50
   },
   "businesses":[
      {
         "id":0,
         "uuid":"2859d6e0-1cb9-4fe9-bc00-97823a9fa4cb",
         "name":"Yundt-Flatley",
         "address":"1386 Lim Brooks",
         "address2":"Suite 517",
         "city":"Lake Betsy",
         "state":"IA",
         "zip":19416,
         "country":"US",
         "phone":4034880719,
         "website":"http://www.halvorson.com/",
         "created_at":"2012-12-10T16:17:58Z"
      }, ...
   ]
}
```

Request:
```http
http://localhost:8080/businesses?page=4;per_page=1
```
Response:
```json
{
   "_metadata":{
      "page":4,
      "per_page":1
   },
   "businesses":[
      {
         "id":3,
         "uuid":"5c2169f1-a319-469b-bc40-8c89623c686e",
         "name":"Kertzmann LLC",
         "address":"7042 Etna Mountains",
         "address2":"",
         "city":"Lake Tobi",
         "state":"KS",
         "zip":63064,
         "country":"US",
         "phone":4442359224,
         "website":"http://okeefe.info/",
         "created_at":"2012-12-16T20:18:08Z"
      }
   ]
}
```

#### /businesses/id
* Returns a specific business based on its id.
* Returns 400 Bad Request if id is not found
##### Example

Request:
```http
http://localhost:8080/businesses/7777
```

Response:
```json
{
   "id":7777,
   "uuid":"9c51ea80-6250-4094-a0e5-390827a5397e",
   "name":"Wiza-Spencer",
   "address":"33379 Wilkinson Trafficway Apt. 185",
   "address2":"Suite 504",
   "city":"Wintheiserton",
   "state":"GA",
   "zip":18290,
   "country":"US",
   "phone":4741831518,
   "website":"http://www.barrows.com/",
   "created_at":"2014-12-05T03:55:49Z"
}
```

### Attributions

Ol uses a couple of fantastic open source projects to work properly:
* [gorm](https://github.com/jinzhu/gorm) - ORM for the GoLang platform
* [mux](https://github.com/gorilla/mux) - Routing library for the GoLang platform
