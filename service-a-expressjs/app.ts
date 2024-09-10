import express, { Request,  Response, NextFunction, Express} from "express";
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
        validateApiSpec: false,
        validateRequests: false,
        validateResponses: false,
        validateSecurity: {
            handlers: {
                ApiKey: function (req, scopes, schema){
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

app.use(apiKeyValidator);

http.createServer(app).listen(port);

console.log(`Running on port ${port}`);
