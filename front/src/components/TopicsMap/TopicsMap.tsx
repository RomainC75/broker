import { useContext } from 'react'
import { SocketContext } from '../../context/socket.context';
import { ISocketContext } from '../../@types/socketContext.type';


const TopicsMap = () => {
  const {data} = useContext(SocketContext) as ISocketContext;

  return (
    <div>
    {data && Object.keys(data).map((topicName)=>
      <div>
        <p>{topicName}</p>
      </div>
    )}
    </div>
  )
}

export default TopicsMap