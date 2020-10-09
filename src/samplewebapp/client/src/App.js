import React from "react";
import "./App.css";

// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";

// import the ToDoList component
import SubscriptionList from "./SubscriptionList";

import Amplify from 'aws-amplify';
import awsmobile from './aws-exports';
import { withAuthenticator } from 'aws-amplify-react';
Amplify.Logger.LOG_LEVEL = 'DEBUG';
Amplify.configure(awsmobile);
function App() {
  return (
    <div>
      <Container>
      </Container>
    </div>
  );
}
//export default App;
export default withAuthenticator(App, true);
