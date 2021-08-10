This is Api service with comments and posts CRUD's and user registration with email or sign in through facebook or google.

Service uses echo minimalistic web framework. 
Service use swagger (echoSwagger) to visualize and interact with the API’s resources.
Service use jwt-token and oauth2 for security.

To download the project use "git clone https://github.com/Shevaister/Api" 

You need to have mysql or any database that could be connected with gorm drivers, if you want to change database from mysql you should go to pkg/db/connection.go and change imported driver and connection.

To set up web service you should change .env variables for yours, you can find them in cmd/main/.env.

Default starting page with swagger "http://localhost:8000/swagger/index.html"
