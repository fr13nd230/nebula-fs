import {
  BadRequestException,
  ConflictException,
  HttpStatus,
  Injectable,
  Logger,
  NotFoundException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { hash, verify } from 'argon2';
import { Response } from '../../config/response';
import { User } from '../../repository/schema/user.schema';
import { CreateUserDto } from './dto/create-user.dto';
import { SignInDto } from './dto/sign-in.dto';
import { NotFoundError } from 'rxjs';

@Injectable()
export class AuthenticationService {
  private readonly logger = new Logger(AuthenticationService.name, {
    timestamp: true,
  });

  constructor(
    private readonly jwtService: JwtService,
    @InjectModel(User.name) private readonly userModel: Model<User>,
  ) {}

  async signUp(createUserDto: CreateUserDto): Promise<Response<User>> {
    try {
      const { email, password, fullName } = createUserDto;
      const user = await this.userModel.findOne({ email });

      if (user)
        throw new ConflictException(
          `User with email: ${email} already exists.`,
        );

      const hashedPassword: string = await hash(password);
      const newUser = await this.userModel.create({
        email,
        password: hashedPassword,
        fullName,
      });
      await newUser.save();

      return {
        status: true,
        code: HttpStatus.CREATED,
        message: 'User has been registered successfully',
        content: newUser,
      };
    } catch (err) {
      this.logger.error('Sign up process has failed.', err);
      throw err;
    }
  }

  async signIn(
    signInDto: SignInDto,
  ): Promise<Response<Partial<User> & { access_token: string }>> {
    try {
      const { email, password } = signInDto;

      const user = await this.userModel.findOne({ email });

      if (!user)
        throw new NotFoundException(`User with email: ${email} do not exist.`);

      const passwordMatch: boolean = await verify(user.password, password);

      if (!passwordMatch)
        throw new BadRequestException(
          'Invalid provided password, try again later.',
        );

      const payload = {
        id: user._id,
        email: user.email,
      };
      const access_token: string = this.jwtService.sign(payload);

      return {
        status: true,
        code: HttpStatus.OK,
        message: 'User signed in successfully.',
        content: {
          id: user._id,
          email: user.email,
          fullName: user.fullName ?? 'anonymous',
          access_token,
          created_at: user.created_at,
          updated_at: user.updated_at,
        },
      };
    } catch (err) {
      this.logger.error('Sign in process has failed.', err);
      throw err;
    }
  }
}
