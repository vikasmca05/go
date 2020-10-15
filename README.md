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


## User Groups feature of Cognito to manage access of Resource Server
- You can create and manage groups in a user pool from the AWS Management Console, the APIs, and the CLI. As a developer (using AWS credentials), you can create, read, update, delete, and list the groups for a user pool. You can also add users and remove users from groups.
- We are using groups in a user pool to control permission with Resource server written in GoLang. The groups that a user is a member of are included in the ID token provided by a user pool when a user signs in. 
- One or many gorups can be assigned to a user. Based on groups value in JWT, we can provide access of resources to user.
- Create User group
![Create User Group](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-14%20at%208.32.32%20PM.png)
- Assign user to a group
![Assign User to Group](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-14%20at%208.32.55%20PM.png)
- cognito:group array in JWT payload define the Group assosciation
![jwt.io](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-14%20at%208.31.29%20PM.png)
- cognito:groups value as admin in logs
![Admin value in console logs](https://github.com/vikasmca05/go/blob/master/Screen%20Shot%202020-10-14%20at%208.29.48%20PM.png)
