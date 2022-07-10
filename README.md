# Buying_Frenzy
Backend service and a database for a food delivery platform 


### Step to start local server:

Step 1 : Build dockerfile 
       - cmd : docker build -t buying-frenzy:latest .

Step 2 : Run docker-compose
       - cmd : docker-compose up -d
     
- Local server have two command : 
  - go run main.go serve : to start the server 
  
  - go run main.go load  : to load the json file

Note : go main.go serve need to run first cause it will run all migrate table (for manual case) otherwise the json will not load correctly


Document API with swagger at : http://localhost:8080/swagger/index.html#/

Local api at : http://localhost:8080/api/v1

## Remote server 
Document API with swagger at: https://lit-temple-98149.herokuapp.com/swagger/index.html#/

API at  : https://lit-temple-98149.herokuapp.com/api/v1 (heroku deploy)


### usage for make file :

Run make all to run the project

Make build will build the binary file

Make run will start the server

Make load data with import data from json file to db

Make swagger to update the api document

Make migrate-up to migrate db

Make test will run all the test case

Make wire will auto inject all define provider
