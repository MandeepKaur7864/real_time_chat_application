import React, { Component } from 'react';
import Message from '../Message/Message';

class ChatHistory extends Component {
    render() {
        console.log(this.props.chatHistory)
        // const messages = this.props.chatHistory.map((msg, index) => (
        //     <p key={index}>{msg.data}</p>
        // ));

        const messages = this.props.chatHistory.map(msg => <Message message={msg.data} />)

        return (
            <div className="ChatHistory">
                <h2>ChatHistory</h2>
                {messages}
            </div>
        )
    }
}

export default ChatHistory;