import { FieldType } from './field.definitions';

export interface CustomFieldModel {
  type: FieldType;
  value: string;
  column: string;
}

export interface InputTestModel extends CustomFieldModel {

}
