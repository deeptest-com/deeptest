import { FieldType } from './field.types';

export interface CustomFieldModel {
  type: FieldType;
  value: string;
  column: string;
}

export class InputTestModel implements CustomFieldModel {

}
