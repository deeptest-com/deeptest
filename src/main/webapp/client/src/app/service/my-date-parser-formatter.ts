import { NgbDateParserFormatter, NgbDateStruct } from '@ng-bootstrap/ng-bootstrap';
import { DatePipe } from '@angular/common';

export class MyDateParserFormatter extends NgbDateParserFormatter {
  datePipe = new DatePipe('en-US');
  constructor(private dateFormatString: string) {
    super();
  }
  format(date: NgbDateStruct): string {
    if (date === null) {
      return '';
    }
    try {
      let returnVal = this.datePipe.transform(new Date(date.year, date.month - 1, date.day), 'yyyy-MM-dd');

      return returnVal;
    } catch (e) {
      return '';
    }
  }
  parse(value: string): NgbDateStruct {
    console.log('===', new Date(value));

    let returnVal: NgbDateStruct;
    if (!value) {
      returnVal = null;
    } else {
      try {
        let dt = new Date(value);
        returnVal = { year: dt.getFullYear(), month: dt.getMonth() + 1, day: dt.getDate() };
      } catch (e) {
        returnVal = null;
      }
    }
    return returnVal;
  }
}
