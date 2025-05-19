import {
  IsDefined,
  IsEmail,
  IsNotEmpty,
  IsOptional,
  IsString,
  IsStrongPassword,
} from 'class-validator';

export class CreateUserDto {
  @IsEmail()
  @IsDefined({ message: 'email field must be defined.' })
  @IsString({ message: 'email field must be a string. ' })
  @IsNotEmpty({ message: 'email field is required.' })
  email: string;

  @IsStrongPassword({ minLength: 5 })
  @IsDefined({ message: 'password filed must be defined.' })
  @IsString({ message: 'password field must be a string.' })
  @IsNotEmpty({ message: 'password field is required.' })
  password: string;

  @IsString({ message: 'fullName filed must be a string.' })
  @IsOptional()
  fullName?: string;
}
