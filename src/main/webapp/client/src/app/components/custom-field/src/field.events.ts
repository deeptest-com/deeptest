export class FieldEvent {
  public constructor(public field: any) {
  }
}

export class FieldChangedEvent extends FieldEvent {
  public constructor(field: any) {
    super(field);
  }
}


