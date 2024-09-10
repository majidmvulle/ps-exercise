import { Request, Response } from "express";

const listContents = (req: Request, res: Response) => {
  return res.status(200).send('ListContents');
}

const fetchContentByID = (req: Request, res: Response) => {
  return res.status(200).send('ContentByID');
}

export const fetchContentBySlug = (req: Request, res: Response) => {
  return res.status(200).send('ContentBySlug');
}

module.exports = {
  ListContents: listContents,
  FetchContentByID: fetchContentByID,
  FetchContentBySlug: fetchContentBySlug,
};
