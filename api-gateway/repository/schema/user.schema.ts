import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, SchemaTypes, Types } from 'mongoose';

@Schema()
export class User extends Document {
  @Prop({
    type: String,
    schema: 'email',
    required: true,
    unique: true,
    min: 8,
    max: 255,
  })
  email: string;

  @Prop({ type: String, required: true, min: 5, max: 255 })
  password: string;

  @Prop({ type: String, min: 3, max: 255 })
  fullName?: string;

  @Prop({ type: Date, default: new Date().toISOString() })
  created_at: Date;

  @Prop({ type: Date, default: new Date().toISOString() })
  updated_at: Date;
}

export const UserSchema = SchemaFactory.createForClass(User);
