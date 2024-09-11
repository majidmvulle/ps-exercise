import { Request, Response } from "express";
import {HealthApi} from "../clients/service-b";
import {ServiceBConfig} from "../config/config";

const healthCheck = async (req: Request, res: Response) => {
    const serviceBApi = new HealthApi(ServiceBConfig(req.headers))
    let statusCode = 200;
    let result = {}

    try{
        result = await serviceBApi.healthCheck()
    //eslint-disable-next-line @typescript-eslint/no-explicit-any
  }catch(error: any){ //contradicts with: TS1196: Catch clause variable type annotation must be any or unknown if specified.
        statusCode = error.status;
    }

    return res.status(statusCode).send(result);
}

module.exports = {
    HealthCheck: healthCheck,
};
