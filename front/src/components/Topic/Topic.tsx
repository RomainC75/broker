import React from 'react'
import { TopicDto } from '../../@types/back_types'
import Content from '../Content/Content';

interface ITopic{
    topic: TopicDto;
}

const Topic = ({topic}: ITopic) => {
  return (
    <div>
        <div>
            <p>content</p>
            {topic.content.map((cont, index)=>
                <Content topicName={`${index}`} messages={cont}/>
            )}
        </div>
    </div>
  )
}

export default Topic