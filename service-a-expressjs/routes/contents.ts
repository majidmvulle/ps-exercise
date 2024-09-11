import { Request, Response } from "express";
import {
  Content,
  ContentsApi,
  FetchContentByIDRequest,
  ListContents200Response, ListContentsRequest,
  ResponseError
} from "../clients/service-b";
import {ServiceBConfig} from "../config/config";

const defaultPage = 1;
const defaultPerPage = 20;

const listContents = async (req: Request, res: Response) => {
  const serviceBApi = new ContentsApi(ServiceBConfig(req.headers))
  const request: ListContentsRequest = {
    page: req.query.page ? Number(req.query.page) : defaultPage,
    perPage: req.query.per_page ? Number(req.query.per_page) : defaultPerPage,
  };
  let statusCode = 200;
  let result: ListContents200Response|ResponseError|undefined;

  try{
    result = await serviceBApi.listContents(request)
    //eslint-disable-next-line @typescript-eslint/no-explicit-any
  }catch(error: any){ //contradicts with: TS1196: Catch clause variable type annotation must be any or unknown if specified.
    if (error instanceof ResponseError) {
      statusCode = error.response.status
      result = await error.response.json()
    }
  }

  console.log(result)

  return res.status(statusCode).send(result);
}

const fetchContentByID = async (req: Request, res: Response) => {
  const serviceBApi = new ContentsApi(ServiceBConfig(req.headers))
  const request: FetchContentByIDRequest = {id: Number(req.params.id)};

  let statusCode = 200;
  let result: Content|ResponseError|undefined;

  try{
    result = await serviceBApi.fetchContentByID(request)
  //eslint-disable-next-line @typescript-eslint/no-explicit-any
  }catch(error: any){ //contradicts with: TS1196: Catch clause variable type annotation must be any or unknown if specified.
    if (error instanceof ResponseError) {
      statusCode = error.response.status
      result = await error.response.json()
    }
  }


  return res.status(statusCode).send(result);
}

module.exports = {
  ListContents: listContents,
  FetchContentByID: fetchContentByID,
};
