import React from "react";
import { ClientDto } from "../../@types/back_types";
import TrafficLight from "../TrafficLight/TrafficLight";

interface IClients {
  clients: ClientDto[];
}

const Clients = ({ clients }: IClients) => {
  return (
    <div className="Clients">
      {clients.map((client, index) => (
        <div key={`client-${index}`} className="client">
          <TrafficLight name="ping_sent" isGood={client.ping.is_ping_sent} />
          <TrafficLight name="is_available" isGood={client.is_available} />
        </div>
      ))}
    </div>
  );
};

export default Clients;
