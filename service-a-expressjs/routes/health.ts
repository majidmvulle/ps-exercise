import { Request, Response } from "express";

const healthCheck = (req: Request, res: Response) => {
  return res.status(200).send({
      status: "UP",
  });
}

module.exports = {
  HealthCheck: healthCheck,
};
