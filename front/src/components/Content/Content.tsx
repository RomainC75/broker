import Column from "../Column/Column";
import { EColumnNames } from "../../@types/column";
import "./content.scss";
import ContentColumn from "../Column/ContentColumn";
import { MessageDto } from "../../@types/broker.type";
interface IContent {
  // topicName: string;
  messages: MessageDto[];
}

const Content = ({ messages }: IContent) => {
  return (
    <div className="Content">
      <Column columnName={EColumnNames.index} messages={messages}/>
      <Column columnName={EColumnNames.key} messages={messages}/>
      {/* <Column columnName={EColumnNames.value} messages={messages}/> */}

      <ContentColumn messages={messages}/>
      <Column columnName={EColumnNames.isSent} messages={messages}/>
      <Column columnName={EColumnNames.isHandled} messages={messages}/>
    </div>
  );
};

export default Content;
