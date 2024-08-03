import React from "react";
import { TopicDto } from "../../@types/back_types";
import Content from "../Content/Content";

interface ITopic {
  topic: TopicDto;
}

const Topic = ({ topic }: ITopic) => {
  return (
    <div className="Topic">
      <Content messages={topic.content} />
    </div>
  );
};

export default Topic;
