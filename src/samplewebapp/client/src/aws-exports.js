const awsmobile = {
    Auth: {
      // Amazon Cognito Region
      region: "us-east-1",
  
      // Amazon Cognito User Pool ID
      userPoolId: "us-east-1_njGiiHrUB",
  
      // Amazon Cognito Web Client ID (26-char alphanumeric string)
      userPoolWebClientId: "ft374gsb7mlau9nhbqqnrjvk6",
      
      oauth: {
        domain: "vikas-test-tool.auth.us-east-1.amazoncognito.com",
        scope:[
        'phone',
        'email',
        'openid',
        'profile',
        'aws.cognito.signin.user.admin'
        ],
        redirectSignIn: "http://localhost:3000/",
        redirectSignOut: "http://localhost:3000/signout.js",
        responseType: "code"
        }

    }
  };
  
  export default awsmobile;