import React from "react";
import { TopicDto } from "../../@types/back_types";
import Content from "../Content/Content";
import Clients from "../Clients/Clients";

interface ITopic {
  topic: TopicDto;
}

const Topic = ({ topic }: ITopic) => {
  return (
    <div className="Topic">
      <Clients clients={topic.consumer_clients}/>
      <div className="index">
        <p>Index : </p>
        <p>{topic.reader_index}</p>
      </div>
      <Content messages={topic.content} />
    </div>
  );
};

export default Topic;
