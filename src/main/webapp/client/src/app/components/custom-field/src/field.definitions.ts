
export class CustomFieldDefinition {
  fieldType: FieldType;
}

export class InputTestDefinition extends CustomFieldDefinition {

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

