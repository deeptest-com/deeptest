import {Pipe, PipeTransform} from "@angular/core";
import {DatePipe} from "@angular/common";

@Pipe({name: 'timePassed'})
export class TimePassedPipe implements PipeTransform {
  transform(timestamp: any): string {
    return this.timePassed(timestamp);
  }

  timePassed(timestamp: any, local = 'en-US', format = 'y/MM/dd HH:mm') {
    let now = new Date().getTime();
    let diffValue = now - timestamp;
    let result = '';
    let minute = 1000 * 60;
    let hour = minute * 60;
    let day = hour * 24;

    let _day = diffValue / day;
    let _hour = diffValue / hour;
    let _min = diffValue / minute;

    if (_day > 7) {
      result = new DatePipe(local).transform(timestamp, format);
    } else if (_day >= 1) {
      result = parseInt(_day + '') + "天前";
    } else if (_hour >= 1) {
      result = parseInt(_hour + '') + "个小时前";
    } else if (_min >= 1) {
      result = parseInt(_min + '') + "分钟前";
    } else {
      result = "刚刚";
    }
    return result;
  }
}
