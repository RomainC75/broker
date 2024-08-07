import React from "react";
import { EColumnNames } from "../../@types/column";
import TrafficLight from "../TrafficLight/TrafficLight";
import './column.scss'
import { MessageDto } from "../../@types/broker.type";

interface IColumn {
  columnName: EColumnNames;
  messages: MessageDto[];
}

const Column = ({ columnName, messages }: IColumn) => {
  return (
    <div className="Column">
      <div>{columnName}</div>
      {messages.map((message, index) =>
        typeof message[columnName] === "boolean" ? (
          <TrafficLight key={`column-${index}`} boolValue={message[columnName]} />
        ) : (
          <div key={`column-${index}`}>{message[columnName]}</div>
        )
      )}
    </div>
  );
};

export default Column;
