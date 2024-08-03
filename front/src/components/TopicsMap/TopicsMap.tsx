import { useContext } from 'react'
import { SocketContext } from '../../context/socket.context';
import { ISocketContext } from '../../@types/socketContext.type';
import Topic from '../Topic/Topic';
import './topicsMap.scss'

const TopicsMap = () => {
  const {data} = useContext(SocketContext) as ISocketContext;

  return (
    <div className='TopicsMap'>
    {data && Object.keys(data).map((topicName, index)=>
      <div className='topics-iterator' key={`topic-${index}`}>
        <h3>{topicName}</h3>
        <Topic topic={data[topicName]}/>
      </div>
    )}
    </div>
  )
}

export default TopicsMap