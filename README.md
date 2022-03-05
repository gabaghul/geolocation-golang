# **geolocation-golang**
Just a project to use as exercise for unit testing and API calls

This project runs the following API:

[GeoDB Cities RapidAPI](https://rapidapi.com/wirefreethought/api/geodb-cities/)

## **Objective**

This homework is focused in change an actual running code to create mocks using structs and interfaces from go, then use these mocks to create a unit test for services package with 100% coverage or close to it. 

So **let's code!!**

## **How to run this project**

You just need **GO** 1.17 or higher version to run this project. 

There is a `.env_copy` in this repo, rename it locally into your directory and fill the two necessary variables used by the service.

Then you can use your favorite browser / API Platform to do the requests. There are only two routes:

`{baseHost}/` - Just a message to test if the API is alive.

`{baseHost}/nearbyCities/{cityName}` - Returns nearby cities given a city name as reference, and also the population nearby this city.

## **Glossary**

`baseHost` - Host where your service is running. (eg: `http://localhost:8080`)

`cityName` - City name used as reference for the `nearbyCities` route. (eg: `osasco`) 

**Obs.:** will return two cities :open_mouth:

### **for .env file**

`PORT`: Port which your service will run. (eg: 8080)

`GEOAPI_API_KEY`: API Key needed to authenticate into RapidAPI and consume GeoDB routes.

## **IMPORTANT NOTES**

Actually is just one note hehe.

The webservice is taking too much time to respond intentionally. Because i used a free subcription from RapidAPI, which enable only 1 request per second. So I added some `time.Sleep()` to doesn't launch any `429 - Too many requests`