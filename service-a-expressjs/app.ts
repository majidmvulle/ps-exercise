import express, {Request, Response, NextFunction, Express} from "express";
import dotenv from "dotenv";
import * as OpenApiValidator from 'express-openapi-validator';
import path from "node:path";
import bodyParser from 'body-parser';
import http from 'http';

dotenv.config();

const app: Express = express();
const port = process.env.PORT || 3000;
const spec = path.join(__dirname, 'docs', 'openapi', 'api.bundled.yaml');

const apiKeyValidator = function (req: Request, res: Response, next: NextFunction) {

    next()
}

app.use(
    OpenApiValidator.middleware({
        apiSpec: spec,
        validateApiSpec: true,
        validateRequests: true,
        validateResponses: true,
        validateSecurity: {
            handlers: {
                ApiKey: function (req){
                    const apiKey = req.headers['x-api-key'];

                    if (apiKey != process.env.API_KEY) {
                        throw { status: 401, message: 'Unauthorized'}
                    }

                    return true;
                }
            }
        },
        operationHandlers: path.join(__dirname, 'routes'),
    })
);

app.use(bodyParser.json());

//err: any is part of the expressjs handler signature.
//eslint-disable-next-line @typescript-eslint/no-explicit-any
const jsonErrorHandler: (err: any, req: Request, res: Response, next: NextFunction) => void = (err: any, req: Request, res: Response, next: NextFunction) => {
  res.status(err.status).send({
    status: err.status,
    message: err.message,
  });
  return next();
};

app.use(jsonErrorHandler);

app.use(apiKeyValidator);

http.createServer(app).listen(port);

console.log(`Running on port ${port}`);
