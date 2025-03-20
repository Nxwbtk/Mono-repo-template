import { ApiProperty } from '@nestjs/swagger';

export class CreateMailDto {
  @ApiProperty()
  otp: number;

  @ApiProperty()
  email: string;
}
