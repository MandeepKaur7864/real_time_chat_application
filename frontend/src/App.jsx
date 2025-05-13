import React, { Component } from 'react'
import './App.css'
import { connect, sendMsg } from './api'
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {
  constructor(props) {
    super(props);
    // connect();
    this.state = {
      chatHistory: []
    }
  }

  componentDidMount() {
    connect((msg)=>{
      console.log("New message");
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send(event) {
    // console.log("hello");
    // sendMsg("hello");
    if(event.keyCode === 13) {
      sendMsg(event.target.value)
      event.target.value = ""
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        {/* <button onClick={this.send}>Hit</button> */}
        <ChatInput send={this.send} />
      </div>
    )
  }
}

export default App;
