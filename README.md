# Sample Full Stack Application for AWS Cognito Integration

## Tech Stack
- React (frontend)
- Go (Backend REST APIs)
- AWS Cognito (Authentication)
- Amplify SDK (Javascript SDK for Cognito Integration)
- Axios (Http calls)


## Command to Run
### Start Frontend
- Navigate to src/samplewebapp/client
- npm start 
- Should start server in localhost:3000

### Start Backend
- Navigate to src/samplewebapp
- go run main.go
- Should start server in localhost:8080
- API supported localhost:8080/api/task
- Test API supported localhost:8080/hello (print Hello World)


## Test Feature
- Open developer tool in Chrome/ or other browser
- Capture Console and Network logs
- Home Page 
 ![Home Page](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.04.49%20PM.png)
- Click on SignIn with AWS button
- It should launch Cognito Hosted Sign In Page.
![Cognito Hosted Sign In Page](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.05.02%20PM.png)
- Enter the User credentials in Sign In Page. User should already be created in Cognito User Pool. Username - test4@test.com, Pwd - Welcome123!
![Sign In](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.05.17%20PM.png)
- If successfully signed in then user will see option to "Check User", "Sign Out", "Get Task"
- **Check User** : it will print the user details in Console for developer tools. **Sign Out** : It will sign out user from Cognito User pool. **Get Task** : It will call backend server with sample "Get" and "Post" call
- Check User : Console log
![Check User](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.26.30%20PM.png)
- Get Task : Console log
![Get Task](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.28.14%20PM.png)
- Get Task : Network log
![Get Task](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-12%20at%2010.28.46%20PM.png)
