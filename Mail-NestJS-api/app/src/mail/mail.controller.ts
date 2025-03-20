import { Body, Controller, Post } from '@nestjs/common';
import { ApiTags } from '@nestjs/swagger';
import { MailService } from './mail.service';
import { CreateMailDto } from './mail.dto';

@ApiTags('Mail')
@Controller('Mail')
export class MailController {
  constructor(private readonly mailService: MailService) {}
  @Post()
  create(@Body() body: CreateMailDto) {
    return this.mailService.create(body);
  }
}
