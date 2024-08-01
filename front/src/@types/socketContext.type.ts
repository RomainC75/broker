import { TopicMapDto } from "./back_types";

export interface ISocketContext {
  myState: number;
  data: TopicMapDto | null;
  setData: (data: TopicMapDto) => void;
}
