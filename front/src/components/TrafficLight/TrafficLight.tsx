import React from 'react'
import './trafficLight.scss'

interface ITrafficLight{
    name: string;
    isGood: boolean;
}

const TrafficLight = ({name, isGood}: ITrafficLight) => {
  return (
    <div className='TrafficLight'>
        <p>{name} : </p>
        <div className={`light ${isGood ? "green":"red"}`}></div>
    </div>
  )
}

export default TrafficLight