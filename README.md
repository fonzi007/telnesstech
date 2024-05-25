# Telness Tech

The project is developed and run on IntelliJ IDEA.

## How to run locally
### Backend
Backend is a REST API Golang service using gin-gonic.

Requirements to run backend part locally: Docker installed.

Build the app:
```
$ make build
```

Start docker and run this command to set up the database:
```
$ make local-deps-up
```
Run the app:
```
$ make run
```

Start the server by clicking on "Run" button on main.go file

Our server is now running on `localhost:8080` 
### Frontend
Frontend is a React Javascript App. I used facebook pre-configured React App (ref: https://github.com/facebook/create-react-app/blob/main/README.md)

Requirements to run frontend part locally: npm installed.

```
$ cd frontend
$ npm install
$ cd my-app
$ npm start
```
`localhost:3000` will now display our web page.

## Improvements

Considering the time restriction, I didn't have time to implement everything. 

I made a very simple front end page which is made of a form with different field to submit, a click button to submit and send request. We should offer a better user experience by having a search bar taking msisdn as text input only, then tick filters we want. 

We should also add validation on the backend to validate user input.
We should also handle pagination.