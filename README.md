# geolocation-golang
Just a project to use as exercise for unit testing and API calls

This project runs the following API:

[GeoDB Cities RapidAPI](https://rapidapi.com/wirefreethought/api/geodb-cities/)

## **Objective**

This homework is focused in change an actual running code to create mocks using structs and interfaces from go, then use these mocks to create a unit test for services package with 100% coverage or close to it. 

So **let's code!!**

## **How to run this project**

You just need **GO** 1.17 or higher version to run this project. 

There is a `.env_copy` in this repo, rename it locally into your directory and fill the two necessary variables used by the service.

## **Glossary for .env file**

`PORT`: Port which your service will run. (eg: 8080)

`GEOAPI_API_KEY`: API Key needed to authenticate into RapidAPI and consume GeoDB routes.