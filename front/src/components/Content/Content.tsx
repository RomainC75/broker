import React from "react";
import { MessageDto } from "../../@types/back_types";
import "./content.scss";
import TrafficLight from "../TrafficLight/TrafficLight";
interface IContent {
  // topicName: string;
  messages: MessageDto[];
}

const Content = ({ messages }: IContent) => {
  return (
    <div className="Content">
      {messages.map((message, index) => (
        <div key={`${index}`} className="Content_message">
          <div className="indicator">
            <p>key : </p>
            <div>{message.key}</div>
          </div>
          <div className="indicator">
            <p>value : </p>
            <div>{message.value}</div>
          </div>
          <TrafficLight name="is_sent" isGood={message.is_sent} />
          <TrafficLight name="is_handled" isGood={message.is_handled} />
        </div>
      ))}
    </div>
  );
};

export default Content;
