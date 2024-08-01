import React from 'react'
import { MessageDto } from '../../@types/back_types'

interface IContent {
    topicName: string;
    messages: MessageDto[];
}

const Content = ({topicName, messages}: IContent) => {
  return (
    <div>
        {messages.map((message, index)=>
            <div key={`${topicName}-${index}`}>
                <p>key : {message.key}</p>
                <p>value : {message.value}</p>
                <p>is_sent : {message.is_sent}</p>
                <p>is_handled : {message.is_handled}</p>
            </div>
        )}
    </div>
  )
}

export default Content