import { Injectable } from '@nestjs/common';
import * as nodemailer from 'nodemailer';
import { CreateMailDto } from './mail.dto';

@Injectable()
export class MailService {
  async create(createMailDto: CreateMailDto) {
    const transporter = nodemailer.createTransport({
      host: process.env.SMTP_HOST,
      port: parseInt(process.env.SMTP_PORT),
      secure: false,
      auth: {
        user: process.env.SMTP_USERNAME,
        pass: process.env.SMTP_PASSWORD,
      },
    });
    const info = await transporter.sendMail({
      from: '',
      to: createMailDto.email,
      subject: '',
      text: ``,
      html: ``,
    });
    if (!info) {
      return {
        message: 'Failed to send email',
        status: 400,
      };
    }
    return {
      message: 'Email sent',
      status: 200,
    };
  }
}
