import { Body, Controller, HttpCode, HttpStatus, Post } from '@nestjs/common';
import { AuthenticationService } from './authentication.service';
import { CreateUserDto } from './dto/create-user.dto';
import { SignInDto } from './dto/sign-in.dto';

@Controller({ version: '1', path: 'auth' })
export class AuthenticationController {
  constructor(private readonly authenticationService: AuthenticationService) {}

  @HttpCode(HttpStatus.OK)
  @Post('signup')
  async signUp(@Body() body: CreateUserDto) {
    return await this.authenticationService.signUp(body);
  }

  @Post('signin')
  async signIn(@Body() body: SignInDto) {
    return await this.authenticationService.signIn(body);
  }
}
