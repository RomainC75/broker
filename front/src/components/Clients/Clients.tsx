import React from "react";
import { ClientDto } from "../../@types/broker.type";
import TrafficLight from "../TrafficLight/TrafficLight";
import './client.scss'

export interface IClients {
  clients: ClientDto[];
}

const Clients = ({ clients }: IClients) => {
  return (
    <div className="Clients">
      {clients.map((client, index) => (
        <div key={`client-${index}`} className="client" data-testid="client-name">
          <p>{index}</p>
          <div>
            <div>
              <p>ping sent : </p>
              <TrafficLight boolValue={client.ping.is_ping_sent} />
            </div>
            <div>
              <p>is available : </p>
              <TrafficLight boolValue={client.is_available} />
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default Clients;
