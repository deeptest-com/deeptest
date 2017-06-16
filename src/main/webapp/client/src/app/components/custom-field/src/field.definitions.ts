
export interface CustomFieldDefinition {
  type: FieldType;
  value: string;
  column: string;
}

export class InputTestDefinition implements CustomFieldDefinition {

}

export enum FieldType {
  string,
  text,
  number,
  url,

  radio,
  checkbox,

  dropdown,
  multi_select,

  date,

  user,
  version,
  steps,
  results
}

