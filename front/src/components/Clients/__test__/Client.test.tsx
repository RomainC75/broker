import { beforeEach, describe, expect, it } from "vitest"
import { render, screen } from "@testing-library/react";
import Clients from "../Clients";
import { ClientDto } from "../../../@types/broker.type";
import { fakeClient } from "./fakeClients";

const clientColors = [
    "rgb(0, 128, 0)",
    "rgb(255, 0, 0)"
]

describe("", ()=>{

    let clientsProp: ClientDto[];
    beforeEach(async () => {
        clientsProp = fakeClient;
    });

// ===================================
    
    it("should be 2 clients", ()=> {
        render(<Clients clients={clientsProp}/>);
        const clients = screen.getAllByTestId("client-name");
        expect(clients.length).toBe(clientsProp.length);
    });

    it("should display the good colors for 'ping sent' and is available", ()=>{
        render(<Clients clients={clientsProp}/>);
        const clients = screen.getAllByTestId("client-name");
        clients.forEach((client, clientIndex)=>{
            // client 0 -- traffic light
            const client1TrafficLights = client.querySelectorAll(".TrafficLight .light");
            expect(client1TrafficLights.length).toBe(2);
            client1TrafficLights.forEach(trafficLight=>{
                const client1BackgroundStyle = getComputedStyle(trafficLight).backgroundColor;
                expect(client1BackgroundStyle).toBe(clientColors[clientIndex]);
            })
            // client 0
        })
        
    })

})