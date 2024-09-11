import {Configuration, HTTPHeaders} from "../clients/service-b";
import {IncomingHttpHeaders} from "node:http";

export const ServiceBConfig = (incomingHttpHeaders?: IncomingHttpHeaders): Configuration => {
    let headers: HTTPHeaders = {}

    headers["accept"] = "application/json";

    if (incomingHttpHeaders && incomingHttpHeaders["Accept-Language"]) {
        headers["Accept-Language"] = incomingHttpHeaders["Accept-Language"][0];
    }

    return new Configuration({
        basePath: process.env.SERVICE_B_URL,
        apiKey: process.env.SERVICE_B_API_KEY,
        headers: headers,
    })
}
