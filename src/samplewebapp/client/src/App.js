import React from "react";
import "./App.css";

// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";

// import the ToDoList component
import SubscriptionList from "./SubscriptionList";

function App() {
  return (
    <div>
      <Container>
        <SubscriptionList />
      </Container>
    </div>
  );
}
export default App;