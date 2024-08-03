import React from 'react'
import './trafficLight.scss'

interface ITrafficLight{
    boolValue: boolean;
}

const TrafficLight = ({boolValue}: ITrafficLight) => {
  return (
    <div className='TrafficLight'>
        <div className={`light ${boolValue ? "green":"red"}`}></div>
    </div>
  )
}

export default TrafficLight