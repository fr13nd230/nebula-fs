import { IsDefined, IsEmail, IsNotEmpty, IsString } from 'class-validator';

export class SignInDto {
  @IsEmail()
  @IsDefined({ message: 'email field must be defined.' })
  @IsString({ message: 'email field must be a string. ' })
  @IsNotEmpty({ message: 'email field is required.' })
  email: string;

  @IsDefined({ message: 'password filed must be defined.' })
  @IsString({ message: 'password field must be a string.' })
  @IsNotEmpty({ message: 'password field is required.' })
  password: string;
}
