import { FieldType } from './field.types';

export interface CustomFieldModel {

  type: FieldType;
  value: string;
  column: string;
  status?: TreeStatus;
}
