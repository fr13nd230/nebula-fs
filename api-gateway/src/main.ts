import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { HttpExceptionFilter } from 'filter/http-exception.filter';
import {
  ConsoleLogger,
  Logger,
  ValidationPipe,
  VersioningType,
} from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import helmet from 'helmet';

async function bootstrap() {
  const { PORT } = process.env;

  const appLogger: Logger = new Logger();
  const app = await NestFactory.create(AppModule, {
    logger: new ConsoleLogger({
      prefix: 'Gateway',
      timestamp: true,
    }),
  });

  app.enableCors({ origin: '*' });
  app.use(helmet());
  app.enableVersioning({
    type: VersioningType.URI,
  });
  app.useGlobalPipes(new ValidationPipe());
  app.useGlobalFilters(new HttpExceptionFilter());
  app.useLogger(appLogger);

  const swaggerConfig = new DocumentBuilder()
    .setTitle('Nebula FS API Gateway')
    .setDescription(
      `This is the API documentation and reference of the NebulaFS project, you will find only the relevant documenations of exposed APIs and authorization ones.`,
    )
    .setVersion('1.0.1')
    .addBearerAuth()
    .build();
  const documentFactory = () =>
    SwaggerModule.createDocument(app, swaggerConfig);
  SwaggerModule.setup('/', app, documentFactory);

  await app.listen(process.env.PORT ?? 3000, () =>
    appLogger.log(
      `Server is up and running at: http://localhost:${PORT ?? 3000}`,
    ),
  );
}

bootstrap().catch((err: Error) => console.error(err));
