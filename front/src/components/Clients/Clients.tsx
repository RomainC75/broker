import React from "react";
import { ClientDto } from "../../@types/back_types";
import TrafficLight from "../TrafficLight/TrafficLight";
import './client.scss'

interface IClients {
  clients: ClientDto[];
}

const Clients = ({ clients }: IClients) => {
  return (
    <div className="Clients">
      {clients.map((client, index) => (
        <div key={`client-${index}`} className="client">
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
