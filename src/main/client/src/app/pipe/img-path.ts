import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {Utils} from '../utils/utils';

@Pipe({name: 'imgPath'})
export class ImgPathPipe implements PipeTransform {
    transform(url: string, external: any, defaultt: string) : string {
      return Utils.imgUrl(url, external, defaultt);
    }
}

@Pipe({name: 'thumbPath'})
export class ThumbPathPipe implements PipeTransform {
  transform(url: string, external: any) : string {
    return Utils.thumbUrl(url, external);
  }
}
