import { Request, Response } from "express";
import {HealthApi, HealthCheck, ResponseError} from "../clients/service-b";
import {ServiceBConfig} from "../config/config";

const healthCheck = async (req: Request, res: Response) => {
    const serviceBApi = new HealthApi(ServiceBConfig(req.headers))
    let statusCode = 200;
    let result = {}

    try{
        result = await serviceBApi.healthCheck()
    }catch(error: any){
        statusCode = error.status;
    }

    return res.status(statusCode).send(result);
}

module.exports = {
    HealthCheck: healthCheck,
};
