import {Component, Input, Output ,OnInit} from '@angular/core';
import {ImgPathPipe} from '../../pipes/img-path';
import {IosDatePipe} from '../../pipes/ios-date';
import {PubSubService} from '../../services/pub-sub-service';

@Component({
    selector: 'timeline',
    templateUrl: 'build/components/timeline/timeline.html',
    providers: [PubSubService],
    pipes: [ImgPathPipe, IosDatePipe]
})

export class TimelineComponent {
    @Input() items: any[];
    @Input() mode: string;

    constructor() {

    }

    ngOnInit() {
        
    }

    select(item) {
        PubSubService.getInstance().changeCategory.emit(item);
    }

}
