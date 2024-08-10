import axiosInstance from "./interceptor"

export class AxiosReq{
    static getTicket = async() =>{
        console.log('========================')
        const res = await axiosInstance.get("/socket/ticket") 
        console.log("-> response ", res)
    }

}