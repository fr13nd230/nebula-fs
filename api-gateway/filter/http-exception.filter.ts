import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { Request, Response } from 'express';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const _req = ctx.getRequest<Request>();
    const res = ctx.getResponse<Response>();
    const code: number =
      exception.getStatus() ?? HttpStatus.INTERNAL_SERVER_ERROR;
    const message =
      (exception.getResponse() as any).message ??
      'Internal server error, try again later.';
    return res.status(code).json({
      status: false,
      code,
      message,
    });
  }
}
