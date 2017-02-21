import {Component, Input, Output ,OnInit} from '@angular/core';
import { ModalController, ViewController } from 'ionic-angular';
import {ImgPathPipe} from '../../pipes/img-path';
import {IosDatePipe} from '../../pipes/ios-date';
import {PubSubService} from '../../services/pub-sub-service';

@Component({
    selector: 'friend',
    templateUrl: 'build/components/modal/friend.html',
    providers: [PubSubService],
    pipes: [ImgPathPipe, IosDatePipe]
})

export class FriendComponent {
    @Input() item: any;

    constructor(private viewCtrl: ViewController) {

    }

    ngOnInit() {

    }
    
    cancel() {
       let data = 'cancel';
       this.viewCtrl.dismiss(data);
    }
    add() {
       let data = 'ok';
       this.viewCtrl.dismiss(data);
    }

    // select(item) {
    //     PubSubService.getInstance().changeCategory.emit(item);
    // }

}
