import "./App.css";
// import useEffect hook
import React, { useEffect, useState , Component} from "react";
import { fetchUser } from "./auth";
import axios from "axios/index";



// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";

// import the ToDoList component
import SubscriptionList from "./SubscriptionList";

import Amplify from 'aws-amplify';
import awsmobile from './aws-exports';
import { withAuthenticator } from 'aws-amplify-react';
import { Auth, Hub,API } from 'aws-amplify'
import Axios from "axios";


Amplify.Logger.LOG_LEVEL = 'DEBUG';
Amplify.configure(awsmobile);

const apiName = "GetTask"
const path = "api/task"


async function GetTaskAPI() {

  // You may have saved off the JWT somewhere when the user logged in.
  // If not, get the token from aws-amplify:
  const user = await Auth.currentAuthenticatedUser();
  const token = user.signInUserSession.idToken.jwtToken;

  const request = {
      body: {
          attr: "value"
      },
      headers: {
          Authorization: token
      }
  };

//const data = {};
//  axios.get("http://localhost:8080/api/task", { crossdomain: true }, { headers: { Authorization: 'Bearer ${token}' }})
//  .then(response => {
//      console.log(response);
//      alert(token);
//      window.location.reload();
//  });


//config = { headers: { 'Authorization': 'Bearer ' + token } }

//return axios({ method: 'get', url: "http://localhost:8080" + '/api/task', 'crossdomain': true }, {headers: { 'Authorization': 'Bearer ' + token } }).then(response => {
  //      console.log(response);
   //     alert(token);
   //     window.location.reload();
   // });

   axios.get("http://localhost:8080/api/task", {
     crossdomain: true,
     headers: { 'Authorization': 'Bearer ' + token},
}).then(res => { 
    console.log(res);
}).catch(error => {
    console.log('error', error);
})


// axios.post(
//   "http://localhost:8080/api/task", {
//       headers: {
//           "Cache-Control": "no-cache",
//           "content-type": "application/vnd.api+json",
//           "x-api-key": "9xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx9",
//           "Access-Control-Allow-Origin": "*",
//           "Crossdomain": true,
//       },
    
//   }).then(response => {
//       console.log(response);
//   }).catch(error => {
//       console.log(error)
//   })

    return axios({
      method: 'post',
      crossdomain:true,
      // data:{
      //   Product: "product-1",
      //   Type: "type-1"
      // },
       url: "http://localhost:8080/api/task",
      headers: { 'Authorization': 'Bearer ' + token },
  })

}

function checkUser() {
  Auth.currentAuthenticatedUser()
    .then(user => console.log({ user }))
    .catch(err => console.log(err));
}

function signOut() {
  Auth.signOut()
    .then(data => console.log(data))
    .catch(err => console.log(err));
}

function App(props) {
  // in useEffect, we create the listener
  useEffect(() => {
    Hub.listen('auth', (data) => {
      const { payload } = data
      console.log('A new auth event has happened: ', data)
       if (payload.event === 'signIn') {
         console.log('a user has signed in!')
       }
       if (payload.event === 'signOut') {
         console.log('a user has signed out!')
       }
    })
  }, [])
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Sample Full Stack App.
        </p>
        <button onClick={checkUser}>Check User</button>
        <button onClick={signOut}>Sign Out</button>
        <button onClick={GetTaskAPI}>Get Task</button>

      </header>
    </div>
  );
}

//export default App;
export default withAuthenticator(App, true);