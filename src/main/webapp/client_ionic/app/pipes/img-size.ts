import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {Utils} from '../utils/utils';

@Pipe({name: 'imgSize'})
export class ImgSizePipe implements PipeTransform {
    transform(url: string, external: any) : string {

        return Utils.ImgSize(url);
    }
}
